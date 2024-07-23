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

	"github.com/singulatron/singulatron/localtron/datastore"

	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
	dockertypes "github.com/singulatron/singulatron/localtron/services/docker/types"
	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"

	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
)

type ModelService struct {
	modelStateMutex sync.Mutex
	modelPortMap    map[int]*modeltypes.ModelState

	modelsStore    datastore.DataStore
	platformsStore datastore.DataStore

	userService     usertypes.UserServiceI
	downloadService downloadtypes.DownloadServiceI
	configService   configtypes.ConfigServiceI
	dockerService   dockertypes.DockerServiceI
}

func NewModelService(
	ds downloadtypes.DownloadServiceI,
	userService usertypes.UserServiceI,
	cs configtypes.ConfigServiceI,
	dockerService dockertypes.DockerServiceI,
	datastoreFactory func(tableName string, insance any) (datastore.DataStore, error),
) (*ModelService, error) {
	srv := &ModelService{
		userService:     userService,
		downloadService: ds,
		configService:   cs,
		dockerService:   dockerService,

		modelPortMap: map[int]*modeltypes.ModelState{},
	}
	modelStore, err := datastoreFactory("models", &modeltypes.Model{})
	if err != nil {
		return nil, err
	}
	srv.modelsStore = modelStore

	platformsStore, err := datastoreFactory("platforms", &modeltypes.Platform{})
	if err != nil {
		return nil, err
	}
	srv.platformsStore = platformsStore

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
