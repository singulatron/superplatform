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
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/singulatron/singulatron/localtron/lib"
)

func (d *DockerService) pullImage(imageName string) error {
	d.imagePullGlobalMutex.Lock()

	imageMutex, exists := d.imagePullMutexes[imageName]
	if !exists {
		imageMutex = &sync.Mutex{}
		d.imagePullMutexes[imageName] = imageMutex
	}

	d.imagePullGlobalMutex.Unlock()

	imageMutex.Lock()
	defer imageMutex.Unlock()

	images, err := d.client.ImageList(context.Background(), image.ListOptions{
		All: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to list Docker images")
	}

	imageExists := false
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == imageName || tag == fmt.Sprintf("%v:latest", imageName) {
				imageExists = true
				break
			}
		}
		if imageExists {
			break
		}
	}

	if imageExists {
		return nil
	}

	lib.Logger.Info("Starting to pull image", slog.String("image", imageName))

	err = pullImageWithProgress(d.client, imageName)
	if err != nil {
		lib.Logger.Error("Failed to pull image",
			slog.String("image", imageName),
			slog.String("error", err.Error()),
		)
		return err
	}

	lib.Logger.Debug("Pulling image is done", slog.String("image", imageName))

	return nil
}

type PullStatus struct {
	Status   string `json:"status"`
	Progress string `json:"progress"`
	ID       string `json:"id"`
}

func pullImageWithProgress(d *client.Client, imageName string) error {
	rc, err := d.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		return errors.Wrap(err, "failed to pull image")
	}
	defer func() {
		if err := rc.Close(); err != nil {
			lib.Logger.Error("Failed to close image pull response",
				slog.String("image", imageName),
				slog.String("error", err.Error()),
			)
		}
	}()

	decoder := json.NewDecoder(rc)
	for {
		var status PullStatus
		if err := decoder.Decode(&status); err == io.EOF {
			break
		} else if err != nil {
			lib.Logger.Error("Error pulling image",
				slog.String("error", err.Error()),
				slog.String("image", imageName),
			)
			return errors.Wrap(err, "Failed to decode image pull output")
		}
		logPullProgress(status)
	}

	return nil
}

func logPullProgress(status PullStatus) {
	if status.Progress != "" {
		lib.Logger.Info("Pulling image progress",
			slog.String("pullImageStatus", status.Status),
			slog.String("pullImageProgress", status.Progress),
			slog.String("imageId", status.ID),
		)
	} else {
		lib.Logger.Info("Pulling image",
			slog.String("pullImageStatus", status.Status),
			slog.String("id", status.ID),
		)
	}
}
