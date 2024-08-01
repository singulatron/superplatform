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
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/logger"

	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"
)

func (a *ChatService) addMessage(chatMessage *chattypes.Message) error {
	if chatMessage.ThreadId == "" {
		return errors.New("empty chat message thread id")
	}
	if chatMessage.Id == "" {
		chatMessage.Id = uuid.New().String()
	}
	if chatMessage.CreatedAt.IsZero() {
		chatMessage.CreatedAt = time.Now()
	}

	threads, err := a.threadsStore.Query(
		datastore.Equal(datastore.Field("id"), chatMessage.ThreadId),
	).Find()
	if err != nil {
		return err
	}

	if len(threads) == 0 {
		return errors.New("thread does not exist")
	}

	logger.Info("Saving chat message",
		slog.String("messageId", chatMessage.Id),
	)

	ev := chattypes.EventMessageAdded{
		ThreadId: chatMessage.ThreadId,
	}
	err = a.router.Post(context.Background(), "firehose-service", "/publish", firehosetypes.PublishRequest{
		Event: &firehosetypes.Event{
			Name: ev.Name(),
			Data: ev,
		},
	}, nil)
	if err != nil {
		logger.Error("Failed to publish: %v", err)
	}

	return a.messagesStore.Query(
		datastore.Equal(datastore.Field("id"), chatMessage.Id),
	).Upsert(chatMessage)
}
