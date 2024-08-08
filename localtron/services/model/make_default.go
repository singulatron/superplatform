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

	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
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
	err = ms.router.Get(context.Background(), "config-svc", "/config", nil, &rsp)
	if err != nil {
		return err
	}

	rsp.Config.Model.CurrentModelId = modelId

	return ms.router.Put(context.Background(), "config-svc", "/config", &configtypes.SaveConfigRequest{
		Config: rsp.Config,
	}, nil)
}
