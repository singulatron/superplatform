/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dockerservice

import (
	"sync"

	"github.com/docker/docker/client"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	downloadservice "github.com/singulatron/singulatron/localtron/services/download"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

type DockerService struct {
	userService          *userservice.UserService
	configService        *configservice.ConfigService
	imagesCache          map[string]bool
	imagePullMutexes     map[string]*sync.Mutex
	imagePullGlobalMutex sync.Mutex
	launchModelMutex     sync.Mutex
	dockerHost           string
	dockerPort           int
	client               *client.Client
	mutex                sync.Mutex
	ds                   *downloadservice.DownloadService
}

func NewDockerService(
	downloadService *downloadservice.DownloadService,
	userService *userservice.UserService,
	configService *configservice.ConfigService,
) (*DockerService, error) {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	service := &DockerService{
		userService:   userService,
		configService: configService,
		ds:            downloadService,

		client:           c,
		imagePullMutexes: make(map[string]*sync.Mutex),
		imagesCache:      make(map[string]bool),
	}
	err = service.registerPermissions()
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (ds *DockerService) GetDockerHost() string {
	if ds.dockerHost == "" {
		return "127.0.0.1"
	}
	return ds.dockerHost
}

func (ds *DockerService) GetDockerPort() int {
	return ds.dockerPort
}
