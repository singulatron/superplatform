package genericservice_test

import (
	"testing"
	"time"

	"github.com/singulatron/singulatron/localtron/datastore"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	genericservice "github.com/singulatron/singulatron/localtron/services/generic"
	generictypes "github.com/singulatron/singulatron/localtron/services/generic/types"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	cs, err := configservice.NewConfigService()
	require.NoError(t, err)
	us, err := userservice.NewUserService(cs)
	require.NoError(t, err)
	fs, err := firehoseservice.NewFirehoseService(us)
	require.NoError(t, err)

	service, err := genericservice.NewGenericService(cs, fs, us)
	require.NoError(t, err)

	userId := "user_1"
	otherUserId := "user_2"

	obj := &generictypes.GenericObject{
		Id:        "1",
		Table:     "test_table",
		CreatedAt: time.Now().String(),
		Data:      map[string]interface{}{"key": "value"},
	}

	err = service.Create("test_table", userId, obj)
	require.NoError(t, err)

	obj2 := &generictypes.GenericObject{
		Id:        "1-2",
		Table:     "test_table2",
		CreatedAt: time.Now().String(),
		Data:      map[string]interface{}{"key": "value"},
	}

	err = service.Create("test_table2", userId, obj2)
	require.NoError(t, err)

	res, err := service.Find("test_table", userId, []datastore.Condition{
		datastore.Id("1"),
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Contains(t, res, obj)

	err = service.Create("test_table", userId, obj)
	// entry already exists
	require.Error(t, err)

	res, err = service.Find("test_table", userId, []datastore.Condition{
		datastore.Id("2"),
	})
	require.NoError(t, err)
	require.Equal(t, 0, len(res))

	err = service.Upsert("test_table", otherUserId, obj)
	// unauthorized
	require.Error(t, err)

	err = service.Upsert("test_table", userId, obj)
	require.NoError(t, err)

	err = service.Delete("test_table", otherUserId, []datastore.Condition{
		datastore.Id(obj.Id),
	})
	// no unauthorized but...
	require.NoError(t, err)

	// ...item wont be deleted
	res, err = service.Find("test_table", otherUserId, []datastore.Condition{
		datastore.All(),
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Contains(t, res, obj)
}
