/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	"github.com/singulatron/singulatron/localtron/datastore"

	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

type ChatService struct {
	configService   *configservice.ConfigService
	userService     *userservice.UserService
	firehoseService *firehoseservice.FirehoseService

	messagesStore datastore.DataStore
	threadsStore  datastore.DataStore
	assetsStore   datastore.DataStore
}

func NewChatService(
	cs *configservice.ConfigService,
	fs *firehoseservice.FirehoseService,
	userService *userservice.UserService,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*ChatService, error) {
	threadsStore, err := datastoreFactory("threads", &chattypes.Thread{})
	if err != nil {
		return nil, err
	}
	messagesStore, err := datastoreFactory("messages", &chattypes.Message{})
	if err != nil {
		return nil, err
	}
	assetsStore, err := datastoreFactory("assets", &chattypes.Asset{})
	if err != nil {
		return nil, err
	}

	service := &ChatService{
		configService:   cs,
		firehoseService: fs,
		userService:     userService,

		messagesStore: messagesStore,
		threadsStore:  threadsStore,
		assetsStore:   assetsStore,
	}

	err = service.registerPermissions()
	if err != nil {
		return nil, err
	}

	return service, nil
}
