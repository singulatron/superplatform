/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dockerservice

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"

	"github.com/singulatron/singulatron/sdk/go/logger"

	configtypes "github.com/singulatron/singulatron/localtron/internal/services/config/types"
	dockertypes "github.com/singulatron/singulatron/localtron/internal/services/docker/types"
	downloadtypes "github.com/singulatron/singulatron/localtron/internal/services/download/types"
)

/*
A low level method for launching containers running models.
For a higher level one use `ModelService.Start“.
*/
func (d *DockerService) launchContainer(image string, internalPort, hostPort int, options *dockertypes.LaunchContainerOptions) (*dockertypes.LaunchInfo, error) {
	err := d.pullImage(image)
	if err != nil {
		return nil, errors.Wrap(err, "image pull failure")
	}

	d.launchModelMutex.Lock()
	defer d.launchModelMutex.Unlock()

	if options == nil {
		options = &dockertypes.LaunchContainerOptions{}
	}
	if options.Name == "" {
		options.Name = "the-singulatron"
	}

	envs, hostBinds, err := d.additionalEnvsAndHostBinds(options.Assets, options.PersistentPaths)
	if err != nil {
		return nil, err
	}

	containerConfig := &container.Config{
		Image: image,
		Env:   append(options.Envs, envs...),
		ExposedPorts: nat.PortSet{
			nat.Port(fmt.Sprintf("%v/tcp", internalPort)): {},
		},
		Labels: map[string]string{},
	}
	hostConfig := &container.HostConfig{
		Binds: hostBinds,
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
			logs, _ := d.getContainerLogsAndStatus(options.Hash, 10)
			logger.Debug("Container state is not running or hash is mismatched, removing...",
				slog.String("containerLogs", logs),
			)

			if err := d.client.ContainerRemove(ctx, existingContainer.ID, container.RemoveOptions{Force: true}); err != nil {
				return nil, errors.Wrap(err, "error removing Docker container")
			}
		} else {
			return &dockertypes.LaunchInfo{
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

	return &dockertypes.LaunchInfo{
		NewContainerStarted: true,
		PortNumber:          hostPort,
	}, nil
}

func (d *DockerService) additionalEnvsAndHostBinds(assets map[string]string, persistentPaths []string) ([]string, []string, error) {
	// We turn the asset map (which is an envar name to file URL map)
	// eg. {"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf"}
	// to an envarNameToFilePath
	// eg. {"MODEL": "/var/lib/some/local/path.gguf"}
	envarNameToFilePath := map[string]string{}

	// We translate URLs in the assets map into local file paths
	// by asking the Download Svc where did it download the file(s).

	for envarName, assetURL := range assets {
		rsp := downloadtypes.GetDownloadResponse{}
		err := d.router.Get(context.Background(), "download-svc", fmt.Sprintf("/download/%v", url.PathEscape(assetURL)), nil, &rsp)
		if err != nil {
			return nil, nil, err
		}
		if !rsp.Exists {
			return nil, nil, fmt.Errorf("asset with URL '%v' cannot be found locally", assetURL)
		}

		assetPath := *rsp.Download.FilePath
		assetPath = transformWinPaths(assetPath)

		envarNameToFilePath[envarName] = assetPath
	}

	envs := []string{}

	for envName, assetPath := range envarNameToFilePath {
		fileName := path.Base(assetPath)
		// eg. MODEL=/root/.singulatron/downloads/mistral-7b-instruct-v0.2.Q2_K.gguf
		envs = append(envs, fmt.Sprintf("%v=/root/.singulatron/downloads/%v", envName, fileName))
	}

	// If the Singulatron daemon is running in Docker, we need to find the volume it mounted so we can share
	// the downloaded files with containers the Singulatron daemon starts.
	// If the Singulatron daemon is running directly on the host, we will just mount the ~/.singulatron folder in
	// the containers the Singulatron daemon starts.

	singulatronVolumeName := os.Getenv("SINGULATRON_VOLUME_NAME")
	if singulatronVolumeName == "" {
		if isRunningInDocker() {
			currentContainerId, err := getContainerID()
			if err != nil {
				return nil, nil, err
			}

			mountedVolume, err := d.getMountedVolume(currentContainerId, "/root/.singulatron")
			if err != nil {
				return nil, nil, err
			}

			singulatronVolumeName = mountedVolume
		} else {
			// If we are not running in Docker, we will ask the Config Svc about the config directory and we mount that.
			var getConfigResponse *configtypes.GetConfigResponse
			err := d.router.Get(context.Background(), "config-svc", "/config", nil, &getConfigResponse)
			if err != nil {
				return nil, nil, err
			}

			configFolderPath := getConfigResponse.Config.Directory
			if configFolderPath == "" {
				return nil, nil, errors.New("config folder not found")
			}

			singulatronVolumeName = configFolderPath
		}

	}

	hostBinds := []string{}

	hostBinds = append(hostBinds, fmt.Sprintf("%v:/root/.singulatron", singulatronVolumeName))

	// Persistent paths are paths in the container we want to persist.
	// eg. /root/.cache/huggingface/diffusers
	// Then here we mount singulatron-data:/root/.cache/huggingface/diffusers
	for _, persistentPath := range persistentPaths {
		hostBinds = append(hostBinds,
			fmt.Sprintf("%v:%v", singulatronVolumeName, path.Dir(persistentPath)),
		)
	}

	return envs, hostBinds, nil
}

func (d *DockerService) getMountedVolume(containerID, mountPoint string) (string, error) {
	container, err := d.client.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return "", err
	}

	for _, mount := range container.Mounts {
		if mount.Destination == mountPoint {
			return mount.Name, nil
		}
	}

	return "", fmt.Errorf("no volume mounted at %s", mountPoint)
}

func isRunningInDocker() bool {
	// if checkDockerSocket() {
	// 	return true
	// }

	if checkDockerenvFile() {
		return true
	}

	if checkContainerenvFile() {
		return true
	}

	if checkContainerEnvVars() {
		return true
	}

	//if checkMountInfoForDockerOrKubernetes() {
	//	return true
	//}

	if checkCgroupForDockerOrKubernetes() {
		return true
	}

	if checkPidNamespace() {
		return true
	}

	if checkHostname() {
		return true
	}

	return false
}

func checkDockerSocket() bool {
	if _, err := os.Stat("/var/run/docker.sock"); err == nil {
		return true
	}
	return false
}

func checkDockerenvFile() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}

