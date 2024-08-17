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
	"log/slog"
	"net/url"
	"os"

	"github.com/pkg/errors"
	"github.com/singulatron/singulatron/localtron/logger"
	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
	dockertypes "github.com/singulatron/singulatron/localtron/services/docker/types"
	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

func (ms *ModelService) status(modelId string) (*modeltypes.ModelStatus, error) {
	// hostReq := dockertypes.GetDockerHostRequest{}
	hostRsp := dockertypes.GetDockerHostResponse{}
	err := ms.router.Get(context.Background(), "docker-svc", "/host", nil, &hostRsp)
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
		err := ms.router.Get(context.Background(), "config-svc", "/config", nil, &rsp)
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
		err := ms.router.Get(context.Background(), "download-svc", fmt.Sprintf("/download/%v", url.PathEscape(assetUrl)), nil, &rsp)
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

	hashRsp := dockertypes.ContainerIsRunningResponse{}
	err = ms.router.Get(context.Background(), "docker-svc", fmt.Sprintf("/container/%v/is-running", hash), nil, &hashRsp)
	if err != nil {
		logger.Warn("Checking if running error",
			slog.String("error", err.Error()),
		)
	}
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
