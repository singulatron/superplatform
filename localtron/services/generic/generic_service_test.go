package genericservice_test

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/di"
	genericservice "github.com/singulatron/singulatron/localtron/services/generic"
	generictypes "github.com/singulatron/singulatron/localtron/services/generic/types"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	uniq := uuid.New().String()
	uniq = strings.Replace(uniq, "-", "", -1)[0:10]

	table1 := "test_table" + uniq
	table2 := "test_table2" + uniq

	universe, err := di.BigBang(di.UniverseOptions{
		Test: true,
	})
	require.NoError(t, err)
	service := universe.GenericService

	userId := "user_1"
	otherUserId := "user_2"

	uuid1 := uuid.New().String()
	uuid2 := uuid.New().String()

	obj := &generictypes.GenericObject{
		Id:        uuid1,
		Table:     table1,
		UserId:    userId,
		CreatedAt: time.Now().String(),
		Data:      map[string]interface{}{"key": "value"},
	}

	err = service.Create(obj)
	require.NoError(t, err)

	t.Run("user 1 can find its own private record", func(t *testing.T) {
		res, err := service.Find(genericservice.FindOptions{
			Table:  table1,
			UserId: userId,
			Public: false,
			Query: &datastore.Query{
				Conditions: []datastore.Condition{
					datastore.All(),
				},
			},
		})
		require.NoError(t, err)
		require.Equal(t, 1, len(res))
		require.Contains(t, res[0].Id, uuid1)
	})

	obj2 := &generictypes.GenericObject{
		Id:        uuid2,
		Table:     table2,
		UserId:    otherUserId,
		CreatedAt: time.Now().String(),
		Data:      map[string]interface{}{"key": "value"},
	}

	err = service.Create(obj2)
	require.NoError(t, err)

	t.Run("user 2 can find its own private record", func(t *testing.T) {
		res, err := service.Find(genericservice.FindOptions{
			Table:  table2,
			UserId: otherUserId,
			Public: false,
			Query: &datastore.Query{
				Conditions: []datastore.Condition{
					datastore.All(),
				}}})
		require.NoError(t, err)
		require.Equal(t, 1, len(res))
		require.Contains(t, res[0].Id, uuid2)
	})

	t.Run("find private for user 1", func(t *testing.T) {
		res, err := service.Find(genericservice.FindOptions{
			Table:  table1,
			UserId: userId,
			Public: false,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.Id(uuid1),
			}}})
		require.NoError(t, err)
		require.Equal(t, 1, len(res))
		require.Equal(t, res[0].Id, uuid1)
	})

	t.Run("find public for user 1", func(t *testing.T) {
		res, err := service.Find(genericservice.FindOptions{
			Table:  table1,
			UserId: userId,
			Public: true,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.Id(uuid1),
			}}})
		require.NoError(t, err)
		require.Equal(t, 0, len(res))
	})

	t.Run("already exists", func(t *testing.T) {
		err = service.Create(obj)
		require.Error(t, err)
	})

	t.Run("user 1 cannot see record of user 2", func(t *testing.T) {
		res, err := service.Find(genericservice.FindOptions{
			Table:  table1,
			UserId: userId,
			Public: false,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.Id(uuid2),
			}}})
		require.NoError(t, err)
		require.Equal(t, 0, len(res))
	})

	t.Run("user 2 cannot update record of user 1", func(t *testing.T) {
		obj.UserId = otherUserId
		err = service.Upsert(obj)
		// unauthorized
		require.Error(t, err)
		obj.UserId = userId
	})

	t.Run("user 1 can upsert its own reord", func(t *testing.T) {
		err = service.Upsert(obj)
		require.NoError(t, err)
	})

	t.Run("user 1 can find its own reord", func(t *testing.T) {
		res, err := service.Find(genericservice.FindOptions{
			Table:  table1,
			UserId: userId,
			Public: false,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.All(),
			}}})
		require.NoError(t, err)
		require.Equal(t, 1, len(res))
		require.Contains(t, res[0].Id, uuid1)
	})

	t.Run("user 2 cannot delete user 1's record", func(t *testing.T) {
		err = service.Delete(table1, otherUserId, []datastore.Condition{
			datastore.Id(obj.Id),
		})
		// no unauthorized but no error either...
		require.NoError(t, err)
	})

	// ...item wont be deleted
	t.Run("user 2 will no see other tables", func(t *testing.T) {
		res, err := service.Find(genericservice.FindOptions{
			Table:  table1,
			UserId: otherUserId,
			Public: false,
			Query: &datastore.Query{Conditions: []datastore.Condition{
				datastore.All(),
			}}})
		require.NoError(t, err)
		require.Equal(t, 0, len(res))
	})
}
