/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package appservice

import (
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/logger"

	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
)

func (a *AppService) AddChatMessage(chatMessage *apptypes.ChatMessage) error {
	if chatMessage.ThreadId == "" {
		return errors.New("empty chat message thread id")
	}
	if chatMessage.Id == "" {
		chatMessage.Id = uuid.New().String()
	}
	if chatMessage.CreatedAt == "" {
		chatMessage.CreatedAt = time.Now().Format(time.RFC3339)
	}

	threads, err := a.threadsStore.Query(
		datastore.Equal("id", chatMessage.ThreadId),
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

	a.firehoseService.Publish(apptypes.EventChatMessageAdded{
		ThreadId: chatMessage.ThreadId,
	})

	return a.messagesStore.Query(
		datastore.Equal("id", chatMessage.Id),
	).Upsert(chatMessage)
}
