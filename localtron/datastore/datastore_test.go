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
package datastore_test

import (
	"fmt"
	"testing"

	"github.com/singulatron/singulatron/localtron/datastore"
	localstore "github.com/singulatron/singulatron/localtron/datastore/localstore"
	"github.com/stretchr/testify/assert"
)

type TestObject struct {
	Name      string
	Value     int
	Age       int
	NickNames []string
}

func TestAll(t *testing.T) {
	stores := map[string]func() datastore.DataStore[TestObject]{
		"localStore": func() datastore.DataStore[TestObject] {
			return localstore.NewLocalStore[TestObject]("")
		},
	}
	tests := map[string]func(t *testing.T, store datastore.DataStore[TestObject]){
		"InClause":               InClause,
		"ReverseInClause":        ReverseInClause,
		"CreateReadUpdateDelete": CreateReadUpdateDelete,
		"CreateManyUpdateDelete": CreateManyUpdateDelete,
		"Query":                  Query,
		"Transactions":           Transactions,
	}

	for storeName, storeFunc := range stores {
		for testName, test := range tests {
			t.Run(fmt.Sprintf("%v %v", storeName, testName), func(t *testing.T) {
				test(t, storeFunc())
			})
		}
	}
}

func InClause(t *testing.T, store datastore.DataStore[TestObject]) {
	obj1 := TestObject{Name: "Alice", Value: 10, Age: 25}
	obj2 := TestObject{Name: "Bob", Value: 20, Age: 30}
	obj3 := TestObject{Name: "Charlie", Value: 30, Age: 35}

	err := store.Create(obj1)
	assert.NoError(t, err)
	err = store.Create(obj2)
	assert.NoError(t, err)
	err = store.Create(obj3)
	assert.NoError(t, err)

	// Test IN clause with string slice
	results, err := store.Query(datastore.Equal("Name", []string{"Alice", "Bob"})).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Contains(t, results, obj1)
	assert.Contains(t, results, obj2)

	// Test IN clause with int slice
	results, err = store.Query(datastore.Equal("Value", []int{10, 30})).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Contains(t, results, obj1)
	assert.Contains(t, results, obj3)

	// Test IN clause with empty slice (should return no results)
	results, err = store.Query(datastore.Equal("Age", []int{})).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 0)

	// Test IN clause with one element slice
	results, err = store.Query(datastore.Equal("Age", []int{30})).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Contains(t, results, obj2)

	// Clean up
	err = store.Query(datastore.Equal("Name", "Alice")).Delete()
	assert.NoError(t, err)
	err = store.Query(datastore.Equal("Name", "Bob")).Delete()
	assert.NoError(t, err)
	err = store.Query(datastore.Equal("Name", "Charlie")).Delete()
	assert.NoError(t, err)
}

func ReverseInClause(t *testing.T, store datastore.DataStore[TestObject]) {
	obj1 := TestObject{Name: "Alice", NickNames: []string{"A1", "A2"}}
	obj2 := TestObject{Name: "Bob", NickNames: []string{"B1"}}
	obj3 := TestObject{Name: "Charlie"}

	err := store.Create(obj1)
	assert.NoError(t, err)
	err = store.Create(obj2)
	assert.NoError(t, err)
	err = store.Create(obj3)
	assert.NoError(t, err)

	results, err := store.Query(
		datastore.Equal("NickNames", "A1"),
	).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Contains(t, results, obj1)

	results, err = store.Query(
		datastore.Equal("NickNames", "A2"),
	).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Contains(t, results, obj1)

	results, err = store.Query(
		datastore.Equal("NickNames", "B1"),
	).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Contains(t, results, obj2)

}

func CreateReadUpdateDelete(t *testing.T, store datastore.DataStore[TestObject]) {
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

func CreateManyUpdateDelete(t *testing.T, store datastore.DataStore[TestObject]) {
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

func Query(t *testing.T, store datastore.DataStore[TestObject]) {
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

	results, err = store.Query(datastore.All()).OrderBy("Name", true).Find()
	assert.NoError(t, err)
	assert.Len(t, results, 3)
	assert.Equal(t, objs[2], results[0])
	assert.Equal(t, objs[1], results[1])
	assert.Equal(t, objs[0], results[2])

	// order is nondeterministic when no OrderBy is supplied
	// results, err = store.Query(datastore.All()).Limit(2).Offset(1).Find()
	// assert.NoError(t, err)
	// assert.Len(t, results, 2)
	// assert.Equal(t, objs[1], results[0])
	// assert.Equal(t, objs[2], results[1])

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

func Transactions(t *testing.T, store datastore.DataStore[TestObject]) {
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
