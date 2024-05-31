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
	"time"

	"github.com/google/uuid"
	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
)

func (a *AppService) AddChatMessage(chatMessage *apptypes.ChatMessage) error {
	if chatMessage.ThreadId == "" {
		return errors.New("empty chat message thread id")
	}
	if chatMessage.Id == "" {
		chatMessage.Id = uuid.New().String()
	}
	if chatMessage.Time == "" {
		chatMessage.Time = time.Now().Format(time.RFC3339)
	}
	var thread *apptypes.ChatThread
	for _, t := range a.chatFile.Threads {
		if t.Id == chatMessage.ThreadId {
			thread = t
		}
	}
	if thread == nil {
		// threads are created when a message is sent

		a.chatFile.AddThread(&apptypes.ChatThread{
			Id:   chatMessage.ThreadId,
			Time: time.Now().Format(time.RFC3339),
		})
	}

	alreadySaved := false

	a.chatFile.MessagesForeach(func(i int, v *apptypes.ChatMessage) {
		if v.Id == chatMessage.Id {
			alreadySaved = true
		}
	})

	if alreadySaved {
		return nil
	}

	a.chatFile.AddMessage(chatMessage)

	return a.saveChatFile()
}
