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
package apptypes

const EventChatMessageAddedName = "chatMessageAdded"

type EventChatMessageAdded struct {
	ThreadId string `json:"threadId"`
}

func (e EventChatMessageAdded) Name() string {
	return EventChatMessageAddedName
}

const EventChatThreadAddedName = "chatThreadAdded"

type EventChatThreadAdded struct {
	ThreadId string `json:"threadId"`
}

func (e EventChatThreadAdded) Name() string {
	return EventChatThreadAddedName
}

const EventChatThreadUpdateName = "chatThreadUpdate"

type EventChatThreadUpdate struct {
	ThreadId string `json:"threadId"`
}

func (e EventChatThreadUpdate) Name() string {
	return EventChatThreadUpdateName
}
