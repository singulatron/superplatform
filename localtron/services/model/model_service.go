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
package modelservice

import (
	"sync"

	"github.com/singulatron/singulatron/localtron/datastore"

	configservice "github.com/singulatron/singulatron/localtron/services/config"
	dockerservice "github.com/singulatron/singulatron/localtron/services/docker"
	downloadservice "github.com/singulatron/singulatron/localtron/services/download"
	storefactoryservice "github.com/singulatron/singulatron/localtron/services/store_factory"
	userservice "github.com/singulatron/singulatron/localtron/services/user"

	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
)

type ModelService struct {
	modelStateMutex sync.Mutex
	modelPortMap    map[int]*modeltypes.ModelState

	modelsStore datastore.DataStore[*modeltypes.Model]

	userService     *userservice.UserService
	downloadService *downloadservice.DownloadService
	configService   *configservice.ConfigService
	dockerService   *dockerservice.DockerService
}

func NewModelService(
	ds *downloadservice.DownloadService,
	userService *userservice.UserService,
	cs *configservice.ConfigService,
	dockerService *dockerservice.DockerService) (*ModelService, error) {
	srv := &ModelService{
		userService:     userService,
		downloadService: ds,
		configService:   cs,
		dockerService:   dockerService,

		modelPortMap: map[int]*modeltypes.ModelState{},
	}
	store, err := storefactoryservice.GetStore[*modeltypes.Model]("models")
	if err != nil {
		return nil, err
	}
	srv.modelsStore = store

	err = srv.registerPermissions()
	if err != nil {
		return nil, err
	}

	err = srv.bootstrapModels()
	if err != nil {
		return nil, err
	}

	return srv, nil
}
