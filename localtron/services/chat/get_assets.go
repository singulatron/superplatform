/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	"github.com/singulatron/singulatron/localtron/datastore"
	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
)

func (a *ChatService) getAssets(assetIds []string) ([]*chattypes.Asset, error) {
	assetIs, err := a.assetsStore.Query(
		datastore.Equal(datastore.Field("id"), assetIds),
	).Find()

	if err != nil {
		return nil, err
	}

	assets := []*chattypes.Asset{}
	for _, assetI := range assetIs {
		assets = append(assets, assetI.(*chattypes.Asset))
	}

	return assets, nil
}
