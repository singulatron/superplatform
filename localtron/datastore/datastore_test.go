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
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	localstore "github.com/singulatron/singulatron/localtron/datastore/localstore"
	"github.com/singulatron/singulatron/localtron/datastore/sqlstore"
	"github.com/stretchr/testify/require"
)

type Friend struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type TestObject struct {
	Name          string    `json:"name"`
	Value         int       `json:"value"`
	Age           int       `json:"age"`
	NickNames     []string  `json:"nickNames"`
	Friend        Friend    `json:"friend"`
	FriendPointer *Friend   `json:"friendPointer"`
	CreatedAt     time.Time `json:"createdAt"`
}

func (t TestObject) GetId() string {
	return t.Name
}

func TestAll(t *testing.T) {
	stores := map[string]func() datastore.DataStore[TestObject]{
		"localStore": func() datastore.DataStore[TestObject] {
			return localstore.NewLocalStore[TestObject]("")
		},
		"sqlStore": func() datastore.DataStore[TestObject] {
			table := uuid.New().String()
			table = strings.Replace(table, "-", "", -1)[0:10]
			store, err := sqlstore.NewSQLStore[TestObject](
				sqlstore.DriverPostGRES,
				"postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable",
				"table_"+table,
				true,
			)
			require.NoError(t, err)
			return store
		},
	}
	tests := map[string]func(t *testing.T, store datastore.DataStore[TestObject]){
		"Create":                 Create,
		"CreatedAt":              CreatedAt,
		"Upsert":                 Upsert,
		"InClause":               InClause,
		"ReverseInClause":        ReverseInClause,
		"CreateReadUpdateDelete": CreateReadUpdateDelete,
		"CreateManyUpdateDelete": CreateManyUpdateDelete,
		"Query":                  Query,
		"Transactions":           Transactions,
		"DotNotation":            DotNotation,
		"Pagination":             Pagination,
	}

	for testName, test := range tests {
		for storeName, storeFunc := range stores {
			t.Run(fmt.Sprintf("%v %v", storeName, testName), func(t *testing.T) {
				store := storeFunc()
				test(t, store)
			})
		}
	}
}

func Pagination(t *testing.T, store datastore.DataStore[TestObject]) {
	for i := 1; i <= 10; i++ {
		obj := TestObject{Name: fmt.Sprintf("PaginationTest%d", i), Value: i}
		err := store.Create(obj)
		require.NoError(t, err)
	}

	results, err := store.Query(datastore.All()).OrderBy("Value", true).Limit(5).Find()
	require.NoError(t, err)
	require.Len(t, results, 5)
	require.Equal(t, "PaginationTest10", results[0].Name)
	require.Equal(t, 10, results[0].Value)

	lastValue := results[len(results)-1].Value
	results, err = store.Query(datastore.All()).OrderBy("Value", true).Limit(5).After(lastValue).Find()
	require.NoError(t, err)
	require.Len(t, results, 5)
	require.Equal(t, "PaginationTest5", results[0].Name)
	require.Equal(t, 5, results[0].Value)
}

func CreatedAt(t *testing.T, store datastore.DataStore[TestObject]) {
	obj1 := TestObject{Name: "A1", Value: 10, CreatedAt: time.Now()}
	obj2 := TestObject{Name: "A2", Value: 10, CreatedAt: time.Now().Add(time.Minute)}
	obj3 := TestObject{Name: "A3", Value: 20, CreatedAt: time.Now().Add(2 * time.Minute)}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	res, err := store.Query(
		datastore.All(),
		datastore.Equal("Value", 101),
	).OrderBy("CreatedAt", false).Find()
	require.NoError(t, err)
	require.Equal(t, 0, len(res))

	res, err = store.Query(
		datastore.All(),
		datastore.Equal("Value", 10),
	).OrderBy("CreatedAt", false).Find()
	require.NoError(t, err)
	require.Equal(t, 2, len(res))
	require.Equal(t, "A1", res[0].Name)

	res, err = store.Query(
		datastore.All(),
		datastore.Equal("Value", 10),
	).OrderBy("CreatedAt", true).Find()

	require.NoError(t, err)
	require.Equal(t, 2, len(res))
	require.Equal(t, "A2", res[0].Name)
}

func Create(t *testing.T, store datastore.DataStore[TestObject]) {
	obj1 := TestObject{Name: "AliceCreate", Value: 10, Age: 25}

	err := store.Create(obj1)
	require.NoError(t, err)

	err = store.Create(obj1)
	require.Error(t, err)
}

