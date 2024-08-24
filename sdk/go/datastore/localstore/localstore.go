/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
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

	"github.com/davecgh/go-spew/spew"
	"github.com/flusflas/dipper"
	"github.com/google/uuid"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/datastore/localstore/statemanager"
	"github.com/singulatron/singulatron/sdk/go/reflector"
)

type LocalStore struct {
	instance      any
	data          map[string]any
	mu            sync.RWMutex
	lastID        int
	inTransaction bool
	originalStore *LocalStore // Reference to the original store in case of transaction
	stateManager  *statemanager.StateManager
}

func NewLocalStore(instance any, filePath string) (*LocalStore, error) {
	if filePath == "" {
		tempFile, err := ioutil.TempFile("", uuid.NewString())
		if err != nil {
			return nil, err
		}
		filePath = tempFile.Name()
	}

	ls := &LocalStore{
		instance: instance,
		data:     make(map[string]any),
	}

	sm := statemanager.New(instance, func() []any {
		vals, _ := ls.Query(datastore.All()).Find()
		is := []any{}
		for _, v := range vals {
			is = append(is, v)
		}
		return is
	}, filePath)
	ls.stateManager = sm

	data, err := sm.LoadState()
	if err != nil {
		return nil, err
	}
	rowSlice := []datastore.Row{}

	if data != nil {
		v := reflect.ValueOf(data)

		if v.Kind() != reflect.Slice {
			panic("not a slice")
		}

		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i).Interface()
			rowSlice = append(rowSlice, elem.(datastore.Row))
		}
	}
	err = ls.CreateMany(rowSlice)
	if err != nil {
		return nil, err
	}

	go sm.PeriodicSaveState(5 * time.Second)

	return ls, nil
}

func (s *LocalStore) SetDebug(debug bool) {
}

func (s *LocalStore) Create(obj datastore.Row) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.createWithoutLock(obj)
}

func (s *LocalStore) createWithoutLock(obj datastore.Row) error {
	id := obj.GetId()
	_, ok := s.data[id]
	if ok {
		return datastore.ErrEntryAlreadyExists
	}

	v, err := reflector.DeepCopyIntoMap(obj)
	if err != nil {
		return err
	}

	s.data[id] = v

	s.stateManager.MarkChanged()
	return nil
}

func (s *LocalStore) CreateMany(objs []datastore.Row) error {
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

		v, err := reflector.DeepCopyIntoMap(obj)
		if err != nil {
			return err
		}

		s.data[id] = v
	}

	s.stateManager.MarkChanged()
	return nil
}

func (s *LocalStore) Upsert(obj datastore.Row) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, err := reflector.DeepCopyIntoMap(obj)
	if err != nil {
		return err
	}

	s.data[obj.GetId()] = v
	s.stateManager.MarkChanged()
	return nil
}

func (s *LocalStore) UpsertMany(objs []datastore.Row) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, obj := range objs {
		v, err := reflector.DeepCopyIntoMap(obj)
		if err != nil {
			return err
		}

		s.data[obj.GetId()] = v
	}
	s.stateManager.MarkChanged()
	return nil
}

func (s *LocalStore) Query(condition datastore.Condition, conditions ...datastore.Condition) datastore.QueryBuilder {
	q := &QueryBuilder{store: s}
	q.conditions = append(q.conditions, condition)
	q.conditions = append(q.conditions, conditions...)
	return q
}

