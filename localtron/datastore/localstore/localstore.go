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
package localstore

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/datastore/localstore/statemanager"
)

type LocalStore[T any] struct {
	data          map[string]T
	mu            sync.RWMutex
	lastID        int
	inTransaction bool
	originalStore *LocalStore[T] // Reference to the original store in case of transaction
	stateManager  *statemanager.StateManager[T]
}

func NewLocalStore[T any](filePath string) *LocalStore[T] {
	if filePath == "" {
		tempFile, err := ioutil.TempFile("", "example")
		if err != nil {
			panic(err)
		}
		filePath = tempFile.Name()
	}

	ls := &LocalStore[T]{
		data: make(map[string]T),
	}

	sm := statemanager.New(func() []T {
		vals, _ := ls.Query(datastore.All()).Find()
		return vals
	}, filePath)

	err := sm.LoadState()
	if err != nil {
		panic(err)
	}

	ls.stateManager = sm
	go sm.PeriodicSaveState(5 * time.Second)

	return ls
}

func (s *LocalStore[T]) Create(obj T) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.newID()
	s.data[id] = obj
	return nil
}

func (s *LocalStore[T]) CreateMany(objs []T) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, obj := range objs {
		id := s.newID()
		s.data[id] = obj
	}
	return nil
}

func (s *LocalStore[T]) Query(condition datastore.Condition, conditions ...datastore.Condition) datastore.QueryBuilder[T] {
	q := &QueryBuilder[T]{store: s}
	q.conditions = append(q.conditions, condition)
	q.conditions = append(q.conditions, conditions...)
	return q
}

func (s *LocalStore[T]) BeginTransaction() (datastore.DataStore[T], error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.inTransaction {
		return nil, errors.New("already in a transaction")
	}

	// Create a copy of the current store data
	newStore := &LocalStore[T]{
		data:          make(map[string]T),
		lastID:        s.lastID,
		inTransaction: true,
		originalStore: s,
	}

	for k, v := range s.data {
		newStore.data[k] = v
	}

	return newStore, nil
}

func (s *LocalStore[T]) Commit() error {
	if !s.inTransaction || s.originalStore == nil {
		return errors.New("not in a transaction")
	}

	s.originalStore.mu.Lock()
	defer s.originalStore.mu.Unlock()

	// Apply the changes to the original store
	for k, v := range s.data {
		s.originalStore.data[k] = v
	}

	// Reset transaction state
	s.inTransaction = false
	s.originalStore.inTransaction = false
	s.originalStore = nil

	return nil
}

func (s *LocalStore[T]) Rollback() error {
	if !s.inTransaction || s.originalStore == nil {
		return errors.New("not in a transaction")
	}

	// Simply discard the transaction store
	s.inTransaction = false
	s.originalStore.inTransaction = false
	s.originalStore = nil

	return nil
}

func (s *LocalStore[T]) IsInTransaction() bool {
	return s.inTransaction
}

func (s *LocalStore[T]) newID() string {
	s.lastID++
	return fmt.Sprintf("%d", s.lastID)
}

type QueryBuilder[T any] struct {
	store        *LocalStore[T]
	conditions   []datastore.Condition
	orderField   string
	orderDesc    bool
	limit        int
	offset       int
	selectFields []string
}

func (q *QueryBuilder[T]) OrderBy(field string, desc bool) datastore.QueryBuilder[T] {
	q.orderField = field
	q.orderDesc = desc
	return q
}

func (q *QueryBuilder[T]) Limit(limit int) datastore.QueryBuilder[T] {
	q.limit = limit
	return q
}

func (q *QueryBuilder[T]) Offset(offset int) datastore.QueryBuilder[T] {
	q.offset = offset
	return q
}

func (q *QueryBuilder[T]) Select(fields ...string) datastore.QueryBuilder[T] {
	q.selectFields = fields
	return q
}

func (q *QueryBuilder[T]) Find() ([]T, error) {
	q.store.mu.RLock()
	defer q.store.mu.RUnlock()
	var result []T
	for _, obj := range q.store.data {
		if q.match(obj) {
			result = append(result, obj)
		}
	}
	if q.orderField != "" {
		sort.Slice(result, func(i, j int) bool {
			vi, vj := getField(result[i], q.orderField), getField(result[j], q.orderField)
			return compare(vi, vj, q.orderDesc)
		})
	}
	if q.offset > 0 && q.offset < len(result) {
		result = result[q.offset:]
	}
	if q.limit > 0 && q.limit < len(result) {
		result = result[:q.limit]
	}
	return result, nil
}

func (q *QueryBuilder[T]) FindOne() (T, bool, error) {
	q.store.mu.RLock()
	defer q.store.mu.RUnlock()
	for _, obj := range q.store.data {
		if q.match(obj) {
			return obj, true, nil
		}
	}
	var empty T
	return empty, false, nil
}

func (q *QueryBuilder[T]) Count() (int64, error) {
	q.store.mu.RLock()
	defer q.store.mu.RUnlock()
	var count int64
	for _, obj := range q.store.data {
		if q.match(obj) {
			count++
		}
	}
	return count, nil
}

func (q *QueryBuilder[T]) Update(obj T) error {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()
	for id, existingObj := range q.store.data {
		if q.match(existingObj) {
			q.store.data[id] = obj
		}
	}
	return nil
}

func (q *QueryBuilder[T]) UpdateFields(fields map[string]interface{}) error {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()
	for id, obj := range q.store.data {
		if q.match(obj) {
			for field, value := range fields {
				setField(&obj, field, value)
			}
			q.store.data[id] = obj
		}
	}
	return nil
}

func (q *QueryBuilder[T]) Delete() error {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()
	for id, obj := range q.store.data {
		if q.match(obj) {
			delete(q.store.data, id)
		}
	}
	return nil
}

func (q *QueryBuilder[T]) match(obj T) bool {
	for _, cond := range q.conditions {
		if cond.Equal != nil && getField(obj, cond.Equal.FieldName) != cond.Equal.Value {
			return false
		}
	}
	return true
}

func getField[T any](obj T, field string) interface{} {
	field = strings.Title(field)

	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	fieldVal := val.FieldByName(field)
	if !fieldVal.IsValid() {
		return nil
	}

	return fieldVal.Interface()
}

func setField[T any](obj *T, field string, value interface{}) {
	field = strings.Title(field)

	val := reflect.ValueOf(obj).Elem()
	fieldVal := val.FieldByName(field)

	if fieldVal.IsValid() && fieldVal.CanSet() {
		fieldVal.Set(reflect.ValueOf(value))
	}
}

func compare(vi, vj interface{}, desc bool) bool {
	viVal := reflect.ValueOf(vi)
	vjVal := reflect.ValueOf(vj)

	if viVal.Kind() == reflect.Ptr {
		viVal = viVal.Elem()
	}
	if vjVal.Kind() == reflect.Ptr {
		vjVal = vjVal.Elem()
	}

	switch viVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if desc {
			return viVal.Int() > vjVal.Int()
		}
		return viVal.Int() < vjVal.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if desc {
			return viVal.Uint() > vjVal.Uint()
		}
		return viVal.Uint() < vjVal.Uint()
	case reflect.Float32, reflect.Float64:
		if desc {
			return viVal.Float() > vjVal.Float()
		}
		return viVal.Float() < vjVal.Float()
	case reflect.String:
		if desc {
			return viVal.String() > vjVal.String()
		}
		return viVal.String() < vjVal.String()
	default:
		return false
	}
}
