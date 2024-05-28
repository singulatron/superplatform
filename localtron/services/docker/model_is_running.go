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

	"github.com/docker/docker/api/types/container"
	"github.com/pkg/errors"
)

func (d *DockerService) ModelIsRunning(modelURL string) (bool, error) {
	modelHash := generateStringHash(modelURL)

	ctx := context.Background()
	containers, err := d.client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return false, errors.Wrap(err, "error listing docker containers when checking for runnign")
	}

	for _, container := range containers {
		if container.State != "running" {
			continue
		}
		if container.Labels["singulatron-model-hash"] == modelHash {
			return true, nil
		}
	}

	return false, nil
}