func (s *LocalStore) BeginTransaction() (datastore.DataStore, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.inTransaction {
		return nil, errors.New("already in a transaction")
	}

	// Create a copy of the current store data
	newStore := &LocalStore{
		data:          make(map[string]any),
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

func (s *LocalStore) Commit() error {
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

func (s *LocalStore) Rollback() error {
	if !s.inTransaction || s.originalStore == nil {
		return errors.New("not in a transaction")
	}

	// Simply discard the transaction store
	s.inTransaction = false
	s.originalStore.inTransaction = false
	s.originalStore = nil

	return nil
}

func (s *LocalStore) IsInTransaction() bool {
	return s.inTransaction
}

type QueryBuilder struct {
	store        *LocalStore
	conditions   []datastore.Condition
	orderField   string
	orderDesc    bool
	orderByRand  bool
	limit        int64
	after        []any
	selectFields []string
}

func (q *QueryBuilder) OrderBy(option datastore.OrderBy, options ...datastore.OrderBy) datastore.QueryBuilder {
	if option.Field != "" {
		q.orderField = option.Field
		q.orderDesc = option.Desc
	} else {
		q.orderByRand = true
	}

	return q
}

func (q *QueryBuilder) Limit(limit int64) datastore.QueryBuilder {
	q.limit = limit
	return q
}

func (q *QueryBuilder) After(value ...any) datastore.QueryBuilder {
	q.after = value
	for i := range q.after {
		str, ok := q.after[i].(string)
		if !ok {
			continue
		}
		t, err := datastore.ParseAnyDate(str)
		if err == nil {
			q.after[i] = t
		}
	}
	return q
}

func (q *QueryBuilder) Select(fields ...string) datastore.QueryBuilder {
	q.selectFields = fields
	return q
}

func (q *QueryBuilder) Find() ([]datastore.Row, error) {
	q.store.mu.RLock()
	defer q.store.mu.RUnlock()

	var result []any
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

			if reflect.DeepEqual(toBaseType(vi), toBaseType(q.after[0])) {
				startIndex = i + 1
				break
			}
		}
		if startIndex != -1 {
			result = result[startIndex:]
		} else {
			result = []any{} // No matching "after" value found
		}
	}

	if q.limit > 0 && q.limit < int64(len(result)) {
		result = result[:q.limit]
	}

	sliceCopy, err := reflector.DeepCopySliceIntoType(result, q.store.instance)
	if err != nil {
		return nil, err
	}

	ret := []datastore.Row{}
	for _, v := range sliceCopy {
		ret = append(ret, v.(datastore.Row))
	}

	return ret, nil
}

func (q *QueryBuilder) FindOne() (datastore.Row, bool, error) {
	q.store.mu.RLock()
	defer q.store.mu.RUnlock()

	var empty datastore.Row

	for _, obj := range q.store.data {
		if q.match(obj) {
			cop, err := reflector.DeepCopyIntoType(obj, q.store.instance)
			if err != nil {
				return nil, false, err
			}

			return cop.(datastore.Row), true, nil
		}
	}

	return empty, false, nil
}

func (q *QueryBuilder) Count() (int64, error) {
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

func (q *QueryBuilder) Update(obj datastore.Row) error {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()

	found := false
	for id, existingObj := range q.store.data {
		if q.match(existingObj) {
			found = true

			v, err := reflector.DeepCopyIntoMap(obj)
			if err != nil {
				return err
			}

			q.store.data[id] = v
		}
	}

	if !found {
		return errors.New("no records to update")
	}

	q.store.stateManager.MarkChanged()

	return nil
}

func (q *QueryBuilder) Upsert(obj datastore.Row) error {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()

	q.store.stateManager.MarkChanged()

	found := false
	for id, existingObj := range q.store.data {
		if q.match(existingObj) {
			found = true

			v, err := reflector.DeepCopyIntoMap(obj)
			if err != nil {
				return err
			}

			q.store.data[id] = v
		}
	}

	if !found {
		return q.store.createWithoutLock(obj)
	}

	return nil
}

func (q *QueryBuilder) UpdateFields(fields map[string]interface{}) error {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()

	for id, obj := range q.store.data {
		if q.match(obj) {
			for field, value := range fields {

				err := setField(&obj, field, value)
				if err != nil {
					return err
				}
			}
			q.store.data[id] = obj
		}
	}
	q.store.stateManager.MarkChanged()
	return nil
}

func (q *QueryBuilder) Delete() error {
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

func (q *QueryBuilder) match(obj any) bool {
	for _, cond := range q.conditions {
		if cond.Equals != nil || cond.Contains != nil || cond.StartsWith != nil {
			var matchFunc func(subject, test any) bool
			var selector *datastore.FieldSelector
			var value any

			switch {
			case cond.Equals != nil:
				matchFunc = func(subject, test any) bool {
					subject = toBaseType(subject)
					test = toBaseType(test)
					if subject == "dipper: field not found" {
						panic("dipper")
					}

					return reflect.DeepEqual(test, subject)
				}
				selector = cond.Equals.Selector
				value = cond.Equals.Value
			case cond.Contains != nil:
				matchFunc = func(subject, test any) bool {
					testStr, testOk := test.(string)
					subjectStr, subjectOk := subject.(string)
					if !testOk || !subjectOk {
						return false
					}

					return strings.Contains(subjectStr, testStr)
				}

				selector = cond.Contains.Selector
				value = cond.Contains.Value
			case cond.StartsWith != nil:
				matchFunc = func(subject, test any) bool {
					testStr, testOk := test.(string)
					subjectStr, subjectOk := subject.(string)
					if !testOk || subjectOk {
						return false
					}
					return strings.HasPrefix(subjectStr, testStr)
				}
				selector = cond.StartsWith.Selector
				value = cond.StartsWith.Value
			}

			fieldNames := []string{}
			if selector.Field != "" {
				fieldNames = append(fieldNames, selector.Field)
			} else if selector.OneOf != nil {
				fieldNames = selector.OneOf
			}

			matched := false
			for _, fieldName := range fieldNames {
				fieldValue := getField(obj, fieldName)

				if fmt.Sprintf("%v", fieldValue) == "dipper: field not found" {
					continue
				}

				condValue := reflect.ValueOf(value)
				fieldV := reflect.ValueOf(fieldValue)

				if fieldV.Kind() == reflect.Slice {
					for i := 0; i < fieldV.Len(); i++ {
						if matchFunc(fieldV.Index(i).Interface(), condValue.Interface()) {
							matched = true
							continue
						}
					}

				} else if condValue.Kind() == reflect.Slice {
					for i := 0; i < condValue.Len(); i++ {
						if matchFunc(fieldValue, condValue.Index(i).Interface()) {
							matched = true
							continue
						}
					}
				} else {
					if matchFunc(fieldValue, value) {
						matched = true
					}
				}
			}
			if !matched {
				return false
			}
		} else if cond.All != nil {
			continue
		} else if cond.Intersects != nil {

		} else {
			spew.Dump(cond)
			panic(fmt.Sprintf("unkown condition %v", cond))
		}
	}
	return true
}

func fixFieldName(s string) string {
	parts := strings.Split(s, ".")
	for i := range parts {
		parts[i] = strings.ToLower(string(parts[i][0])) + parts[i][1:]
	}

	return strings.Join(parts, ".")
}

func getField(obj any, field string) interface{} {
	field = fixFieldName(field)

	return dipper.Get(obj, field)
}

func setField(obj any, field string, value interface{}) error {
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

func toBaseType(input interface{}) interface{} {
	val := reflect.ValueOf(input)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Invalid {
		return input
	}

	// Recursively decompose until we reach the base type
	for val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.String {
		return val.String()
	} else if val.Kind() == reflect.Int || val.Kind() == reflect.Int8 || val.Kind() == reflect.Int16 || val.Kind() == reflect.Int32 || val.Kind() == reflect.Int64 {
		return float64(val.Int())
	} else if val.Kind() == reflect.Uint || val.Kind() == reflect.Uint8 || val.Kind() == reflect.Uint16 || val.Kind() == reflect.Uint32 || val.Kind() == reflect.Uint64 {
		return val.Uint()
	} else if val.Kind() == reflect.Float32 || val.Kind() == reflect.Float64 {
		return val.Float()
	} else if val.Kind() == reflect.Bool {
		return val.Bool()
	}

	return val.Interface()
}
