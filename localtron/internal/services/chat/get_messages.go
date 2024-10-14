/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	"github.com/singulatron/superplatform/sdk/go/datastore"
	chattypes "github.com/singulatron/superplatform/server/internal/services/chat/types"
)

func (a *ChatService) getMessages(threadId string) ([]*chattypes.Message, error) {
	messageIs, err := a.messagesStore.Query(
		datastore.Equals(datastore.Field("threadId"), threadId),
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
