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
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"

	"github.com/singulatron/singulatron/localtron/logger"
)

type LaunchOptions struct {
	Name       string
	Envs       []string
	Labels     map[string]string
	HostBinds  []string
	GPUEnabled bool
	Hash       string
}

type LaunchInfo struct {
	NewContainerStarted bool
	PortNumber          int
}

/*
A low level method for launching containers running models.
For a higher level one use `ModelService.Start“.
*/
func (d *DockerService) LaunchContainer(image string, internalPort, hostPort int, options *LaunchOptions) (*LaunchInfo, error) {
	err := d.pullImage(image)
	if err != nil {
		return nil, errors.Wrap(err, "image pull failure")
	}

	d.launchModelMutex.Lock()
	defer d.launchModelMutex.Unlock()

	if options == nil {
		options = &LaunchOptions{}
	}
	if options.Name == "" {
		options.Name = "the-singulatron"
	}

	containerConfig := &container.Config{
		Image: image,
		Env:   options.Envs,
		ExposedPorts: nat.PortSet{
			nat.Port(fmt.Sprintf("%v/tcp", internalPort)): {},
		},
		Labels: map[string]string{},
	}
	hostConfig := &container.HostConfig{
		Binds: options.HostBinds,
		PortBindings: map[nat.Port][]nat.PortBinding{
			nat.Port(fmt.Sprintf("%v/tcp", internalPort)): {
				{
					HostPort: fmt.Sprintf("%v", hostPort),
				},
			},
		},
		Resources: container.Resources{
			DeviceRequests: []container.DeviceRequest{},
		},
	}

	if options.GPUEnabled {
		hostConfig.Resources.DeviceRequests = append(hostConfig.Resources.DeviceRequests, container.DeviceRequest{
			Capabilities: [][]string{
				{"gpu"},
			},
			Count: -1,
		})
	}

	ctx := context.Background()

	containers, err := d.client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return nil, errors.Wrap(err, "error listing docker containers when launching")
	}

	var existingContainer *types.Container
	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+options.Name || name == options.Name || strings.Contains(name, options.Name) {
				existingContainer = &container
				break
			}
		}
		if existingContainer != nil {
			break
		}
	}

	if existingContainer != nil {
		if existingContainer.State != "running" || existingContainer.Labels["singulatron-hash"] != options.Hash {
			logger.Debug("Container state is not running or hash is mismatched, removing...")
			if err := d.client.ContainerRemove(ctx, existingContainer.ID, container.RemoveOptions{Force: true}); err != nil {
				return nil, errors.Wrap(err, "error removing Docker container")
			}
		} else {
			return &LaunchInfo{
				NewContainerStarted: false,
				PortNumber:          hostPort,
			}, nil
		}
	}

	containerConfig.Labels["singulatron-hash"] = options.Hash

	createdContainer, err := d.client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, options.Name)
	if err != nil {
		return nil, errors.Wrap(err, "error creating Docker container")
	}

	if err := d.client.ContainerStart(ctx, createdContainer.ID, container.StartOptions{}); err != nil {
		return nil, errors.Wrap(err, "error starting Docker container")
	}

	return &LaunchInfo{
		NewContainerStarted: true,
		PortNumber:          hostPort,
	}, nil
}
