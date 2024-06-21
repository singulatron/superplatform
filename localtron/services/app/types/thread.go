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

import (
	"time"
)

type ChatThread struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	TopicIds []string `json:"topicIds,omitempty"`
	UserIds  []string `json:"userIds,omitempty"`

	Title string `json:"title"`
}

func (c *ChatThread) GetId() string {
	return c.Id
}

func (c *ChatThread) GetUpdatedAt() string {
	return c.Id
}

type ThreadByTime []*ChatThread

func (a ThreadByTime) Len() int      { return len(a) }
func (a ThreadByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ThreadByTime) Less(i, j int) bool {
	ti := a[i].CreatedAt
	tj := a[j].CreatedAt

	return ti.After(tj)
}

type AddChatThreadRequest struct {
	Thread *ChatThread `json:"thread"`
}

type UpdateChatThreadRequest struct {
	Thread *ChatThread `json:"thread"`
}

type AddChatThreadResponse struct {
	Thread *ChatThread `json:"thread"`
}

type DeleteChatThreadRequest struct {
	ThreadId string `json:"threadId"`
}

type GetChatThreadRequest struct {
	ThreadId string `json:"threadId"`
}

type GetChatThreadResponse struct {
	Thread ChatThread `json:"thread"`
}

type GetChatThreadsRequest struct{}

type GetChatThreadsResponse struct {
	Threads []*ChatThread `json:"threads"`
}
