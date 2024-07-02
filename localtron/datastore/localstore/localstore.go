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
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/flusflas/dipper"
	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/datastore/localstore/statemanager"
)

type LocalStore[T datastore.Row] struct {
	data          map[string]T
	mu            sync.RWMutex
	lastID        int
	inTransaction bool
	originalStore *LocalStore[T] // Reference to the original store in case of transaction
	stateManager  *statemanager.StateManager[T]
}

func NewLocalStore[T datastore.Row](filePath string) *LocalStore[T] {
	if filePath == "" {
		tempFile, err := ioutil.TempFile("", uuid.NewString())
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
	ls.stateManager = sm

	data, err := sm.LoadState()
	if err != nil {
		panic(err)
	}
	err = ls.CreateMany(data)
	if err != nil {
		panic(err)
	}

	go sm.PeriodicSaveState(5 * time.Second)

	return ls
}

func (s *LocalStore[T]) SetDebug(debug bool) {
}

func (s *LocalStore[T]) Create(obj T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.createWithoutLock(obj)
}

func (s *LocalStore[T]) createWithoutLock(obj T) error {
	id := obj.GetId()
	_, ok := s.data[id]
	if ok {
		return datastore.ErrEntryAlreadyExists
	}
	s.data[id] = obj
	s.stateManager.MarkChanged()
	return nil
}

func (s *LocalStore[T]) CreateMany(objs []T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, obj := range objs {
		id := obj.GetId()
		_, ok := s.data[id]
		if ok {
			return datastore.ErrEntryAlreadyExists
		}
	}

	for _, obj := range objs {
		id := obj.GetId()
		s.data[id] = obj
	}

	s.stateManager.MarkChanged()
	return nil
}

func (s *LocalStore[T]) Upsert(obj T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[obj.GetId()] = obj
	s.stateManager.MarkChanged()
	return nil
}

func (s *LocalStore[T]) UpsertMany(objs []T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, obj := range objs {
		s.data[obj.GetId()] = obj
	}
	s.stateManager.MarkChanged()
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
		stateManager:  s.stateManager,
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

type QueryBuilder[T datastore.Row] struct {
	store        *LocalStore[T]
	conditions   []datastore.Condition
	orderField   string
	orderDesc    bool
	limit        int
	after        []any
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

func (q *QueryBuilder[T]) After(value ...any) datastore.QueryBuilder[T] {
	q.after = value
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

	if len(q.after) > 0 {
		startIndex := -1
		for i, obj := range result {
			vi := getField(obj, q.orderField)
			if reflect.DeepEqual(vi, q.after[0]) {
				startIndex = i + 1
				break
			}
		}
		if startIndex != -1 {
			result = result[startIndex:]
		} else {
			result = []T{} // No matching "after" value found
		}
	}

	if q.limit > 0 && q.limit < len(result) {
		result = result[:q.limit]
	}

	// deep copy result before returning
	var resultCopy []T
	bs, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return resultCopy, json.Unmarshal(bs, &resultCopy)
}

func (q *QueryBuilder[T]) FindOne() (T, bool, error) {
	q.store.mu.RLock()
	defer q.store.mu.RUnlock()

	var empty T

	for _, obj := range q.store.data {
		if q.match(obj) {
			var cop T
			// deep copy result before returning
			bs, err := json.Marshal(obj)
			if err != nil {
				return empty, false, err
			}

			return cop, true, json.Unmarshal(bs, &cop)
		}
	}

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

	found := false
	for id, existingObj := range q.store.data {
		if q.match(existingObj) {
			found = true
			q.store.data[id] = obj
		}
	}

	if !found {
		return errors.New("no records to update")
	}

	q.store.stateManager.MarkChanged()

	return nil
}

func (q *QueryBuilder[T]) Upsert(obj T) error {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()

	q.store.stateManager.MarkChanged()

	found := false
	for id, existingObj := range q.store.data {
		if q.match(existingObj) {
			found = true
			q.store.data[id] = obj
		}
	}

	if !found {
		return q.store.createWithoutLock(obj)
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
	q.store.stateManager.MarkChanged()
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
	q.store.stateManager.MarkChanged()
	return nil
}

func (q *QueryBuilder[T]) match(obj T) bool {
	for _, cond := range q.conditions {
		if cond.Equal != nil {
			fieldValue := getField(obj, cond.Equal.FieldName)

			condValue := reflect.ValueOf(cond.Equal.Value)
			if fieldV := reflect.ValueOf(fieldValue); fieldV.Kind() == reflect.Slice {
				matched := false
				for i := 0; i < fieldV.Len(); i++ {
					if reflect.DeepEqual(fieldV.Index(i).Interface(), condValue.Interface()) {
						matched = true
						continue
					}
				}
				if !matched {
					return false
				}
			} else if condValue.Kind() == reflect.Slice {
				matched := false
				for i := 0; i < condValue.Len(); i++ {
					if reflect.DeepEqual(fieldValue, condValue.Index(i).Interface()) {
						matched = true
						continue
					}
				}
				if !matched {
					return false
				}
			} else {
				if !reflect.DeepEqual(fieldValue, cond.Equal.Value) {
					return false
				}
			}
		} else if cond.All != nil {
			continue
		}
	}
	return true
}

func fixFieldName(s string) string {
	parts := strings.Split(s, ".")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, ".")
}

func getField[T any](obj T, field string) interface{} {
	field = fixFieldName(field)

	return dipper.Get(obj, field)
}

func setField[T any](obj *T, field string, value interface{}) error {
	field = fixFieldName(field)

	return dipper.Set(obj, field, value)
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
	case reflect.Struct:
		if viVal.Type() == reflect.TypeOf(time.Time{}) {
			viTime := viVal.Interface().(time.Time)
			vjTime := vjVal.Interface().(time.Time)
			if desc {
				return viTime.After(vjTime)
			}
			return viTime.Before(vjTime)
		}
	default:
		// Handle pointers to time.Time explicitly
		if viVal.Type() == reflect.TypeOf(&time.Time{}) && vjVal.Type() == reflect.TypeOf(&time.Time{}) {
			viTime := viVal.Interface().(*time.Time)
			vjTime := vjVal.Interface().(*time.Time)
			if viTime == nil || vjTime == nil {
				return false
			}
			if desc {
				return viTime.After(*vjTime)
			}
			return viTime.Before(*vjTime)
		}
	}
	return false
}
