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
package lib

import (
	"encoding/json"
	"sync"
)

type Row interface {
	GetId() string
	GetUpdatedAt() string
}

type MemoryStore[T Row] struct {
	items []T
	mutex sync.Mutex
}

func NewMemoryStore[T Row]() *MemoryStore[T] {
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

func (ms *MemoryStore[T]) ForeachStop(f func(i int, item T) bool) bool {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	for i, v := range ms.items {
		r := f(i, v)
		if r {
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

func (ms *MemoryStore[T]) Count(f func(item T) bool) int64 {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	var count int64
	for _, v := range ms.items {
		if f(v) {
			count++
		}
	}

	return count
}

func (ms *MemoryStore[T]) Find(matchFunc func(T) bool) (T, bool) {
	var def T
	ret := ms.Filter(matchFunc)
	if len(ret) > 0 {
		return ret[0], true
	}

	return def, false
}

func (ms *MemoryStore[T]) FindById(id string) (T, bool) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	for _, v := range ms.items {
		if v.GetId() == id {
			return v, true
		}
	}

	var def T
	return def, false
}

func (ms *MemoryStore[T]) FindByIds(ids []string) []T {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	index := map[string]struct{}{}
	for _, v := range ids {
		index[v] = struct{}{}
	}

	ret := []T{}
	for _, v := range ms.items {
		_, exists := index[v.GetId()]
		if exists {
			ret = append(ret, v)
		}
	}

	return ret
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

func (ms *MemoryStore[T]) DeleteByFunc(matchFunc func(T) bool) bool {
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
		return true
	}

	return false
}

func (ms *MemoryStore[T]) DeleteById(id string) bool {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	position := -1
	for i, item := range ms.items {
		if item.GetId() == id {
			position = i
		}
	}
	if position >= 0 {
		ms.items = append(ms.items[:position], ms.items[position+1:]...)
		return true
	}

	return false
}
