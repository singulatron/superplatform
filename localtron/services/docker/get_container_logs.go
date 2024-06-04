package dockerservice

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/pkg/errors"
)

func (d *DockerService) GetContainerLogsAndStatus(modelURL string, logCount int) (string, error) {
	modelHash := generateStringHash(modelURL)

	ctx := context.Background()
	containers, err := d.client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return "", errors.Wrap(err, "error listing docker containers when getting logs")
	}

	for _, modelContainer := range containers {
		if modelContainer.Labels["singulatron-model-hash"] == modelHash {
			logOptions := container.LogsOptions{
				ShowStdout: true,
				ShowStderr: true,
				Tail:       fmt.Sprintf("%v", logCount),
			}
			logsReader, err := d.client.ContainerLogs(ctx, modelContainer.ID, logOptions)
			if err != nil {
				return "", errors.Wrap(err, "error getting container logs")
			}
			defer logsReader.Close()

			logs := new(strings.Builder)
			_, err = io.Copy(logs, logsReader)
			if err != nil {
				return "", errors.Wrap(err, "error reading container logs")
			}

			containerJSON, err := d.client.ContainerInspect(ctx, modelContainer.ID)
			if err != nil {
				return "", errors.Wrap(err, "error inspecting container")
			}

			portMappings := []string{}
			if containerJSON.NetworkSettings != nil {
				for port, bindings := range containerJSON.NetworkSettings.Ports {
					for _, binding := range bindings {
						portMappings = append(portMappings, fmt.Sprintf("%s:%s -> %s", binding.HostIP, binding.HostPort, port))
					}
				}
			} else {
				portMappings = append(portMappings, "unknown")
			}

			state := "unknown"
			healthStatus := "unknown"
			startedAt := "unkown"
			if containerJSON.State != nil {
				state = containerJSON.State.Status
				startedAt = containerJSON.State.StartedAt
				if containerJSON.State.Health != nil {
					healthStatus = containerJSON.State.Health.Status
				}
			}

			containerStatus := fmt.Sprintf(
				"ID: %s\nName: %s\nImage: %s\nState: %s\nStatus: %s\nCreated: %s\nStarted: %s\nPorts: %s\n",
				containerJSON.ID,
				containerJSON.Name,
				containerJSON.Image,
				state,
				healthStatus,
				containerJSON.Created,
				startedAt,
				strings.Join(portMappings, ", "),
			)

			return fmt.Sprintf("Container Status:\n%s\n\nContainer Logs:\n%s", containerStatus, logs.String()), nil
		}
	}

	return "", errors.New("no matching container found for the provided model URL")
}
