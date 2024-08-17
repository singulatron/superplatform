/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package modelservice

import (
	"sync"

	sdk "github.com/singulatron/singulatron/sdk/go"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/router"

	modeltypes "github.com/singulatron/singulatron/localtron/internal/services/model/types"
)

type ModelService struct {
	modelStateMutex sync.Mutex
	modelPortMap    map[int]*modeltypes.ModelState

	router         *router.Router
	modelsStore    datastore.DataStore
	platformsStore datastore.DataStore

	credentialStore datastore.DataStore
}

func NewModelService(
	router *router.Router,
	datastoreFactory func(tableName string, insance any) (datastore.DataStore, error),
) (*ModelService, error) {
	srv := &ModelService{
		router: router,

		modelPortMap: map[int]*modeltypes.ModelState{},
	}
	modelStore, err := datastoreFactory("modelSvcModels", &modeltypes.Model{})
	if err != nil {
		return nil, err
	}
	srv.modelsStore = modelStore

	platformsStore, err := datastoreFactory("modelSvcPlatforms", &modeltypes.Platform{})
	if err != nil {
		return nil, err
	}
	srv.platformsStore = platformsStore

	credentialStore, err := datastoreFactory("modelSvcCredentials", &sdk.Credential{})
	if err != nil {
		return nil, err
	}
	srv.credentialStore = credentialStore

	err = srv.bootstrapModels()
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func (ms *ModelService) Start() error {
	token, err := sdk.RegisterService("model-svc", "Model Service", ms.router, ms.credentialStore)
	if err != nil {
		return err
	}
	ms.router = ms.router.SetBearerToken(token)

	return ms.registerPermissions()
}
