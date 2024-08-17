/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	"github.com/singulatron/singulatron/localtron/router"
	sdk "github.com/singulatron/singulatron/sdk/go"
	"github.com/singulatron/singulatron/sdk/go/datastore"

	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

type ChatService struct {
	router *router.Router

	messagesStore   datastore.DataStore
	threadsStore    datastore.DataStore
	assetsStore     datastore.DataStore
	credentialStore datastore.DataStore
}

func NewChatService(
	router *router.Router,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*ChatService, error) {
	threadsStore, err := datastoreFactory("chatThreads", &chattypes.Thread{})
	if err != nil {
		return nil, err
	}
	messagesStore, err := datastoreFactory("chatMessages", &chattypes.Message{})
	if err != nil {
		return nil, err
	}
	assetsStore, err := datastoreFactory("chatAssets", &chattypes.Asset{})
	if err != nil {
		return nil, err
	}
	credentialStore, err := datastoreFactory("chatCredentials", &usertypes.Credential{})
	if err != nil {
		return nil, err
	}

	service := &ChatService{
		router:          router,
		messagesStore:   messagesStore,
		threadsStore:    threadsStore,
		assetsStore:     assetsStore,
		credentialStore: credentialStore,
	}

	return service, nil
}

func (cs *ChatService) Start() error {
	token, err := sdk.RegisterService("chat-svc", "Chat Service", cs.router, cs.credentialStore)
	if err != nil {
		return err
	}
	cs.router = cs.router.SetBearerToken(token)

	return cs.registerPermissions()
}
