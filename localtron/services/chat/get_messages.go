/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

func (a *ChatService) getMessages(threadId string) ([]*chattypes.Message, error) {
	messageIs, err := a.messagesStore.Query(
		datastore.Equal(datastore.Field("threadId"), threadId),
	).OrderBy(datastore.OrderByField("createdAt", false)).Find()
	if err != nil {
		return nil, err
	}

	messages := []*chattypes.Message{}
	for _, messageI := range messageIs {
		messages = append(messages, messageI.(*chattypes.Message))
	}

	return messages, nil
}
