package lib

import (
	"encoding/json"
	"sync"
)

type MemoryStore[T any] struct {
	items []T
	mutex sync.Mutex
}

func NewMemoryStore[T any]() *MemoryStore[T] {
	return &MemoryStore[T]{
		items: []T{},
	}
}

func (ms *MemoryStore[T]) Add(item T) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ms.items = append(ms.items, item)
}

func (ms *MemoryStore[T]) Reset(items []T) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ms.items = items
}

func (ms *MemoryStore[T]) Foreach(f func(i int, item T)) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	for i, v := range ms.items {
		f(i, v)
	}
}

func (ms *MemoryStore[T]) UpdateByFunc(matchFunc func(T) bool, updateValue T) bool {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	for i, item := range ms.items {
		if matchFunc(item) {
			ms.items[i] = updateValue
			return true
		}
	}
	return false
}

func (ms *MemoryStore[T]) Filter(f func(item T) bool) []T {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	var ret []T
	for _, v := range ms.items {
		if f(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func (ms *MemoryStore[T]) Find(matchFunc func(T) bool) (T, bool) {
	var def T
	ret := ms.Filter(matchFunc)
	if len(ret) > 0 {
		return ret[0], true
	}

	return def, false
}

func (ms *MemoryStore[T]) SliceCopy() []T {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	newArray := make([]T, len(ms.items))
	copy(newArray, ms.items)
	return newArray
}

func (ms *MemoryStore[T]) DeepCopy() ([]T, error) {
	var ret []T
	bytes, err := json.Marshal(ms.items)
	if err != nil {
		return nil, err
	}

	return ret, json.Unmarshal(bytes, &ret)
}

func (ms *MemoryStore[T]) DeleteByFunc(matchFunc func(T) bool) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	position := -1
	for i, item := range ms.items {
		if matchFunc(item) {
			position = i
			break
		}
	}
	if position >= 0 {
		ms.items = append(ms.items[:position], ms.items[position+1:]...)
	}
}
