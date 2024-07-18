/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package modelservice

import (
	"github.com/singulatron/singulatron/localtron/datastore"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
)

func (ms *ModelService) GetModels() ([]*modeltypes.Model, error) {
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
