package memorystore_test

import (
	"testing"

	"github.com/singulatron/singulatron/localtron/datastore/memorystore"
	"github.com/stretchr/testify/assert"
)

type TestObject struct {
	ID    string
	Name  string
	Value int
}

func TestMemoryStore_CreateReadUpdateDelete(t *testing.T) {
	store := memorystore.NewMemoryStore[TestObject]()

	obj := TestObject{ID: "1", Name: "test", Value: 10}
	err := store.Create(obj)
	assert.NoError(t, err)

	readObj, found, err := store.Read("1")
	assert.NoError(t, err)
	assert.Equal(t, true, found)
	assert.Equal(t, obj, readObj)

	obj.Value = 20
	err = store.Update("1", obj)
	assert.NoError(t, err)

	readObj, found, err = store.Read("1")
	assert.NoError(t, err)
	assert.Equal(t, true, found)
	assert.Equal(t, obj, readObj)

	err = store.Delete("1")
	assert.NoError(t, err)

	_, found, err = store.Read("1")
	assert.NoError(t, err)
	assert.Equal(t, false, found)
}

func TestMemoryStore_BatchCreateUpdateDelete(t *testing.T) {
	store := memorystore.NewMemoryStore[TestObject]()

	objs := []TestObject{
		{ID: "1", Name: "test1", Value: 10},
		{ID: "2", Name: "test2", Value: 20},
	}

	err := store.BatchCreate(objs)
	assert.NoError(t, err)

	readObj1, found, err := store.Read("1")
	assert.NoError(t, err)
	assert.Equal(t, true, found)
	assert.Equal(t, objs[0], readObj1)

	readObj2, found, err := store.Read("2")
	assert.NoError(t, err)
	assert.Equal(t, true, found)
	assert.Equal(t, objs[1], readObj2)

	objs[0].Value = 30
	objs[1].Value = 40
	err = store.BatchUpdate([]string{"1", "2"}, objs)
	assert.NoError(t, err)

	readObj1, found, err = store.Read("1")
	assert.NoError(t, err)
	assert.Equal(t, true, found)
	assert.Equal(t, objs[0], readObj1)

	readObj2, found, err = store.Read("2")
	assert.NoError(t, err)
	assert.Equal(t, true, found)
	assert.Equal(t, objs[1], readObj2)

	err = store.BatchDelete([]string{"1", "2"})
	assert.NoError(t, err)

	_, found, err = store.Read("1")
	assert.NoError(t, err)
	assert.Equal(t, false, found)

	_, found, err = store.Read("2")
	assert.NoError(t, err)
	assert.Equal(t, false, found)
}

func TestMemoryStore_Query(t *testing.T) {
	store := memorystore.NewMemoryStore[TestObject]()

	objs := []TestObject{
		{ID: "1", Name: "test1", Value: 10},
		{ID: "2", Name: "test2", Value: 20},
		{ID: "3", Name: "test3", Value: 30},
	}

	err := store.BatchCreate(objs)
	assert.NoError(t, err)

	results, err := store.Query().Where("Value", 20).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, objs[1], results[0])

	results, err = store.Query().OrderBy("Value", true).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 3)
	assert.Equal(t, objs[2], results[0])
	assert.Equal(t, objs[1], results[1])
	assert.Equal(t, objs[0], results[2])

	results, err = store.Query().Limit(2).Offset(1).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, objs[1], results[0])
	assert.Equal(t, objs[2], results[1])

	count, err := store.Query().Where("Value", 10).Count()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)

	err = store.Query().Where("Value", 10).UpdateFields(map[string]interface{}{
		"Value": 100,
	})
	assert.NoError(t, err)

	readObj, found, err := store.Read("1")
	assert.NoError(t, err)
	assert.Equal(t, true, found)
	assert.Equal(t, 100, readObj.Value)

	err = store.Query().Where("Value", 100).Delete()
	assert.NoError(t, err)

	_, found, err = store.Read("1")
	assert.NoError(t, err)
	assert.Equal(t, false, found)
}

func TestMemoryStore_Transactions(t *testing.T) {
	store := memorystore.NewMemoryStore[TestObject]()
	tx, err := store.BeginTransaction()
	assert.NoError(t, err)

	obj := TestObject{ID: "1", Name: "test", Value: 10}
	err = tx.DataStore().Create(obj)
	assert.NoError(t, err)

	_, found, err := store.Read("1")
	assert.NoError(t, err)
	assert.Equal(t, false, found)

	err = tx.Commit()
	assert.NoError(t, err)

	readObj, found, err := store.Read("1")
	assert.NoError(t, err)
	assert.Equal(t, true, found)
	assert.Equal(t, obj, readObj)
}
