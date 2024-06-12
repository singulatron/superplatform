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
	"os"

	"github.com/pkg/errors"
	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
)

func (ms *ModelService) Status(modelId string) (*modeltypes.Status, error) {
	if modelId == "" {
		conf, err := ms.confiService.GetConfig()
		if err != nil {
			return nil, err
		}
		modelId = conf.Model.CurrentModelId
	}

	if modelId == "" {
		return nil, errors.New("cannot locate current model id")
	}

	dockerHost := ms.dockerService.GetDockerHost()
	singulatronLLMHost := os.Getenv("SINGULATRON_LLM_HOST")
	if singulatronLLMHost != "" {
		dockerHost = singulatronLLMHost
	}

	modelAddress := fmt.Sprintf("%v:%v", dockerHost, portNum)

	downl, exists := ms.downloadService.GetDownload(modelId)
	if !exists || downl.Status != downloadtypes.DownloadStatusCompleted {
		return &modeltypes.Status{
			CurrentModelId: modelId,
			SelectedExists: false,
			ModelAddress:   modelAddress,
		}, nil
	}

	isRunning := false
	if v, err := ms.dockerService.ModelIsRunning(modelId); err == nil && v {
		isRunning = true
	}
	// @todo this is not threadsafe, needs locking, will panic
	if v, ok := ms.modelStateMap[portNum]; ok {
		if !v.Answering {
			isRunning = false
		}
	}

	return &modeltypes.Status{
		CurrentModelId: modelId,
		Running:        isRunning,
		SelectedExists: true,
		ModelAddress:   modelAddress,
	}, nil
}