func Upsert(t *testing.T, store datastore.DataStore[TestObject]) {
	obj1 := TestObject{Name: "AliceCreate", Value: 10, Age: 25}

	err := store.Upsert(obj1)
	require.NoError(t, err)

	err = store.Upsert(obj1)
	require.NoError(t, err)
}

func InClause(t *testing.T, store datastore.DataStore[TestObject]) {
	obj1 := TestObject{Name: "Alice", Value: 10, Age: 25}
	obj2 := TestObject{Name: "Bob", Value: 20, Age: 30}
	obj3 := TestObject{Name: "Charlie", Value: 30, Age: 35}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	// Test IN clause with string slice
	results, err := store.Query(datastore.Equal("Name", []string{"Alice", "Bob"})).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj2)

	// Test IN clause with int slice
	results, err = store.Query(datastore.Equal("Value", []int{10, 30})).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj3)

	// Test IN clause with empty slice (should return no results)
	results, err = store.Query(datastore.Equal("Age", []int{})).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	// Test IN clause with one element slice
	results, err = store.Query(datastore.Equal("Age", []int{30})).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj2)

	// Clean up
	err = store.Query(datastore.Equal("Name", "Alice")).Delete()
	require.NoError(t, err)
	err = store.Query(datastore.Equal("Name", "Bob")).Delete()
	require.NoError(t, err)
	err = store.Query(datastore.Equal("Name", "Charlie")).Delete()
	require.NoError(t, err)
}

func DotNotation(t *testing.T, store datastore.DataStore[TestObject]) {
	obj1 := TestObject{Name: "Alice", Value: 10, Age: 25,
		Friend:        Friend{Name: "AliceFriend", Age: 26},
		FriendPointer: &Friend{Name: "AliceFriendP", Age: 27},
	}
	obj2 := TestObject{Name: "Bob", Value: 20, Age: 30,
		Friend:        Friend{Name: "BobFriend", Age: 31},
		FriendPointer: &Friend{Name: "BobFriendP", Age: 32},
	}
	obj3 := TestObject{Name: "Charlie", Value: 30, Age: 35,
		Friend:        Friend{Name: "CharlieFriend", Age: 36},
		FriendPointer: &Friend{Name: "CharlieFriendP", Age: 37},
	}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	// Test IN clause with string slice
	results, err := store.Query(datastore.Equal("Friend.Name", []string{"AliceFriend", "BobFriend"})).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj2)

	results, err = store.Query(datastore.Equal("friendPointer.name", []string{"AliceFriendP", "BobFriendP"})).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj2)

	// Test IN clause with int slice
	results, err = store.Query(datastore.Equal("Friend.Age", []int{26, 36})).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj3)

	// Test IN clause with empty slice (should return no results)
	results, err = store.Query(datastore.Equal("Friend.Age", []int{})).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	// Test Ordering
	results, err = store.Query(datastore.All()).OrderBy("Friend.Age", false).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, "Alice", results[0].Name)
	require.Equal(t, "Bob", results[1].Name)
	require.Equal(t, "Charlie", results[2].Name)

	results, err = store.Query(datastore.All()).OrderBy("Friend.Age", true).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, "Charlie", results[0].Name)
	require.Equal(t, "Bob", results[1].Name)
	require.Equal(t, "Alice", results[2].Name)

	// Test IN clause with one element slice
	results, err = store.Query(datastore.Equal("FriendPointer.Age", []int{32})).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj2)

	// Clean up
	err = store.Query(datastore.Equal("Friend.Name", "AliceFriend")).Delete()
	require.NoError(t, err)
	err = store.Query(datastore.Equal("Friend.Name", "BobFriendP")).Delete()
	require.NoError(t, err)
	err = store.Query(datastore.Equal("Friend.Name", "CharlieFriend")).Delete()
	require.NoError(t, err)
}

func ReverseInClause(t *testing.T, store datastore.DataStore[TestObject]) {
	obj1 := TestObject{Name: "Alice", NickNames: []string{"A1", "A2"}}
	obj2 := TestObject{Name: "Bob", NickNames: []string{"B1"}}
	obj3 := TestObject{Name: "Charlie"}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	results, err := store.Query(
		datastore.Equal("NickNames", "A1"),
	).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj1)

	results, err = store.Query(
		datastore.Equal("NickNames", "A2"),
	).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj1)

	results, err = store.Query(
		datastore.Equal("NickNames", "B1"),
	).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj2)

}

