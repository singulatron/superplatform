/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dockerservice

import (
	"fmt"
	"net"
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

	credentialStore, err := datastoreFactory("dockerSvcCredentials", &sdk.Credential{})
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

func (ds *DockerService) getDockerHost() (string, error) {
	// Docker host should only exist for cases like WSL when the
	// Docker host address is not localhost.
	// Here instead of trying to return localhost we will try to find the docker bridge
	// ip so containers can address each other.
	if ds.dockerHost == "" {
		return getDockerBridgeIP()
	}
	return ds.dockerHost, nil
}

func (ds *DockerService) getDockerPort() int {
	return ds.dockerPort
}

func getDockerBridgeIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("failed to get network interfaces: %w", err)
	}

	for _, iface := range interfaces {
		if iface.Name == "docker0" {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", fmt.Errorf("failed to get addresses for interface %s: %w", iface.Name, err)
			}

			for _, addr := range addrs {
				ip, _, err := net.ParseCIDR(addr.String())
				if err != nil {
					continue
				}

				if ip.To4() != nil {
					return ip.String(), nil
				}
			}
		}
	}

	return "", fmt.Errorf("docker bridge interface not found")
}
