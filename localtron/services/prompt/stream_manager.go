/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptservice

import (
	"sync"

	"github.com/singulatron/singulatron/localtron/clients/llm"
)

type SubscriberChan chan *llm.CompletionResponse

type StreamManager struct {
	streams map[string][]SubscriberChan
	history map[string][]*llm.CompletionResponse
	lock    sync.RWMutex
}

func NewStreamManager() *StreamManager {
	return &StreamManager{
		streams: make(map[string][]SubscriberChan),
		history: make(map[string][]*llm.CompletionResponse),
	}
}

func (sm *StreamManager) Subscribe(threadId string, subscriber SubscriberChan) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	// Add subscriber to the list
	sm.streams[threadId] = append(sm.streams[threadId], subscriber)

	// Send historical messages to the new subscriber
	go func() {
		sm.lock.RLock()
		defer sm.lock.RUnlock()
		for _, msg := range sm.history[threadId] {
			subscriber <- msg
		}
	}()
}

func (sm *StreamManager) Unsubscribe(threadId string, subscriber SubscriberChan) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	subs := sm.streams[threadId]
	for i, sub := range subs {
		if sub == subscriber {
			sm.streams[threadId] = append(subs[:i], subs[i+1:]...)
			close(subscriber) // Close the channel to signify no more data will be sent
			break
		}
	}
}

func (sm *StreamManager) Broadcast(threadId string, response *llm.CompletionResponse) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	if subscribers, ok := sm.streams[threadId]; ok {
		for _, subscriber := range subscribers {
			select {
			case subscriber <- response:
			default:
				// Handle full channel or unresponsive subscriber
			}
		}
	}
	// Append the new response to the history
	sm.history[threadId] = append(sm.history[threadId], response)
}
