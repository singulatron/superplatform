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

	sdk "github.com/singulatron/singulatron/sdk/go"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/logger"
	"github.com/singulatron/singulatron/sdk/go/router"

	firehosetypes "github.com/singulatron/singulatron/localtron/internal/services/firehose/types"
)

type FirehoseService struct {
	router *router.Router

	subscribers map[int]func(events []*firehosetypes.Event)
	mu          sync.Mutex
	nextID      int

	credentialStore datastore.DataStore
}

func NewFirehoseService(r *router.Router, datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)) (*FirehoseService, error) {
	credentialStore, err := datastoreFactory("firehoseCredentials", &sdk.Credential{})
	if err != nil {
		return nil, err
	}

	service := &FirehoseService{
		router:          r,
		credentialStore: credentialStore,
		subscribers:     make(map[int]func(events []*firehosetypes.Event)),
	}

	return service, nil
}

func (fs *FirehoseService) Start() error {
	token, err := sdk.RegisterService("firehose-svc", "Firehose Service", fs.router, fs.credentialStore)
	if err != nil {
		return err
	}
	fs.router = fs.router.SetBearerToken(token)

	return fs.registerPermissions()
}

func (fs *FirehoseService) publishMany(events ...*firehosetypes.Event) {
	for _, event := range events {
		logger.Debug("Event published",
			slog.String("eventName", event.Name),
		)
	}
	fs.mu.Lock()
	defer fs.mu.Unlock()
	for _, subscriber := range fs.subscribers {
		go func(subscriber func(events []*firehosetypes.Event)) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered in subscriber: %v", r)
				}
			}()
			subscriber(events)
		}(subscriber)
	}
}

func (fs *FirehoseService) publish(event *firehosetypes.Event) {
	fs.publishMany(event)
}

func (fs *FirehoseService) subscribe(callback func(events []*firehosetypes.Event)) int {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	id := fs.nextID
	fs.subscribers[id] = callback
	fs.nextID++
	return id
}

func (fs *FirehoseService) unsubscribe(id int) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	delete(fs.subscribers, id)
}
