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

type MessagesMem struct {
	Messages []*ChatMessage `json:"messages"`
}

type Log struct {
	Level    string         `json:"level"`
	Time     string         `json:"time"`
	Source   string         `json:"source"`
	Ip       string         `json:"ip"`
	ClientId string         `json:"clientId"`
	Platform string         `json:"platform"`
	Message  string         `json:"message"`
	Fields   map[string]any `json:"fields"`
}

type LogRequest struct {
	Logs []Log `json:"logs"`
}

type LoggingStatus struct {
	Enabled bool `json:"enabled"`
}

type ChatThread struct {
	Id      string `json:"id"`
	TopicId string `json:"topicId,omitempty"`
	Title   string `json:"title"`
	Time    string `json:"time"`
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

	ti, err := parseTime(a[i].Time)
	if err != nil {
		lib.Logger.Error("Error parsing thread time",
			slog.String("threadId", a[i].Id),
			slog.String("error", err.Error()))
		return false // Could handle error differently if required
	}

	tj, err := parseTime(a[j].Time)
	if err != nil {
		lib.Logger.Error("Error parsing thread time",
			slog.String("threadId", a[j].Id),
			slog.String("error", err.Error()))
		return false // Could handle error differently if required
	}

	return ti.After(tj)
}

type ChatMessage struct {
	Id             string `json:"id"`
	ThreadId       string `json:"threadId"`
	MessageContent string `json:"messageContent"`
	IsUserMessage  bool   `json:"isUserMessage,omitempty"`
	// UserId is saved when the user is logged in to an account
	// @todo not used yet
	UserId string `json:"userId,omitempty"`
	Time   string `json:"time"`
}

type ByTime []*ChatMessage

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByTime) Less(i, j int) bool {
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

	ti, err := parseTime(a[i].Time)
	if err != nil {
		lib.Logger.Error("Error parsing message time",
			slog.String("messageId", a[i].Id),
			slog.String("error", err.Error()))
		return false // Could handle error differently if required
	}

	tj, err := parseTime(a[j].Time)
	if err != nil {
		lib.Logger.Error("Error parsing message time",
			slog.String("messageId", a[j].Id),
			slog.String("error", err.Error()))
		return false // Could handle error differently if required
	}

	return ti.Before(tj)
}

type AddChatMessageRequest struct {
	Message *ChatMessage `json:"message"`
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

type DeleteChatMessageRequest struct {
	MessageId string `json:"messageId"`
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

type GetChatMessagesRequest struct {
	ThreadId string `json:"threadId"`
}

type GetChatMessagesResponse struct {
	Messages []*ChatMessage `json:"messages"`
}

//
// Events
//

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
