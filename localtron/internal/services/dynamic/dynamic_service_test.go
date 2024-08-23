package dynamicservice_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	clients "github.com/singulatron/singulatron/clients/go"
	"github.com/singulatron/singulatron/localtron/internal/di"
	dynamictypes "github.com/singulatron/singulatron/localtron/internal/services/dynamic/types"
	sdk "github.com/singulatron/singulatron/sdk/go"
	"github.com/singulatron/singulatron/sdk/go/datastore"
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

	token1, err := sdk.RegisterUser(router, "someuser", "pw123", "Some name")
	require.NoError(t, err)
	user1Router := router.SetBearerToken(token1)

	token2, err := sdk.RegisterUser(router, "someuser2", "pw123", "Some name 2")
	require.NoError(t, err)
	user2Router := router.SetBearerToken(token2)

	client1 := clients.NewAPIClient(&clients.Configuration{
		Servers: clients.ServerConfigurations{
			{
				URL:         server.URL,
				Description: "Default server",
			},
		},
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + token1,
		},
	})

	client2 := clients.NewAPIClient(&clients.Configuration{
		Servers: clients.ServerConfigurations{
			{
				URL:         server.URL,
				Description: "Default server",
			},
		},
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + token2,
		},
	})

	tokenReadRsp1, _, err := client1.UserSvcAPI.ReadUserByToken(context.Background()).Body(clients.UserSvcReadUserByTokenRequest{
		Token: clients.PtrString(token1),
	}).Execute()
	require.NoError(t, err)

	tokenReadRsp2, _, err := client2.UserSvcAPI.ReadUserByToken(context.Background()).Body(clients.UserSvcReadUserByTokenRequest{
		Token: clients.PtrString(token1),
	}).Execute()
	require.NoError(t, err)

	uuid1 := uuid.New().String()
	uuid2 := uuid.New().String()

	obj := &dynamictypes.Object{
		ObjectCreateFields: dynamictypes.ObjectCreateFields{
			Id:      uuid1,
			Table:   table1,
			Readers: []string{*tokenReadRsp1.User.Id},
			Data:    map[string]interface{}{"key": "value"},
		},
		CreatedAt: time.Now().String(),
	}

	err = user1Router.Post(context.Background(), "dynamic-svc", "/object", &dynamictypes.CreateObjectRequest{
		Object: &obj.ObjectCreateFields,
	}, nil)
	require.NoError(t, err)

	t.Run("user 1 can find its own private record", func(t *testing.T) {
		req := dynamictypes.QueryRequest{
			Table: table1,
			Query: &datastore.Query{
				Conditions: []datastore.Condition{
					datastore.All(),
				},
			},
		}
		rsp := dynamictypes.QueryResponse{}
		err = user1Router.Post(context.Background(), "dynamic-svc", "/objects", req, &rsp)
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Contains(t, rsp.Objects[0].Id, uuid1)
	})

	obj2 := &dynamictypes.Object{
		ObjectCreateFields: dynamictypes.ObjectCreateFields{
			Id:      uuid2,
			Table:   table2,
			Readers: []string{*tokenReadRsp2.User.Id},
			Data:    map[string]interface{}{"key": "value"},
		},
		CreatedAt: time.Now().String(),
	}

	err = user2Router.Post(context.Background(), "dynamic-svc", "/object", &dynamictypes.CreateObjectRequest{
		Object: &obj2.ObjectCreateFields,
	}, nil)
	require.NoError(t, err)

	t.Run("query user2 records", func(t *testing.T) {
		req := dynamictypes.QueryRequest{
			Table: table2,
			Query: &datastore.Query{
				Conditions: []datastore.Condition{
					datastore.All(),
				}},
		}
		rsp := dynamictypes.QueryResponse{}
		err = user2Router.Post(context.Background(), "dynamic-svc", "/objects", req, &rsp)
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Contains(t, rsp.Objects[0].Id, uuid2)
	})

	t.Run("query user1 records", func(t *testing.T) {
		req := dynamictypes.QueryRequest{
			Table: table1,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.Id(uuid1),
			}},
		}
		rsp := dynamictypes.QueryResponse{}
		err = user1Router.Post(context.Background(), "dynamic-svc", "/objects", req, &rsp)
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, rsp.Objects[0].Id, uuid1)
	})

	t.Run("already exists", func(t *testing.T) {
		err = user1Router.Post(context.Background(), "dynamic-svc", "/create",
			&dynamictypes.CreateObjectRequest{
				Object: &obj.ObjectCreateFields,
			},
			nil)
		require.Error(t, err)
	})

	t.Run("user 1 cannot see record of user 2", func(t *testing.T) {
		req := dynamictypes.QueryRequest{
			Table: table1,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.Id(uuid2),
			}},
		}
		rsp := dynamictypes.QueryResponse{}
		err = user1Router.Post(context.Background(), "dynamic-svc", "/objects", req, &rsp)
		require.NoError(t, err)
		require.Equal(t, 0, len(rsp.Objects))
	})

	t.Run("user 2 cannot update record of user 1", func(t *testing.T) {
		req := &dynamictypes.UpsertObjectRequest{
			Object: &obj.ObjectCreateFields,
		}
		err = user2Router.Put(context.Background(), "dynamic-svc", fmt.Sprintf("/object/%v", req.Object.Id), req, nil)
		// unauthorized
		require.Error(t, err)
	})

	t.Run("user 1 can upsert its own reord", func(t *testing.T) {
		req := &dynamictypes.UpsertObjectRequest{
			Object: &obj.ObjectCreateFields,
		}
		err = user1Router.Put(context.Background(), "dynamic-svc", fmt.Sprintf("/object/%v", req.Object.Id), req, nil)
		require.NoError(t, err)
	})

	t.Run("user 1 can find its own reord", func(t *testing.T) {
		req := dynamictypes.QueryRequest{
			Table: table1,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.All(),
			}},
		}
		rsp := dynamictypes.QueryResponse{}
		err = user1Router.Post(context.Background(), "dynamic-svc", "/objects", req, &rsp)
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Contains(t, rsp.Objects[0].Id, uuid1)
	})

	t.Run("user 2 cannot delete user 1's record", func(t *testing.T) {
		req := dynamictypes.DeleteObjectRequest{
			Table: table1,
			Conditions: []datastore.Condition{
				datastore.Id(obj.Id),
			},
		}

		err = user2Router.Post(context.Background(), "dynamic-svc", "/objects/delete", req, nil)
		// no unauthorized but no error either...
		require.NoError(t, err)
	})

	// ...item wont be deleted
	t.Run("user 2 will no see other tables", func(t *testing.T) {
		req := dynamictypes.QueryRequest{
			Table: table1,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.All(),
			}},
		}
		rsp := dynamictypes.QueryResponse{}
		err = user2Router.Post(context.Background(), "dynamic-svc", "/objects", req, &rsp)

		require.NoError(t, err)
		require.Equal(t, 0, len(rsp.Objects))
	})
}
