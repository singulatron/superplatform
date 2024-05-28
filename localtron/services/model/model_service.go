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

	configservice "github.com/singulatron/singulatron/localtron/services/config"
	dockerservice "github.com/singulatron/singulatron/localtron/services/docker"
	downloadservice "github.com/singulatron/singulatron/localtron/services/download"

	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
)

type ModelService struct {
	modelStateMutex sync.Mutex
	modelStateMap   map[int]*modeltypes.ModelState

	downloadService *downloadservice.DownloadService
	confiService    *configservice.ConfigService
	dockerService   *dockerservice.DockerService
}

func NewModelService(
	ds *downloadservice.DownloadService,
	cs *configservice.ConfigService,
	dockerService *dockerservice.DockerService) (*ModelService, error) {
	srv := &ModelService{
		modelStateMap:   map[int]*modeltypes.ModelState{},
		downloadService: ds,
		confiService:    cs,
		dockerService:   dockerService,
	}

	return srv, nil
}
