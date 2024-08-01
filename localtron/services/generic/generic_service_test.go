package genericservice_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/di"
	generictypes "github.com/singulatron/singulatron/localtron/services/generic/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	uniq := uuid.New().String()
	uniq = strings.Replace(uniq, "-", "", -1)[0:10]

	table1 := "test_table" + uniq
	table2 := "test_table2" + uniq

	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)
	router := options.Router

	err = starterFunc()
	require.NoError(t, err)

	token, err := usertypes.RegisterUser(router, "someuser", "pw123", "Some name")
	require.NoError(t, err)
	user1Router := router.SetBearerToken(token)

	token, err = usertypes.RegisterUser(router, "someuser2", "pw123", "Some name 2")
	require.NoError(t, err)
	user2Router := router.SetBearerToken(token)

	uuid1 := uuid.New().String()
	uuid2 := uuid.New().String()

	obj := &generictypes.GenericObject{
		GenericObjectCreateFields: generictypes.GenericObjectCreateFields{
			Id:    uuid1,
			Table: table1,
			Data:  map[string]interface{}{"key": "value"},
		},
		CreatedAt: time.Now().String(),
	}

	err = user1Router.Post(context.Background(), "generic-service", "/object", &generictypes.CreateRequest{
		Object: &obj.GenericObjectCreateFields,
	}, nil)
	require.NoError(t, err)

	t.Run("user 1 can find its own private record", func(t *testing.T) {
		req := generictypes.QueryRequest{
			Table:  table1,
			Public: false,
			Query: &datastore.Query{
				Conditions: []datastore.Condition{
					datastore.All(),
				},
			},
		}
		rsp := generictypes.QueryResponse{}
		err = user1Router.Post(context.Background(), "generic-service", "/objects/query", req, &rsp)
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Contains(t, rsp.Objects[0].Id, uuid1)
	})

	obj2 := &generictypes.GenericObject{
		GenericObjectCreateFields: generictypes.GenericObjectCreateFields{
			Id:    uuid2,
			Table: table2,
			Data:  map[string]interface{}{"key": "value"},
		},
		CreatedAt: time.Now().String(),
	}

	err = user2Router.Post(context.Background(), "generic", "/create", &generictypes.CreateRequest{
		Object: &obj2.GenericObjectCreateFields,
	}, nil)
	require.NoError(t, err)

	t.Run("user 2 can find its own private record", func(t *testing.T) {
		req := generictypes.QueryRequest{
			Table:  table2,
			Public: false,
			Query: &datastore.Query{
				Conditions: []datastore.Condition{
					datastore.All(),
				}},
		}
		rsp := generictypes.QueryResponse{}
		err = user2Router.Post(context.Background(), "generic-service", "/objects/query", req, &rsp)
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Contains(t, rsp.Objects[0].Id, uuid2)
	})

	t.Run("find private for user 1", func(t *testing.T) {
		req := generictypes.QueryRequest{
			Table:  table1,
			Public: false,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.Id(uuid1),
			}},
		}
		rsp := generictypes.QueryResponse{}
		err = user1Router.Post(context.Background(), "generic-service", "/find", req, &rsp)
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, rsp.Objects[0].Id, uuid1)
	})

	t.Run("find public for user 1", func(t *testing.T) {
		req := generictypes.QueryRequest{
			Table:  table1,
			Public: true,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.Id(uuid1),
			}},
		}
		rsp := generictypes.QueryResponse{}
		err = user1Router.Post(context.Background(), "generic-service", "/objects/query", req, &rsp)
		require.NoError(t, err)
		require.Equal(t, 0, len(rsp.Objects))
	})

	t.Run("already exists", func(t *testing.T) {
		err = user1Router.Post(context.Background(), "generic-service", "/create",
			&generictypes.CreateRequest{
				Object: &obj.GenericObjectCreateFields,
			},
			nil)
		require.Error(t, err)
	})

	t.Run("user 1 cannot see record of user 2", func(t *testing.T) {
		req := generictypes.QueryRequest{
			Table:  table1,
			Public: false,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.Id(uuid2),
			}},
		}
		rsp := generictypes.QueryResponse{}
		err = user1Router.Post(context.Background(), "generic-service", "/objects/query", req, &rsp)
		require.NoError(t, err)
		require.Equal(t, 0, len(rsp.Objects))
	})

	t.Run("user 2 cannot update record of user 1", func(t *testing.T) {
		req := &generictypes.UpsertRequest{
			Object: &obj.GenericObjectCreateFields,
		}
		err = user2Router.Put(context.Background(), "generic-service", fmt.Sprintf("/object/%v", req.Object.Id), req, nil)
		// unauthorized
		require.Error(t, err)
	})

	t.Run("user 1 can upsert its own reord", func(t *testing.T) {
		req := &generictypes.UpsertRequest{
			Object: &obj.GenericObjectCreateFields,
		}
		err = user1Router.Put(context.Background(), "generic-service", fmt.Sprintf("/object/%v", req.Object.Id), req, nil)
		require.NoError(t, err)
	})

	t.Run("user 1 can find its own reord", func(t *testing.T) {
		req := generictypes.QueryRequest{
			Table:  table1,
			Public: false,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.All(),
			}},
		}
		rsp := generictypes.QueryResponse{}
		err = user1Router.Post(context.Background(), "generic", "/find", req, &rsp)
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Contains(t, rsp.Objects[0].Id, uuid1)
	})

	t.Run("user 2 cannot delete user 1's record", func(t *testing.T) {
		req := generictypes.DeleteRequest{
			Table: table1,
			Conditions: []datastore.Condition{
				datastore.Id(obj.Id),
			},
		}

		err = user2Router.Delete(context.Background(), "generic-service", "/objects/delete", req, nil)
		// no unauthorized but no error either...
		require.NoError(t, err)
	})

	// ...item wont be deleted
	t.Run("user 2 will no see other tables", func(t *testing.T) {
		req := generictypes.QueryRequest{
			Table:  table1,
			Public: false,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.All(),
			}},
		}
		rsp := generictypes.QueryResponse{}
		err = user2Router.Post(context.Background(), "generic-service", "/objects/query", req, &rsp)

		require.NoError(t, err)
		require.Equal(t, 0, len(rsp.Objects))
	})
}
