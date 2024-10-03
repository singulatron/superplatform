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

	chattypes "github.com/singulatron/singulatron/localtron/internal/services/chat/types"
	firehosetypes "github.com/singulatron/singulatron/localtron/internal/services/firehose/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/logger"
)

func (a *ChatService) updateThread(chatThread *chattypes.Thread) (*chattypes.Thread, error) {
	err := a.threadsStore.Query(
		datastore.Equals(datastore.Field("id"), chatThread.Id),
	).Update(chatThread)

	if err != nil {
		return nil, err
	}

	ev := chattypes.EventThreadUpdate{
		ThreadId: chatThread.Id,
	}
	err = a.router.Post(context.Background(), "firehose-svc", "/event", firehosetypes.EventPublishRequest{
		Event: &firehosetypes.Event{
			Name: ev.Name(),
			Data: ev,
		},
	}, nil)
	if err != nil {
		logger.Error("Failed to publish: %v", err)
	}

	return chatThread, nil
}
