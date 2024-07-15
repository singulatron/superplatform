/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	"github.com/singulatron/singulatron/localtron/datastore"
	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
)

func (a *ChatService) GetThreads(userId string) ([]*chattypes.Thread, error) {
	threadIs, err := a.threadsStore.Query(
		datastore.Equal(datastore.Field("userIds"), userId),
	).OrderBy("createdAt", true).Find()
	if err != nil {
		return nil, err
	}

	threads := []*chattypes.Thread{}
	for _, threadI := range threadIs {
		threads = append(threads, threadI.(*chattypes.Thread))
	}

	return threads, nil
}
