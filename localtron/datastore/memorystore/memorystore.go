package memorystore

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"sync"

	"github.com/singulatron/singulatron/localtron/datastore"
)

type MemoryStore[T any] struct {
	data   map[string]T
	mu     sync.RWMutex
	lastID int
}

func NewMemoryStore[T any]() *MemoryStore[T] {
	return &MemoryStore[T]{
		data: make(map[string]T),
	}
}

func (s *MemoryStore[T]) Create(obj T) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.newID()
	s.data[id] = obj
	return nil
}

func (s *MemoryStore[T]) Read(id string) (T, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	obj, exists := s.data[id]
	if !exists {
		var empty T
		return empty, false, nil
	}
	return obj, true, nil
}

func (s *MemoryStore[T]) Update(id string, obj T) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.data[id]; !exists {
		return errors.New("not found")
	}
	s.data[id] = obj
	return nil
}

func (s *MemoryStore[T]) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.data[id]; !exists {
		return errors.New("not found")
	}
	delete(s.data, id)
	return nil
}

func (s *MemoryStore[T]) Query() datastore.QueryBuilder[T] {
	return &QueryBuilder[T]{store: s}
}

func (s *MemoryStore[T]) BatchCreate(objs []T) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, obj := range objs {
		id := s.newID()
		s.data[id] = obj
	}
	return nil
}

func (s *MemoryStore[T]) BatchUpdate(ids []string, objs []T) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(ids) != len(objs) {
		return errors.New("mismatched ids and objs")
	}
	for i, id := range ids {
		if _, exists := s.data[id]; !exists {
			return errors.New("not found")
		}
		s.data[id] = objs[i]
	}
	return nil
}

func (s *MemoryStore[T]) BatchDelete(ids []string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, id := range ids {
		if _, exists := s.data[id]; !exists {
			return errors.New("not found")
		}
		delete(s.data, id)
	}
	return nil
}

func (s *MemoryStore[T]) BeginTransaction() (datastore.Transaction[T], error) {
	return &MemoryTransaction[T]{store: s}, nil
}

func (s *MemoryStore[T]) newID() string {
	s.lastID++
	return fmt.Sprintf("%d", s.lastID)
}

type QueryBuilder[T any] struct {
	store        *MemoryStore[T]
	conditions   []func(T) bool
	orderField   string
	orderDesc    bool
	limit        int
	offset       int
	selectFields []string
}

func (q *QueryBuilder[T]) Where(field string, value interface{}) datastore.QueryBuilder[T] {
	q.conditions = append(q.conditions, func(obj T) bool {
		return getField(obj, field) == value
	})
	return q
}

func (q *QueryBuilder[T]) AndWhere(field string, value interface{}) datastore.QueryBuilder[T] {
	return q.Where(field, value)
}

func (q *QueryBuilder[T]) OrWhere(field string, value interface{}) datastore.QueryBuilder[T] {
	return q.Where(field, value)
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
		if !cond(obj) {
			return false
		}
	}
	return true
}

type MemoryTransaction[T any] struct {
	store *MemoryStore[T]
}

func (t *MemoryTransaction[T]) Commit() error {
	return nil
}

func (t *MemoryTransaction[T]) Rollback() error {
	return nil
}

func (t *MemoryTransaction[T]) DataStore() datastore.DataStore[T] {
	return t.store
}

func getField[T any](obj T, field string) interface{} {
	val := reflect.ValueOf(obj)
	return val.FieldByName(field).Interface()
}

func setField[T any](obj *T, field string, value interface{}) {
	val := reflect.ValueOf(obj).Elem()
	val.FieldByName(field).Set(reflect.ValueOf(value))
}

func compare(vi, vj interface{}, desc bool) bool {
	viVal := reflect.ValueOf(vi)
	vjVal := reflect.ValueOf(vj)

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
		panic("unsupported type for comparison")
	}
}