func checkContainerenvFile() bool {
	if _, err := os.Stat("/.containerenv"); err == nil {
		return true
	}
	return false
}

func checkContainerEnvVars() bool {
	if os.Getenv("DOCKER_CONTAINER") != "" || os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		return true
	}
	return false
}

func checkMountInfoForDockerOrKubernetes() bool {
	file, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "/docker/") || strings.Contains(line, "/kubepods/") {
			return true
		}
	}
	return false
}

func checkCgroupForDockerOrKubernetes() bool {
	file, err := os.Open("/proc/1/cgroup")
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "docker") || strings.Contains(scanner.Text(), "kubepods") {
			return true
		}
	}
	return false
}

func checkPidNamespace() bool {
	pid1, err := os.Readlink("/proc/1/ns/pid")
	if err != nil {
		return false
	}

	self, err := os.Readlink("/proc/self/ns/pid")
	if err != nil {
		return false
	}

	return pid1 != self
}

func checkHostname() bool {
	hostname, err := os.Hostname()
	if err != nil {
		return false
	}

	return strings.HasPrefix(hostname, "docker-") || strings.HasPrefix(hostname, "container-")
}

func getContainerID() (string, error) {
	id, err := getContainerIDFromCgroup()
	if err == nil {
		return id, nil
	}

	id, err = getContainerIDFromHostname()
	if err == nil {
		return id, nil
	}

	id, err = getContainerIDFromEnv()
	if err == nil {
		return id, nil
	}

	return "", fmt.Errorf("could not find container ID")
}

func getContainerIDFromCgroup() (string, error) {
	file, err := os.Open("/proc/self/cgroup")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "/")
		if len(parts) > 1 {
			containerID := parts[len(parts)-1]
			if len(containerID) == 64 && isHex(containerID) {
				return containerID, nil
			}
		}
	}

	return "", fmt.Errorf("could not find container ID in /proc/self/cgroup")
}

func getContainerIDFromHostname() (string, error) {
	hostname, err := ioutil.ReadFile("/etc/hostname")
	if err != nil {
		return "", err
	}

	id := strings.TrimSpace(string(hostname))
	if len(id) >= 12 && len(id) <= 64 && isHex(id) {
		return id, nil
	}

	return "", fmt.Errorf("hostname does not contain a valid container ID")
}

func getContainerIDFromEnv() (string, error) {
	id := os.Getenv("HOSTNAME")
	if len(id) >= 12 && len(id) <= 64 && isHex(id) {
		return id, nil
	}
	return "", fmt.Errorf("environment variable HOSTNAME does not contain a valid container ID")
}

func isHex(s string) bool {
	for _, c := range s {
		if !strings.Contains("0123456789abcdef", strings.ToLower(string(c))) {
			return false
		}
	}
	return true
}

// transformWinPaths maps win paths to unix paths so WSL can understand it
// eg. C:\users -> /mnt/c/users
func transformWinPaths(modelDir string) string {
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
