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
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"

	"github.com/singulatron/singulatron/localtron/logger"
)

type LaunchInfo struct {
	NewContainerStarted bool
	PortNumber          int
}

/*
A low level method for launching containers running models.
For a higher level one use `ModelService.Start“.
*/
func (d *DockerService) LaunchModel(containerName string, hostPort int, image, modelURL string) (*LaunchInfo, error) {
	err := d.pullImage(image)
	if err != nil {
		return nil, errors.Wrap(err, "image pull failure")
	}

	d.launchModelMutex.Lock()
	defer d.launchModelMutex.Unlock()

	download, exists := d.ds.GetDownload(modelURL)
	if !exists {
		return nil, fmt.Errorf("model '%v' is cannot be found locally", modelURL)
	}

	modelPath := download.FilePath
	modelPath = transformModelDir(modelPath)

	// @todo this is a hack
	configFolderPath := d.configService.ConfigDirectory
	sfolder := os.Getenv("SINGULATRON_HOST_FOLDER")
	if sfolder != "" {
		modelPath = strings.Replace(modelPath, configFolderPath, sfolder, 1)
	}

	urlParts := strings.Split(download.URL, "/")
	filenameFromUrl := urlParts[len(urlParts)-1]

	containerConfig := &container.Config{
		Image: image,
		// ie. MODEL=/models/mistral-7b-instruct-v0.2.Q2_K.gguf
		Env: []string{
			fmt.Sprintf("MODEL=/models/%v", filenameFromUrl),
			// only applicable to nvidia but should not affect others?
			"NVIDIA_VISIBLE_DEVICES=all",
		},
		// @todo port 8000 here is llama cpp python specific
		ExposedPorts: nat.PortSet{
			nat.Port("8000/tcp"): {},
		},
		Labels: map[string]string{},
	}
	hostConfig := &container.HostConfig{
		Binds: []string{fmt.Sprintf("%v:/models/%v", modelPath, filenameFromUrl)},
		PortBindings: map[nat.Port][]nat.PortBinding{
			// @todo port 8000 here is llama cpp python specific
			nat.Port("8000/tcp"): {
				{
					HostPort: fmt.Sprintf("%v", hostPort),
				},
			},
		},
		Resources: container.Resources{
			DeviceRequests: []container.DeviceRequest{},
		},
	}

	gpuEnabled := os.Getenv("SINGULATRON_GPU_ENABLED")
	if gpuEnabled == "true" {
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

	containerJSON, err := json.Marshal(containerConfig)
	if err != nil {
		return nil, err
	}
	hostJSON, err := json.Marshal(hostConfig)
	if err != nil {
		return nil, err
	}
	hash := generateStringHash(string(containerJSON) + string(hostJSON))

	var existingContainer *types.Container
	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+containerName || name == containerName || strings.Contains(name, containerName) {
				existingContainer = &container
				break
			}
		}
		if existingContainer != nil {
			break
		}
	}

	if existingContainer != nil {
		if existingContainer.State != "running" || existingContainer.Labels["singulatron-hash"] != hash {
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

	containerConfig.Labels["singulatron-hash"] = hash
	containerConfig.Labels["singulatron-model-hash"] = generateStringHash(modelURL)

	createdContainer, err := d.client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, containerName)
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

func transformModelDir(modelDir string) string {
	parts := strings.SplitN(modelDir, "\\", 2)
	if len(parts) == 1 {
		return modelDir
	}

	driveRegex := regexp.MustCompile(`^([A-Z]):`)
	newFirstPart := driveRegex.ReplaceAllStringFunc(parts[0], func(match string) string {
		driveLetter := strings.ToLower(match[:1])
		return "/mnt/" + driveLetter
	})

	newModelDir := newFirstPart
	if len(parts) > 1 {
		newModelDir += "/" + strings.ReplaceAll(parts[1], "\\", "/")
	}

	return newModelDir
}

func generateStringHash(vals string) string {
	hasher := sha256.New()
	hasher.Write([]byte(vals))
	return hex.EncodeToString(hasher.Sum(nil))
}
