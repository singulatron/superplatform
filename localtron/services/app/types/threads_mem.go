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
package apptypes

import (
	"errors"
	"sync"
	"time"
)

type ThreadsMem struct {
	Threads []*ChatThread `json:"threads"`
	mutex   sync.Mutex
}

func NewThreadsMem() *ThreadsMem {
	return &ThreadsMem{
		Threads: []*ChatThread{},
	}
}

func (cf *ThreadsMem) ThreadsForeach(f func(i int, thread *ChatThread)) {
	cf.mutex.Lock()
	defer cf.mutex.Unlock()

	for i, v := range cf.Threads {
		f(i, v)
	}
}

func (cf *ThreadsMem) GetThreadsCopy() []*ChatThread {
	cf.mutex.Lock()
	defer cf.mutex.Unlock()

	ret := []*ChatThread{}

	for _, v := range cf.Threads {
		ret = append(ret, &ChatThread{
			Id:      v.Id,
			TopicId: v.TopicId,
			Name:    v.Name,
			Time:    v.Time,
		})
	}
	return ret
}

func (cf *ThreadsMem) AddThread(thread *ChatThread) {
	cf.mutex.Lock()
	defer cf.mutex.Unlock()

	cf.Threads = append(cf.Threads, thread)
}

func (cf *ThreadsMem) UpdateThread(thread *ChatThread) (*ChatThread, error) {
	cf.mutex.Lock()
	defer cf.mutex.Unlock()

	if thread.Id == "" {
		return nil, errors.New("no thread id supplied to update")
	}
	found := false
	for i, v := range cf.Threads {
		if v.Id == thread.Id {
			old := cf.Threads[i]
			if thread.Time == "" {
				thread.Time = old.Time
			}
			cf.Threads[i] = thread
			found = true
		}
	}
	if !found {
		if thread.Time == "" {
			thread.Time = time.Now().Format(time.RFC3339)
		}
		cf.Threads = append(cf.Threads, thread)
	}

	return thread, nil
}

func (cf *ThreadsMem) GetThreadById(threadId string) (*ChatThread, bool) {
	cf.mutex.Lock()
	defer cf.mutex.Unlock()

	for _, v := range cf.Threads {
		if v.Id == threadId {
			return v, true
		}
	}

	return nil, false
}

func (cf *ThreadsMem) DeleteThreadById(id string) {
	position := -1
	for i, chatThread := range cf.Threads {
		if chatThread.Id == id {
			position = i
		}
	}
	if position < 0 {
		return
	}

	cf.mutex.Lock()
	defer cf.mutex.Unlock()

	cf.Threads = append(cf.Threads[:position], cf.Threads[position+1:]...)
}
