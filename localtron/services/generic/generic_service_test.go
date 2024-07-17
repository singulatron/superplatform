package genericservice_test

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/di"
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
		CreatedAt: time.Now().String(),
		Data:      map[string]interface{}{"key": "value"},
	}

	err = service.Create(table1, userId, obj)
	require.NoError(t, err)

	obj2 := &generictypes.GenericObject{
		Id:        uuid2,
		Table:     table2,
		CreatedAt: time.Now().String(),
		Data:      map[string]interface{}{"key": "value"},
	}

	err = service.Create(table2, userId, obj2)
	require.NoError(t, err)

	res, err := service.Find(table1, userId, false, []datastore.Condition{
		datastore.Id(uuid1),
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Equal(t, res[0].Id, uuid1)

	err = service.Create(table1, userId, obj)
	// entry already exists
	require.Error(t, err)

	res, err = service.Find(table1, userId, false, []datastore.Condition{
		datastore.Id(uuid2),
	})
	require.NoError(t, err)
	require.Equal(t, 0, len(res))

	err = service.Upsert(table1, otherUserId, obj)
	// unauthorized
	require.Error(t, err)

	err = service.Upsert(table1, userId, obj)
	require.NoError(t, err)

	res, err = service.Find(table1, userId, false, []datastore.Condition{
		datastore.All(),
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Contains(t, res[0].Id, uuid1)

	err = service.Delete(table1, otherUserId, []datastore.Condition{
		datastore.Id(obj.Id),
	})
	// no unauthorized but...
	require.NoError(t, err)

	// ...item wont be deleted
	// res, err = service.Find(table1, otherUserId, false, []datastore.Condition{
	// 	datastore.All(),
	// })
	// require.NoError(t, err)
	// require.Equal(t, 1, len(res))
	// require.Contains(t, res[0].Id, uuid1)
}
