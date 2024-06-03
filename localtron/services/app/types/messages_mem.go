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
	"sort"
	"sync"
)

type MessagesMem struct {
	Messages []*ChatMessage `json:"messages"`
	mutex    sync.Mutex
}

func NewMessagesMem() *MessagesMem {
	return &MessagesMem{
		Messages: []*ChatMessage{},
	}
}

func (cf *MessagesMem) AddMessage(message *ChatMessage) {
	cf.mutex.Lock()
	defer cf.mutex.Unlock()

	cf.Messages = append(cf.Messages, message)
}

func (cf *MessagesMem) MessagesForeach(f func(i int, message *ChatMessage)) {
	cf.mutex.Lock()
	defer cf.mutex.Unlock()

	for i, v := range cf.Messages {
		f(i, v)
	}
}

func (cf *MessagesMem) GetMessagesByThreadId(threadId string) []*ChatMessage {
	cf.mutex.Lock()
	defer cf.mutex.Unlock()

	ret := []*ChatMessage{}
	for _, v := range cf.Messages {
		if v.ThreadId == threadId {
			ret = append(ret, v)
		}
	}

	sort.Sort(ByTime(ret))
	return ret
}

func (cf *MessagesMem) DeleteMessageById(id string) {
	position := -1
	for i, chatMessage := range cf.Messages {
		if chatMessage.Id == id {
			position = i
		}
	}
	if position < 0 {
		return
	}

	cf.mutex.Lock()
	defer cf.mutex.Unlock()

	cf.Messages = append(cf.Messages[:position], cf.Messages[position+1:]...)
}
