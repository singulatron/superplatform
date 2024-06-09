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
package firehoseservice

import (
	"log"
	"log/slog"
	"sync"

	"github.com/singulatron/singulatron/localtron/lib"
	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

type FirehoseService struct {
	userService *userservice.UserService

	subscribers map[int]func(events []firehosetypes.Event)
	mu          sync.Mutex
	nextID      int
}

func NewFirehoseService(userService *userservice.UserService) (*FirehoseService, error) {
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
		lib.Logger.Debug("Event published",
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
