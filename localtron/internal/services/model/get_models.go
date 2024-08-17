/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package modelservice

import (
	modeltypes "github.com/singulatron/singulatron/localtron/internal/services/model/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

func (ms *ModelService) getModels() ([]*modeltypes.Model, error) {
	modelIs, err := ms.modelsStore.Query(
		datastore.All(),
	).Find()
	if err != nil {
		return nil, err
	}

	models := []*modeltypes.Model{}
	for _, modelI := range modelIs {
		models = append(models, modelI.(*modeltypes.Model))
	}

	return models, nil
}

func (ms *ModelService) getModel(modelId string) (*modeltypes.Model, bool, error) {
	modelIs, err := ms.modelsStore.Query(
		datastore.Id(modelId),
	).Find()
	if err != nil {
		return nil, false, err
	}

	models := []*modeltypes.Model{}
	for _, modelI := range modelIs {
		models = append(models, modelI.(*modeltypes.Model))
	}

	if len(models) == 0 {
		return nil, false, nil
	}

	return models[0], true, nil
}

func (ms *ModelService) getPlatform(platformId string) (*modeltypes.Platform, bool, error) {
	platformIs, err := ms.platformsStore.Query(
		datastore.Id(platformId),
	).Find()
	if err != nil {
		return nil, false, err
	}

	platforms := []*modeltypes.Platform{}
	for _, platformI := range platformIs {
		platforms = append(platforms, platformI.(*modeltypes.Platform))
	}

	return platforms[0], true, nil
}