func CreateReadUpdateDelete(t *testing.T, store datastore.DataStore[TestObject]) {
	obj := TestObject{Name: "test", Value: 10}
	err := store.Create(obj)
	require.NoError(t, err)

	results, err := store.Query(datastore.Equal("Name", "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	readObj := results[0]

	require.Equal(t, obj, readObj)

	obj.Value = 20
	err = store.Query(datastore.Equal("Name", "test")).UpdateFields(map[string]interface{}{
		"Value": obj.Value,
	})
	require.NoError(t, err)

	results, err = store.Query(datastore.Equal("Name", "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	readObj = results[0]

	require.Equal(t, obj, readObj)

	err = store.Query(datastore.Equal("Name", "test")).Delete()
	require.NoError(t, err)

	results, err = store.Query(datastore.Equal("Name", "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)
}

func CreateManyUpdateDelete(t *testing.T, store datastore.DataStore[TestObject]) {
	objs := []TestObject{
		{Name: "test1", Value: 10},
		{Name: "test2", Value: 20},
	}

	err := store.CreateMany(objs)
	require.NoError(t, err)

	results, err := store.Query(datastore.Equal("Name", "test1")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, objs[0], results[0])

	results, err = store.Query(datastore.Equal("Name", "test2")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, objs[1], results[0])

	err = store.Query(datastore.Equal("Name", "test1")).UpdateFields(map[string]interface{}{
		"Value": 30,
	})
	require.NoError(t, err)

	err = store.Query(datastore.Equal("Name", "test2")).UpdateFields(map[string]interface{}{
		"Value": 40,
	})
	require.NoError(t, err)

	results, err = store.Query(datastore.Equal("Name", "test1")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, 30, results[0].Value)

	results, err = store.Query(datastore.Equal("Name", "test2")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, 40, results[0].Value)

	err = store.Query(datastore.Equal("Name", "test1")).Delete()
	require.NoError(t, err)

	err = store.Query(datastore.Equal("Name", "test2")).Delete()
	require.NoError(t, err)

	results, err = store.Query(datastore.Equal("Name", "test1")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	results, err = store.Query(datastore.Equal("Name", "test2")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)
}

func Query(t *testing.T, store datastore.DataStore[TestObject]) {
	objs := []TestObject{
		{Name: "queryTest1", Value: 10},
		{Name: "queryTest2", Value: 20},
		{Name: "queryTest3", Value: 30},
	}

	err := store.CreateMany(objs)
	require.NoError(t, err)

	results, err := store.Query(datastore.Equal("Value", 20)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, objs[1], results[0])

	results, err = store.Query(datastore.All()).OrderBy("Value", true).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, objs[2], results[0])
	require.Equal(t, objs[1], results[1])
	require.Equal(t, objs[0], results[2])

	results, err = store.Query(datastore.All()).OrderBy("Name", true).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, objs[2], results[0])
	require.Equal(t, objs[1], results[1])
	require.Equal(t, objs[0], results[2])

	// order is nondeterministic when no OrderBy is supplied
	// results, err = store.Query(datastore.All()).Limit(2).Offset(1).Find()
	// require.NoError(t, err)
	// require.Len(t, results, 2)
	// require.Equal(t, objs[1], results[0])
	// require.Equal(t, objs[2], results[1])

	count, err := store.Query(datastore.Equal("Value", 10)).Count()
	require.NoError(t, err)
	require.Equal(t, int64(1), count)

	err = store.Query(datastore.Equal("Value", 10)).UpdateFields(map[string]interface{}{
		"Value": 100,
	})
	require.NoError(t, err)

	results, err = store.Query(datastore.Equal("Value", 100)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, 100, results[0].Value)

	err = store.Query(datastore.Equal("Value", 100)).Delete()
	require.NoError(t, err)

	results, err = store.Query(datastore.Equal("Value", 100)).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)
}

func Transactions(t *testing.T, store datastore.DataStore[TestObject]) {
	tx, err := store.BeginTransaction()
	require.NoError(t, err)

	obj := TestObject{Name: "test", Value: 10}
	err = tx.Create(obj)
	require.NoError(t, err)

	results, err := store.Query(datastore.Equal("Name", "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	err = tx.Commit()
	require.NoError(t, err)

	results, err = store.Query(datastore.Equal("Name", "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	readObj := results[0]
	require.Equal(t, obj, readObj)
}
