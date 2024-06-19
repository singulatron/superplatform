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
	"github.com/singulatron/singulatron/localtron/datastore"
	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
)

func (ms *ModelService) Status(modelId string) (*modeltypes.ModelStatus, error) {
	dockerHost := ms.dockerService.GetDockerHost()
	singulatronLLMHost := os.Getenv("SINGULATRON_LLM_HOST")
	if singulatronLLMHost != "" {
		dockerHost = singulatronLLMHost
	}

	modelAddress := fmt.Sprintf("%v:%v", dockerHost, hostPortNum)

	if modelId == "" {
		conf, err := ms.configService.GetConfig()
		if err != nil {
			return nil, err
		}
		if conf.Model.CurrentModelId == "" {
			return nil, errors.New("cannot detect model id")
		}
		modelId = conf.Model.CurrentModelId
	}

	model, found, err := ms.modelsStore.Query(
		datastore.Id(modelId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("model not found")
	}

	for _, assetUrl := range model.Assets {
		downl, exists := ms.downloadService.GetDownload(assetUrl)
		if !exists || downl.Status != downloadtypes.DownloadStatusCompleted {
			return &modeltypes.ModelStatus{
				AssetsReady: false,
				Address:     modelAddress,
			}, nil
		}
	}

	isRunning := false
	if v, err := ms.dockerService.HashIsRunning(modelId); err == nil && v {
		isRunning = true
	}

	// @todo lock this
	if v, ok := ms.modelPortMap[hostPortNum]; ok {
		if !v.Answering {
			isRunning = false
		}
	}

	return &modeltypes.ModelStatus{
		Running:     isRunning,
		AssetsReady: true,
		Address:     modelAddress,
	}, nil
}
