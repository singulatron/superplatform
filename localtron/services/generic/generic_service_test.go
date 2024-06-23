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
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	cs, _ := configservice.NewConfigService()
	us, _ := userservice.NewUserService(cs)
	fs, _ := firehoseservice.NewFirehoseService(us)

	service, err := genericservice.NewGenericService(cs, fs, us)
	assert.NoError(t, err)

	userId := "user_1"
	otherUserId := "user_2"

	obj := &generictypes.GenericObject{
		Id:        "1",
		Table:     "test_table",
		CreatedAt: time.Now().String(),
		Data:      map[string]interface{}{"key": "value"},
	}

	err = service.Create("test_table", userId, obj)
	assert.NoError(t, err)

	obj2 := &generictypes.GenericObject{
		Id:        "1-2",
		Table:     "test_table2",
		CreatedAt: time.Now().String(),
		Data:      map[string]interface{}{"key": "value"},
	}

	err = service.Create("test_table2", userId, obj2)
	assert.NoError(t, err)

	res, err := service.Find("test_table", userId, []datastore.Condition{
		datastore.Id("1"),
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res, obj)

	err = service.Create("test_table", userId, obj)
	// entry already exists
	assert.Error(t, err)

	res, err = service.Find("test_table", userId, []datastore.Condition{
		datastore.Id("2"),
	})
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))

	err = service.Upsert("test_table", otherUserId, obj)
	// unauthorized
	assert.Error(t, err)

	err = service.Upsert("test_table", userId, obj)
	assert.NoError(t, err)

	err = service.Delete("test_table", otherUserId, []datastore.Condition{
		datastore.Id(obj.Id),
	})
	// no unauthorized but...
	assert.NoError(t, err)

	// ...item wont be deleted
	res, err = service.Find("test_table", otherUserId, []datastore.Condition{
		datastore.All(),
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res, obj)
}
