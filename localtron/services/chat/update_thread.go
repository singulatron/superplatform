/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	"context"

	"github.com/singulatron/singulatron/localtron/datastore"
	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"
)

func (a *ChatService) updateThread(chatThread *chattypes.Thread) (*chattypes.Thread, error) {
	err := a.threadsStore.Query(
		datastore.Equal(datastore.Field("id"), chatThread.Id),
	).Update(chatThread)

	if err != nil {
		return nil, err
	}

	ev := chattypes.EventThreadUpdate{
		ThreadId: chatThread.Id,
	}
	a.router.Post(context.Background(), "firehose", "/publish", firehosetypes.PublishRequest{
		Event: &firehosetypes.Event{
			Name: ev.Name(),
			Data: ev,
		},
	}, nil)

	return chatThread, nil
}
