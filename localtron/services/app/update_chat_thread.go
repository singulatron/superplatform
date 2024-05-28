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

	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
)

func (a *AppService) UpdateChatThread(chatThread *apptypes.ChatThread) (*apptypes.ChatThread, error) {
	if chatThread.Id == "" {
		return nil, errors.New("no thread id supplied to update")
	}
	found := false
	for i, v := range a.chatFile.Threads {
		if v.Id == chatThread.Id {
			old := a.chatFile.Threads[i]
			if chatThread.Time == "" {
				chatThread.Time = old.Time
			}
			a.chatFile.Threads[i] = chatThread
			found = true
		}
	}
	if !found {
		if chatThread.Time == "" {
			chatThread.Time = time.Now().Format(time.RFC3339)
		}
		a.chatFile.Threads = append(a.chatFile.Threads, chatThread)
	}

	return chatThread, a.saveChatFile()
}
