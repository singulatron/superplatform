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
package appservice

import (
	"sync"

	"github.com/pkg/errors"

	"github.com/singulatron/singulatron/localtron/datastore"

	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	userservice "github.com/singulatron/singulatron/localtron/services/user"

	storefactoryservice "github.com/singulatron/singulatron/localtron/services/store_factory"
)

type AppService struct {
	configService   *configservice.ConfigService
	userService     *userservice.UserService
	firehoseService *firehoseservice.FirehoseService

	clientId string

	messagesStore datastore.DataStore[*apptypes.ChatMessage]
	threadsStore  datastore.DataStore[*apptypes.ChatThread]
	assetsStore   datastore.DataStore[*apptypes.Asset]

	logMutex sync.Mutex
}

func NewAppService(
	cs *configservice.ConfigService,
	fs *firehoseservice.FirehoseService,
	userService *userservice.UserService,
) (*AppService, error) {
	threadsStore, err := storefactoryservice.GetStore[*apptypes.ChatThread]("threads")
	if err != nil {
		return nil, err
	}
	messagesStore, err := storefactoryservice.GetStore[*apptypes.ChatMessage]("messages")
	if err != nil {
		return nil, err
	}
	assetsStore, err := storefactoryservice.GetStore[*apptypes.Asset]("assets")
	if err != nil {
		return nil, err
	}

	ci, err := cs.GetClientId()
	if err != nil {
		return nil, errors.Wrap(err, "app service canno get client id")
	}

	service := &AppService{
		configService:   cs,
		firehoseService: fs,
		userService:     userService,

		messagesStore: messagesStore,
		threadsStore:  threadsStore,
		assetsStore:   assetsStore,

		clientId: ci,
	}

	err = service.registerPermissions()
	if err != nil {
		return nil, err
	}

	return service, nil
}
