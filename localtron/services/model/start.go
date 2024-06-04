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
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/singulatron/singulatron/localtron/llm"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"

	"github.com/singulatron/singulatron/localtron/lib"
)

const portNum = 8001

/*
Starts the currently activated model
*/
func (ms *ModelService) Start(modelId string) error {
	stat, err := ms.Status(modelId)
	if err != nil {
		return err
	}
	if !stat.SelectedExists {
		return fmt.Errorf("cannot start selected model as it is not downloaded yet")
	}

	launchInfo, err := ms.dockerService.LaunchModel("the-singulatron", portNum, "crufter/llama-cpp-python-simple", stat.CurrentModelId)
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

		lib.Logger.Debug("Checking for answer started", slog.Int("port", port))

		isModelRunning, err := ms.dockerService.ModelIsRunning(modelId)
		if err != nil {
			lib.Logger.Warn("Model check error",
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
		if !strings.HasPrefix(dockerHost, "http") {
			dockerHost = "http://" + dockerHost
		}

		llmClient := llm.Client{
			LLMAddress: fmt.Sprintf("%v:%v", dockerHost, port),
		}

		rsp, err := llmClient.PostCompletions(llm.PostCompletionsRequest{
			MaxTokens: 32,
			Prompt:    "My name is John. Please say hello to me.",
		})
		if err != nil {
			lib.Logger.Debug("Answer failed for port", slog.Int("port", port), slog.String("error", err.Error()))
			state.SetAnswering(false)
			ms.printContainerLogs(modelId)
			continue
		}

		answer := ""
		for _, v := range rsp.Choices {
			answer += v.Text
		}

		if !strings.Contains(answer, "John") {
			lib.Logger.Debug("Answer failed to contain test sequence", slog.Int("port", port), slog.String("answer", answer))
			state.SetAnswering(false)
			continue
		} else {
			lib.Logger.Debug("LLM answered correctly", slog.Int("port", port))
			state.SetAnswering(true)
			return
		}
	}
}

func (ms *ModelService) printContainerLogs(modelId string) {
	logs, err := ms.dockerService.GetContainerLogsAndStatus(modelId, 100)
	if err != nil {
		lib.Logger.Warn("Error getting container logs",
			slog.String("modelId", modelId),
			slog.String("error", err.Error()),
		)
	} else {
		lib.Logger.Info("Container logs for model that is not running",
			slog.String("logs", logs),
		)
	}
}
