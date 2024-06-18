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
package modelservice

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/singulatron/singulatron/localtron/llm"
	"github.com/singulatron/singulatron/localtron/logger"

	dockerservice "github.com/singulatron/singulatron/localtron/services/docker"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
)

const portNum = 8001

/*
Starts the currently activated model
*/
func (ms *ModelService) Start(platform *modeltypes.Platform, assets modeltypes.Assets) error {
	env := map[string]string{}
	for assetName, assetURL := range assets {
		download, exists := ms.downloadService.GetDownload(assetURL)
		if !exists {
			return fmt.Errorf("asset with URL '%v' is cannot be found locally", assetURL)
		}

		assetPath := download.FilePath
		assetPath = transformWinPaths(assetPath)

		env[assetName] = assetPath
	}

	launchOptions := &dockerservice.LaunchOptions{}

	configFolderPath := ms.configService.ConfigDirectory
	singulatronHostFolder := os.Getenv("SINGULATRON_HOST_FOLDER")
	for envName, assetPath := range env {
		if singulatronHostFolder != "" {
			assetPath = strings.Replace(assetPath, configFolderPath, sfolder, 1)
		}
		fileName := path.Base(assetPath)
		// ie. MODEL=/models/mistral-7b-instruct-v0.2.Q2_K.gguf
		launchOptions.Envs = append(launchOptions.Envs, fmt.Sprintf("%v:/assets/%v", envName, fileName))
		launchOptions.Binds = append(launchOptions.Binds, fmt.Sprintf("%v:/assets/%v", assetPath, fileName))
	}

	image := platform.Container.Images.Default
	gpuEnabled := os.Getenv("SINGULATRON_GPU_ENABLED")
	if gpuEnabled == "true" {
		launchOptions.GPUEnabled = true
		switch os.Getenv("SINGULATRON_GPU_PLATFORM") {
		case "cuda":
			image = platform.Container.Images.Cuda
		}

		// only applicable to nvidia but should not affect others?
		launchOptions.Env = append(launchOptions.Env, "NVIDIA_VISIBLE_DEVICES=all")
	}

	launchInfo, err := ms.dockerService.LaunchContainer(image, portNum, platform)
	if err != nil {
		return errors.Wrap(err, "failed to launch container")
	}

	if launchInfo.NewContainerStarted {
		state := ms.get(launchInfo.PortNumber)
		if !state.HasCheckerRunning {
			go ms.checkIfAnswers(stat.CurrentModelId, launchInfo.PortNumber, state)
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

	_, ok := ms.modelStateMap[port]
	if !ok {
		ms.modelStateMap[port] = &modeltypes.ModelState{}
	}

	return ms.modelStateMap[port]
}

func (ms *ModelService) checkIfAnswers(modelId string, port int, state *modeltypes.ModelState) {
	state.SetHasCheckerRunning(true)

	defer func() {
		state.SetHasCheckerRunning(false)
	}()

	first := true
	for {
		if !first {
			time.Sleep(5 * time.Second)
		}
		first = false

		logger.Debug("Checking for answer started", slog.Int("port", port))

		isModelRunning, err := ms.dockerService.ModelIsRunning(modelId)
		if err != nil {
			logger.Warn("Model check error",
				slog.String("modelId", modelId),
				slog.String("error", err.Error()),
			)
			continue
		}
		if !isModelRunning {
			ms.printContainerLogs(modelId)
			continue
		}

		dockerHost := ms.dockerService.GetDockerHost()

		// @todo document this
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
			ms.printContainerLogs(modelId)
			continue
		}

		llmClient := llm.Client{
			LLMAddress: fmt.Sprintf("%v:%v", dockerHost, port),
		}

		rsp, err := llmClient.PostCompletions(llm.PostCompletionsRequest{
			MaxTokens: 32,
			Prompt:    "My name is John. Please say hello to me.",
		})
		if err != nil {
			logger.Debug("Answer failed for port",
				slog.Int("port", port),
				slog.String("error", err.Error()),
			)
			state.SetAnswering(false)
			ms.printContainerLogs(modelId)
			continue
		}

		answer := ""
		for _, v := range rsp.Choices {
			answer += v.Text
		}

		if !strings.Contains(answer, "John") {
			logger.Debug("Answer failed to contain test sequence", slog.Int("port", port), slog.String("answer", answer))
			state.SetAnswering(false)
			continue
		} else {
			logger.Debug("LLM answered correctly", slog.Int("port", port))
			state.SetAnswering(true)
			return
		}
	}
}

func (ms *ModelService) printContainerLogs(modelId string) {
	logs, err := ms.dockerService.GetContainerLogsAndStatus(modelId, 100)
	if err != nil {
		logger.Warn("Error getting container logs",
			slog.String("modelId", modelId),
			slog.String("error", err.Error()),
		)
	} else {
		logger.Info("Container logs for model that is not running",
			slog.String("logs", logs),
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
