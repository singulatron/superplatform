/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dockerservice

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/pkg/errors"
)

func (d *DockerService) hashIsRunning(hash string) (bool, error) {
	ctx := context.Background()
	containers, err := d.client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return false, errors.Wrap(err, "error listing docker containers when checking for runnign")
	}

	for _, container := range containers {
		if container.State != "running" {
			continue
		}
		if container.Labels["superplatform-hash"] == hash {
			return true, nil
		}
	}

	return false, nil
}
