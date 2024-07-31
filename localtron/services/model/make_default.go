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

	"github.com/pkg/errors"
	"github.com/singulatron/singulatron/localtron/datastore"
	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
)

func (ms *ModelService) makeDefault(modelId string) error {
	stat, err := ms.status(modelId)
	if err != nil {
		return err
	}
	if !stat.AssetsReady {
		return fmt.Errorf("cannot set model as it is not downloaded yet")
	}

	rsp := configtypes.GetConfigResponse{}
	err = ms.router.Get(context.Background(), "config", "/get", nil, &rsp)
	if err != nil {
		return err
	}

	rsp.Config.Model.CurrentModelId = modelId

	return ms.router.Post(context.Background(), "config", "/save", &configtypes.SaveConfigRequest{
		Config: rsp.Config,
	}, rsp)
}

func (ms *ModelService) getPlatformByModelId(modelId string) (*modeltypes.Platform, error) {
	modelI, found, err := ms.modelsStore.Query(
		datastore.Id(modelId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("cannot find model")
	}
	model := modelI.(*modeltypes.Model)

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

	return platform, nil
}
