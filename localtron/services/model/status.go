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
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/singulatron/singulatron/localtron/datastore"
	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
	dockertypes "github.com/singulatron/singulatron/localtron/services/docker/types"
	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
)

func (ms *ModelService) status(modelId string) (*modeltypes.ModelStatus, error) {
	hostReq := dockertypes.GetDockerHostRequest{}
	hostRsp := dockertypes.GetDockerHostResponse{}
	err := ms.router.Post(context.Background(), "docker", "/host", hostReq, &hostRsp)
	if err != nil {
		return nil, err
	}

	dockerHost := hostRsp.Host
	singulatronLLMHost := os.Getenv("SINGULATRON_LLM_HOST")
	if singulatronLLMHost != "" {
		dockerHost = singulatronLLMHost
	}

	modelAddress := fmt.Sprintf("%v:%v", dockerHost, hostPortNum)

	if modelId == "" {
		rsp := configtypes.GetConfigResponse{}
		err := ms.router.Get(context.Background(), "config", "/get", nil, &rsp)
		if err != nil {
			return nil, err
		}

		if rsp.Config.Model.CurrentModelId == "" {
			return nil, errors.New("no model id specified and no default model")
		}
		modelId = rsp.Config.Model.CurrentModelId
	}

	modelI, found, err := ms.modelsStore.Query(
		datastore.Id(modelId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("model not found")
	}
	model := modelI.(*modeltypes.Model)

	for _, assetUrl := range model.Assets {
		rsp := downloadtypes.GetDownloadResponse{}
		err := ms.router.Post(context.Background(), "download", "/get", &downloadtypes.GetDownloadRequest{
			Url: assetUrl,
		}, &rsp)
		if err != nil {
			return nil, err
		}
		if !rsp.Exists || rsp.Download.Status != string(downloadtypes.DownloadStatusCompleted) {
			return &modeltypes.ModelStatus{
				AssetsReady: false,
				Address:     modelAddress,
			}, nil
		}
	}

	platformI, found, err := ms.platformsStore.Query(
		datastore.Id(model.PlatformId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("cannot find platform")
	}
	platform := platformI.(*modeltypes.Platform)

	hash, err := modelToHash(model, platform)
	if err != nil {
		return nil, err
	}

	isRunning := false
	hashReq := dockertypes.HashIsRunningRequest{
		Hash: hash,
	}
	hashRsp := dockertypes.HashIsRunningResponse{}
	err = ms.router.Post(context.Background(), "docker", "/hash-is-running", hashReq, &hashRsp)
	if err == nil && hashRsp.IsRunning {
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
