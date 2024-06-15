/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package localstore_test

import (
	"testing"

	"github.com/singulatron/singulatron/localtron/datastore"
	localstore "github.com/singulatron/singulatron/localtron/datastore/localstore"
	"github.com/stretchr/testify/assert"
)

type TestObject struct {
	Name  string
	Value int
}

func TestMemoryStore_CreateReadUpdateDelete(t *testing.T) {
	store := localstore.NewLocalStore[TestObject]("")

	obj := TestObject{Name: "test", Value: 10}
	err := store.Create(obj)
	assert.NoError(t, err)

	results, err := store.Query(datastore.Equal("Name", "test")).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	readObj := results[0]

	assert.Equal(t, obj, readObj)

	obj.Value = 20
	err = store.Query(datastore.Equal("Name", "test")).UpdateFields(map[string]interface{}{
		"Value": obj.Value,
	})
	assert.NoError(t, err)

	results, err = store.Query(datastore.Equal("Name", "test")).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	readObj = results[0]

	assert.Equal(t, obj, readObj)

	err = store.Query(datastore.Equal("Name", "test")).Delete()
	assert.NoError(t, err)

	results, err = store.Query(datastore.Equal("Name", "test")).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 0)
}

func TestMemoryStore_CreateManyUpdateDelete(t *testing.T) {
	store := localstore.NewLocalStore[TestObject]("")

	objs := []TestObject{
		{Name: "test1", Value: 10},
		{Name: "test2", Value: 20},
	}

	err := store.CreateMany(objs)
	assert.NoError(t, err)

	results, err := store.Query(datastore.Equal("Name", "test1")).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, objs[0], results[0])

	results, err = store.Query(datastore.Equal("Name", "test2")).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, objs[1], results[0])

	err = store.Query(datastore.Equal("Name", "test1")).UpdateFields(map[string]interface{}{
		"Value": 30,
	})
	assert.NoError(t, err)

	err = store.Query(datastore.Equal("Name", "test2")).UpdateFields(map[string]interface{}{
		"Value": 40,
	})
	assert.NoError(t, err)

	results, err = store.Query(datastore.Equal("Name", "test1")).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, 30, results[0].Value)

	results, err = store.Query(datastore.Equal("Name", "test2")).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, 40, results[0].Value)

	err = store.Query(datastore.Equal("Name", "test1")).Delete()
	assert.NoError(t, err)

	err = store.Query(datastore.Equal("Name", "test2")).Delete()
	assert.NoError(t, err)

	results, err = store.Query(datastore.Equal("Name", "test1")).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 0)

	results, err = store.Query(datastore.Equal("Name", "test2")).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 0)
}

func TestMemoryStore_Query(t *testing.T) {
	store := localstore.NewLocalStore[TestObject]("")

	objs := []TestObject{
		{Name: "test1", Value: 10},
		{Name: "test2", Value: 20},
		{Name: "test3", Value: 30},
	}

	err := store.CreateMany(objs)
	assert.NoError(t, err)

	results, err := store.Query(datastore.Equal("Value", 20)).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, objs[1], results[0])

	results, err = store.Query(datastore.All()).OrderBy("Value", true).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 3)
	assert.Equal(t, objs[2], results[0])
	assert.Equal(t, objs[1], results[1])
	assert.Equal(t, objs[0], results[2])

	results, err = store.Query(datastore.All()).Limit(2).Offset(1).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, objs[1], results[0])
	assert.Equal(t, objs[2], results[1])

	count, err := store.Query(datastore.Equal("Value", 10)).Count()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)

	err = store.Query(datastore.Equal("Value", 10)).UpdateFields(map[string]interface{}{
		"Value": 100,
	})
	assert.NoError(t, err)

	results, err = store.Query(datastore.Equal("Value", 100)).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, 100, results[0].Value)

	err = store.Query(datastore.Equal("Value", 100)).Delete()
	assert.NoError(t, err)

	results, err = store.Query(datastore.Equal("Value", 100)).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 0)
}

func TestMemoryStore_Transactions(t *testing.T) {
	store := localstore.NewLocalStore[TestObject]("")
	tx, err := store.BeginTransaction()
	assert.NoError(t, err)

	obj := TestObject{Name: "test", Value: 10}
	err = tx.Create(obj)
	assert.NoError(t, err)

	results, err := store.Query(datastore.Equal("Name", "test")).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 0)

	err = tx.Commit()
	assert.NoError(t, err)

	results, err = store.Query(datastore.Equal("Name", "test")).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	readObj := results[0]
	assert.Equal(t, obj, readObj)
}
