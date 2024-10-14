package dynamicservice_test

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/test"
	"github.com/singulatron/superplatform/server/internal/di"
	"github.com/stretchr/testify/require"

	client "github.com/singulatron/superplatform/clients/go"
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

	err = starterFunc()
	require.NoError(t, err)

	manyClients, err := test.MakeClients(options.Router, 2)
	require.NoError(t, err)
	client1 := manyClients[0]
	client2 := manyClients[1]

	tokenReadRsp1, _, err := client1.UserSvcAPI.ReadUserByToken(context.Background()).Execute()
	require.NoError(t, err)

	tokenReadRsp2, _, err := client2.UserSvcAPI.ReadUserByToken(context.Background()).Execute()
	require.NoError(t, err)

	uuid1 := sdk.Id(table1)
	uuid2 := sdk.Id(table2)

	obj := client.DynamicSvcObjectCreateFields{
		Id:       &uuid1,
		Table:    table1,
		Readers:  []string{"_self"},
		Writers:  []string{"_self"},
		Deleters: []string{"_self"},
		Data:     map[string]interface{}{"key": "value"},
	}

	_, _, err = client1.DynamicSvcAPI.CreateObject(context.Background()).Body(client.DynamicSvcCreateObjectRequest{
		Object: &obj,
	}).Execute()
	require.NoError(t, err)

	t.Run("user 1 can find its own private record", func(t *testing.T) {
		req := client.DynamicSvcQueryRequest{
			Table:   &table1,
			Readers: []string{*tokenReadRsp1.User.Id},
		}

		rsp, _, err := client1.DynamicSvcAPI.Query(context.Background()).Body(req).Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, uuid1, *rsp.Objects[0].Id)
	})

	obj2 := client.DynamicSvcObjectCreateFields{
		Id:      &uuid2,
		Table:   table2,
		Readers: []string{*tokenReadRsp2.User.Id},
		Data:    map[string]interface{}{"key": "value"},
	}

	_, _, err = client2.DynamicSvcAPI.CreateObject(context.Background()).Body(client.DynamicSvcCreateObjectRequest{
		Object: &obj2,
	}).Execute()
	require.NoError(t, err)

	t.Run("query user2 records", func(t *testing.T) {
		req := client.DynamicSvcQueryRequest{
			Table:   &table2,
			Readers: []string{*tokenReadRsp2.User.Id},
		}

		rsp, _, err := client2.DynamicSvcAPI.Query(context.Background()).Body(req).Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, uuid2, *rsp.Objects[0].Id)
	})

	t.Run("query user1 records", func(t *testing.T) {
		req := client.DynamicSvcQueryRequest{
			Table: &table1,
			Query: &client.DatastoreQuery{Filters: []client.DatastoreFilter{
				{
					Fields:     []string{"id"},
					Op:         client.OpEquals.Ptr(),
					JsonValues: sdk.Marshal([]any{uuid1}),
				},
			}},

			Readers: []string{*tokenReadRsp1.User.Id},
		}

		rsp, _, err := client1.DynamicSvcAPI.Query(context.Background()).Body(req).Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, uuid1, *rsp.Objects[0].Id)
	})

	t.Run("query user1 records with _self", func(t *testing.T) {
		req := client.DynamicSvcQueryRequest{
			Table: &table1,
			Query: &client.DatastoreQuery{Filters: []client.DatastoreFilter{
				{
					Fields:     []string{"id"},
					Op:         client.OpEquals.Ptr(),
					JsonValues: sdk.Marshal([]any{uuid1}),
				},
			}},

			Readers: []string{"_self"},
		}

		rsp, _, err := client1.DynamicSvcAPI.Query(context.Background()).Body(req).Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, uuid1, *rsp.Objects[0].Id)
	})

	t.Run("already exists", func(t *testing.T) {
		_, _, err = client1.DynamicSvcAPI.CreateObject(context.Background()).Body(client.DynamicSvcCreateObjectRequest{
			Object: &obj,
		}).Execute()

		require.Error(t, err)
	})

	t.Run("user 1 cannot see record of user 2", func(t *testing.T) {
		req := client.DynamicSvcQueryRequest{
			Table: &table1,
			Query: &client.DatastoreQuery{Filters: []client.DatastoreFilter{
				{
					Fields:     []string{"id"},
					Op:         client.OpEquals.Ptr(),
					JsonValues: sdk.Marshal([]any{uuid2}),
				},
			}},
			Readers: []string{*tokenReadRsp2.User.Id},
		}
		rsp, _, err := client1.DynamicSvcAPI.Query(context.Background()).Body(req).Execute()
		require.NoError(t, err)
		require.Equal(t, 0, len(rsp.Objects))
	})

	t.Run("user 2 cannot update record of user 1", func(t *testing.T) {
		req := &client.DynamicSvcUpsertObjectRequest{
			Object: &obj,
		}
		_, _, err = client2.DynamicSvcAPI.UpsertObject(context.Background(), *obj.Id).Body(*req).Execute()

		// unauthorized
		require.Error(t, err)
	})

	t.Run("user 1 can upsert its own record", func(t *testing.T) {
		req := &client.DynamicSvcUpsertObjectRequest{
			Object: &obj,
		}
		_, _, err = client1.DynamicSvcAPI.UpsertObject(context.Background(), *obj.Id).Body(*req).Execute()

		require.NoError(t, err)
	})

	t.Run("user 1 can find its own record", func(t *testing.T) {
		req := &client.DynamicSvcQueryRequest{
			Table:   client.PtrString(table1),
			Readers: []string{*tokenReadRsp1.User.Id},
		}
		rsp, _, err := client1.DynamicSvcAPI.Query(context.Background()).Body(*req).Execute()

		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, uuid1, *rsp.Objects[0].Id)
	})

	t.Run("user 2 cannot delete user 1's record", func(t *testing.T) {
		req := &client.DynamicSvcDeleteObjectRequest{
			Table: client.PtrString(table1),
			Filters: []client.DatastoreFilter{
				{
					Fields:     []string{"id"},
					Op:         client.OpEquals.Ptr(),
					JsonValues: sdk.Marshal([]any{obj.Id}),
				},
			},
		}

		_, _, err = client2.DynamicSvcAPI.DeleteObjects(context.Background()).Body(*req).Execute()

		require.Error(t, err)
	})

	// ...item wont be deleted
	t.Run("user 2 will no see other tables", func(t *testing.T) {
		req := &client.DynamicSvcQueryRequest{
			Table: client.PtrString(table1),
		}
		rsp, _, err := client2.DynamicSvcAPI.Query(context.Background()).Body(*req).Execute()

		require.NoError(t, err)
		require.Equal(t, 0, len(rsp.Objects))
	})
}
