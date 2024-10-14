/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package streammanager

import (
	"sync"

	"github.com/singulatron/superplatform/sdk/go/clients/llm"
	prompttypes "github.com/singulatron/superplatform/server/internal/services/prompt/types"
)

type StreamManager struct {
	streams map[string][]prompttypes.SubscriberChan
	History map[string][]*llm.CompletionResponse
	lock    sync.RWMutex
}

func NewStreamManager() *StreamManager {
	return &StreamManager{
		streams: make(map[string][]prompttypes.SubscriberChan),
		History: make(map[string][]*llm.CompletionResponse),
	}
}

func (sm *StreamManager) Subscribe(threadId string, subscriber prompttypes.SubscriberChan) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	// Add subscriber to the list
	sm.streams[threadId] = append(sm.streams[threadId], subscriber)

	// Send historical messages to the new subscriber
	go func() {
		sm.lock.RLock()
		defer sm.lock.RUnlock()
		for _, msg := range sm.History[threadId] {
			subscriber <- msg
		}
	}()
}

func (sm *StreamManager) Unsubscribe(threadId string, subscriber prompttypes.SubscriberChan) {
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
	sm.History[threadId] = append(sm.History[threadId], response)
}
