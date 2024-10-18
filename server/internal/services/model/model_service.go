/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package modelservice

import (
	"context"
	"sync"

	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/datastore"
	"github.com/singulatron/superplatform/sdk/go/lock"
	"github.com/singulatron/superplatform/sdk/go/router"

	modeltypes "github.com/singulatron/superplatform/server/internal/services/model/types"
)

type ModelService struct {
	modelStateMutex sync.Mutex
	modelPortMap    map[int]*modeltypes.ModelState

	router *router.Router
	lock   lock.DistributedLock

	modelsStore    datastore.DataStore
	platformsStore datastore.DataStore

	credentialStore datastore.DataStore

	gpuPlatform string
	llmHost     string
}

func NewModelService(
	gpuPlatform string,
	llmHost string,
	router *router.Router,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, insance any) (datastore.DataStore, error),
) (*ModelService, error) {
	srv := &ModelService{
		gpuPlatform:  gpuPlatform,
		router:       router,
		lock:         lock,
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
	ctx := context.Background()
	ms.lock.Acquire(ctx, "model-svc-start")
	defer ms.lock.Release(ctx, "model-svc-start")

	token, err := sdk.RegisterService("model-svc", "Model Service", ms.router, ms.credentialStore)
	if err != nil {
		return err
	}
	ms.router = ms.router.SetBearerToken(token)

	return ms.registerPermissions()
}
