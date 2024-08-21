/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package modelservice

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/logger"

	configtypes "github.com/singulatron/singulatron/localtron/internal/services/config/types"
	dockertypes "github.com/singulatron/singulatron/localtron/internal/services/docker/types"
	downloadtypes "github.com/singulatron/singulatron/localtron/internal/services/download/types"
	modeltypes "github.com/singulatron/singulatron/localtron/internal/services/model/types"
)

const hostPortNum = 8001

/*
Starts the model which has the supplied modelId or the currently activated one of
the modelId is empty.
*/
func (ms *ModelService) start(modelId string) error {
	var getConfigResponse *configtypes.GetConfigResponse

	if modelId == "" {
		rsp := configtypes.GetConfigResponse{}
		err := ms.router.Get(context.Background(), "config-svc", "/config", nil, &rsp)
		if err != nil {
			return err
		}
		getConfigResponse = &rsp
		conf := rsp.Config
		if conf.Model.CurrentModelId == "" {
			return errors.New("no model id specified and no default model")
		}
		modelId = conf.Model.CurrentModelId
	}

	modelI, found, err := ms.modelsStore.Query(
		datastore.Id(modelId),
	).FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("model not found")
	}
	model := modelI.(*modeltypes.Model)

	env := map[string]string{}
	for envarName, assetURL := range model.Assets {
		rsp := downloadtypes.GetDownloadResponse{}
		err := ms.router.Get(context.Background(), "download-svc", fmt.Sprintf("/download/%v", url.PathEscape(assetURL)), nil, &rsp)
		if err != nil {
			return err
		}
		if !rsp.Exists {
			return fmt.Errorf("asset with URL '%v' cannot be found locally", assetURL)
		}

		assetPath := *rsp.Download.FilePath
		assetPath = transformWinPaths(assetPath)

		env[envarName] = assetPath
	}

	platformI, found, err := ms.platformsStore.Query(
		datastore.Id(model.PlatformId),
	).FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("cannot find platform")
	}
	platform := platformI.(*modeltypes.Platform)

	launchOptions := &dockertypes.LaunchOptions{
		Name: platform.Id,
	}

	image := platform.Architectures.Default.Image
	port := platform.Architectures.Default.Port
	launchOptions.Envs = platform.Architectures.Default.Envars
	persistentPaths := platform.Architectures.Default.PersistentPaths

	switch os.Getenv("SINGULATRON_GPU_PLATFORM") {
	case "cuda":
		launchOptions.GPUEnabled = true
		if platform.Architectures.Cuda.Image != "" {
			image = platform.Architectures.Cuda.Image
		}
		if platform.Architectures.Cuda.Port != 0 {
			port = platform.Architectures.Cuda.Port
		}
		if len(platform.Architectures.Cuda.Envars) > 0 {
			launchOptions.Envs = platform.Architectures.Cuda.Envars
		}
		if len(platform.Architectures.Cuda.PersistentPaths) > 0 {
			persistentPaths = platform.Architectures.Cuda.PersistentPaths
		}
	}

	if getConfigResponse != nil {
		rsp := configtypes.GetConfigResponse{}
		err := ms.router.Get(context.Background(), "config-svc", "/config", nil, &rsp)
		if err != nil {
			return err
		}
		getConfigResponse = &rsp
	}

	configFolderPath := getConfigResponse.Config.Directory

	for envName, assetPath := range env {
		fileName := path.Base(assetPath)
		// eg. MODEL=/root/.singulatron/downloads/mistral-7b-instruct-v0.2.Q2_K.gguf
		launchOptions.Envs = append(launchOptions.Envs, fmt.Sprintf("%v=/root/.singulatron/downloads/%v", envName, fileName))
	}

	singulatronVolumeName := os.Getenv("SINGULATRON_VOLUME_NAME")
	if singulatronVolumeName == "" {
		if configFolderPath == "" {
			return errors.New("config folder not found")
		}
		singulatronVolumeName = configFolderPath
	}
	launchOptions.HostBinds = append(launchOptions.HostBinds, fmt.Sprintf("%v:/root/.singulatron", singulatronVolumeName))

	// Persistent paths are paths in the container we want to persist.
	// eg. /root/.cache/huggingface/diffusers
	// Then here we mount singulatron-data:/root/.cache/huggingface/diffusers
	for _, persistentPath := range persistentPaths {
		launchOptions.HostBinds = append(launchOptions.HostBinds,
			fmt.Sprintf("%v:%v", singulatronVolumeName, path.Dir(persistentPath)),
		)
	}

	hash, err := modelToHash(model, platform)
	if err != nil {
		return err
	}
	launchOptions.Hash = hash

	launchReq := &dockertypes.LaunchContainerRequest{
		Image:    image,
		Port:     port,
		HostPort: hostPortNum,
		Options:  launchOptions,
	}
	launchRsp := &dockertypes.LaunchContainerResponse{}
	err = ms.router.Put(context.Background(), "docker-svc", "/container", launchReq, &launchRsp)
	if err != nil {
		return errors.Wrap(err, "failed to launch container")
	}

	if launchRsp.Info.NewContainerStarted {
		state := ms.get(launchRsp.Info.PortNumber)
		if !state.HasCheckerRunning {
			go ms.checkIfAnswers(model, platform, launchRsp.Info.PortNumber, state)
		}
	}

	return nil
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

