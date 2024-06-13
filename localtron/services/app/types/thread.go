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
	"log/slog"
	"time"

	"github.com/singulatron/singulatron/localtron/lib"
)

type ChatThread struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

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
	// Attempt to parse time using RFC3339 and then with JavaScript ISO8601 format
	parseTime := func(t string) (time.Time, error) {
		// First, try parsing in RFC 3339 format
		parsedTime, err := time.Parse(time.RFC3339, t)
		if err != nil {
			// If RFC 3339 fails, try parsing in a format that includes milliseconds (common in JavaScript)
			parsedTime, err = time.Parse("2006-01-02T15:04:05.999Z07:00", t)
		}
		return parsedTime, err
	}

	ti, err := parseTime(a[i].CreatedAt)
	if err != nil {
		lib.Logger.Error("Error parsing thread time",
			slog.String("threadId", a[i].Id),
			slog.String("error", err.Error()))
		return false // Could handle error differently if required
	}

	tj, err := parseTime(a[j].CreatedAt)
	if err != nil {
		lib.Logger.Error("Error parsing thread time",
			slog.String("threadId", a[j].Id),
			slog.String("error", err.Error()))
		return false // Could handle error differently if required
	}

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
