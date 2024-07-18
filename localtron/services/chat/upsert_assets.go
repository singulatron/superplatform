/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	"time"

	"github.com/singulatron/singulatron/localtron/datastore"
	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
)

func (a *ChatService) UpsertAssets(assets []*chattypes.Asset) error {
	now := time.Now()
	for _, v := range assets {
		if v.CreatedAt.IsZero() {
			v.CreatedAt = now
		}
	}

	assetIs := []datastore.Row{}
	for _, v := range assets {
		assetIs = append(assetIs, v)
	}
	return a.assetsStore.UpsertMany(assetIs)
}