func (ms *ModelService) get(port int) *modeltypes.ModelState {
	ms.modelStateMutex.Lock()
	defer ms.modelStateMutex.Unlock()

	_, ok := ms.modelPortMap[port]
	if !ok {
		ms.modelPortMap[port] = &modeltypes.ModelState{}
	}

	return ms.modelPortMap[port]
}

func modelToHash(model *modeltypes.Model, platform *modeltypes.Platform) (string, error) {
	bs, err := json.Marshal(platform)
	if err != nil {
		return "", err
	}

	bs1, err := json.Marshal(model.Assets)
	if err != nil {
		return "", err
	}

	return generateStringHash(string(bs) + string(bs1)), nil
}

func generateStringHash(vals string) string {
	hasher := sha256.New()
	hasher.Write([]byte(vals))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (ms *ModelService) checkIfAnswers(
	model *modeltypes.Model,
	platform *modeltypes.Platform,
	port int,
	state *modeltypes.ModelState,
) {
	state.SetHasCheckerRunning(true)

	defer func() {
		state.SetHasCheckerRunning(false)
	}()

	hash, err := modelToHash(model, platform)
	if err != nil {
		logger.Error("cannot get hash to print logs", slog.Any("error", err))
		return
	}

	first := true
	for {
		if !first {
			time.Sleep(5 * time.Second)
		}
		first = false

		logger.Debug("Checking for answer started", slog.Int("port", port))

		hashRsp := dockertypes.ContainerIsRunningResponse{}
		err := ms.router.Get(context.Background(), "docker-svc", fmt.Sprintf("/container/%v/is-running", hash), nil, &hashRsp)
		if err != nil {
			logger.Warn("Model check error",
				slog.String("modelId", model.Id),
				slog.String("error", err.Error()),
			)
			continue
		}

		if !hashRsp.IsRunning {
			ms.printContainerLogs(model.Id, hash)
			continue
		}

		hostRsp := dockertypes.GetDockerHostResponse{}
		err = ms.router.Get(context.Background(), "docker-svc", "/host", nil, &hostRsp)
		if err != nil {
			logger.Warn("Docker host error",
				slog.String("error", err.Error()),
			)
			continue
		}
		dockerHost := hostRsp.Host

		singulatronLLMHost := os.Getenv("SINGULATRON_LLM_HOST")
		if singulatronLLMHost != "" {
			dockerHost = singulatronLLMHost
		}

		if !strings.HasPrefix(dockerHost, "http") {
			dockerHost = "http://" + dockerHost
		}

		host := strings.TrimPrefix(dockerHost, "http://")

		err = pingAddress(host, port)
		if err != nil {
			logger.Warn("Ping to LLM address failed",
				slog.String("address", host),
				slog.Int("port", port),
				slog.String("error", err.Error()),
			)
			state.SetAnswering(false)

			ms.printContainerLogs(model.Id, hash)
			continue
		}

		logger.Debug("LLM pinged successfully", slog.Int("port", port))
		state.SetAnswering(true)
		return
	}
}

func (ms *ModelService) printContainerLogs(modelId, hash string) {
	rsp := dockertypes.GetContainerSummaryResponse{}
	err := ms.router.Get(context.Background(), "docker-svc", fmt.Sprintf("/container/%v/summary/%v", hash, 10), nil, &rsp)
	if err != nil {
		logger.Warn("Error getting container logs",
			slog.String("modelId", modelId),
			slog.String("error", err.Error()),
		)
	} else {
		logger.Info("Container logs for model that is not running",
			slog.String("logs", rsp.Summary),
		)
	}
}

func pingAddress(host string, port int) error {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
