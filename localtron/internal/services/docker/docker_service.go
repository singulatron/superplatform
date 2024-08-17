/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dockerservice

import (
	"sync"

	"github.com/docker/docker/client"
	sdk "github.com/singulatron/singulatron/sdk/go"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/router"
)

type DockerService struct {
	router               *router.Router
	imagesCache          map[string]bool
	imagePullMutexes     map[string]*sync.Mutex
	imagePullGlobalMutex sync.Mutex
	launchModelMutex     sync.Mutex
	dockerHost           string
	dockerPort           int
	client               *client.Client
	mutex                sync.Mutex

	credentialStore datastore.DataStore
}

func NewDockerService(
	router *router.Router,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*DockerService, error) {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	credentialStore, err := datastoreFactory("dockerCredentials", &sdk.Credential{})
	if err != nil {
		return nil, err
	}

	service := &DockerService{
		router:          router,
		credentialStore: credentialStore,

		client:           c,
		imagePullMutexes: make(map[string]*sync.Mutex),
		imagesCache:      make(map[string]bool),
	}

	return service, nil
}

func (ds *DockerService) Start() error {
	token, err := sdk.RegisterService("docker-svc", "Docker Service", ds.router, ds.credentialStore)
	if err != nil {
		return err
	}
	ds.router = ds.router.SetBearerToken(token)

	return ds.registerPermissions()
}

func (ds *DockerService) getDockerHost() string {
	if ds.dockerHost == "" {
		return "127.0.0.1"
	}
	return ds.dockerHost
}

func (ds *DockerService) getDockerPort() int {
	return ds.dockerPort
}
