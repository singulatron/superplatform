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
package dockerservice

import (
	"sync"

	"github.com/docker/docker/client"
	downloadservice "github.com/singulatron/singulatron/localtron/services/download"
)

// DockerService manages Docker interactions
type DockerService struct {
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

// NewDockerService acts as a constructor for DockerService
func NewDockerService(downloadService *downloadservice.DownloadService) (*DockerService, error) {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &DockerService{
		client:           c,
		imagePullMutexes: make(map[string]*sync.Mutex),
		imagesCache:      make(map[string]bool),
		ds:               downloadService,
	}, nil
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
