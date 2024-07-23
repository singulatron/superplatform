/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package firehoseservice

import (
	"log"
	"log/slog"
	"sync"

	"github.com/singulatron/singulatron/localtron/logger"

	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

type FirehoseService struct {
	userService usertypes.UserServiceI

	subscribers map[int]func(events []firehosetypes.Event)
	mu          sync.Mutex
	nextID      int
}

func NewFirehoseService(userService usertypes.UserServiceI) (*FirehoseService, error) {
	service := &FirehoseService{
		userService: userService,
		subscribers: make(map[int]func(events []firehosetypes.Event)),
	}
	err := service.registerPermissions()
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (fs *FirehoseService) PublishMany(events ...firehosetypes.Event) {
	for _, event := range events {
		logger.Debug("Event published",
			slog.String("eventName", event.Name()),
		)
	}
	fs.mu.Lock()
	defer fs.mu.Unlock()
	for _, subscriber := range fs.subscribers {
		go func(subscriber func(events []firehosetypes.Event)) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered in subscriber: %v", r)
				}
			}()
			subscriber(events)
		}(subscriber)
	}
}

func (fs *FirehoseService) Publish(event firehosetypes.Event) {
	fs.PublishMany(event)
}

func (fs *FirehoseService) Subscribe(callback func(events []firehosetypes.Event)) int {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	id := fs.nextID
	fs.subscribers[id] = callback
	fs.nextID++
	return id
}

func (fs *FirehoseService) Unsubscribe(id int) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	delete(fs.subscribers, id)
}
