package dockerservice

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/pkg/errors"
)

func (d *DockerService) GetContainerLogs(modelURL string, logCount int) (string, error) {
	modelHash := generateStringHash(modelURL)

	ctx := context.Background()
	containers, err := d.client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return "", errors.Wrap(err, "error listing docker containers when getting logs")
	}

	for _, modelContainer := range containers {
		if modelContainer.Labels["singulatron-model-hash"] == modelHash {
			options := container.LogsOptions{
				ShowStdout: true,
				ShowStderr: true,
				Tail:       fmt.Sprintf("%v", logCount),
			}
			out, err := d.client.ContainerLogs(ctx, modelContainer.ID, options)
			if err != nil {
				return "", errors.Wrap(err, "error getting container logs")
			}
			defer out.Close()

			logs := new(strings.Builder)
			_, err = io.Copy(logs, out)
			if err != nil {
				return "", errors.Wrap(err, "error reading container logs")
			}

			return logs.String(), nil
		}
	}

	return "", errors.New("no matching container found for the provided model URL")
}
