/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	chattypes "github.com/singulatron/singulatron/localtron/internal/services/chat/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

func (a *ChatService) getThread(threadId string) (*chattypes.Thread, bool, error) {
	threadI, found, err := a.threadsStore.Query(
		datastore.Equal(datastore.Field("id"), threadId),
	).FindOne()
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	return threadI.(*chattypes.Thread), false, nil
}
