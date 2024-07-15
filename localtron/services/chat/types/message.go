/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chattypes

import (
	"time"
)

type Message struct {
	Id       string `json:"id"`
	ThreadId string `json:"threadId"`
	Content  string `json:"content"`
	// UserId is saved when the user is logged in to an account
	// @todo not used yet
	UserId    string    `json:"userId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	AssetIds  []string  `json:"assetIds,omitempty"`
}

type Asset struct {
	Id  string `json:"id"`
	Url string `json:"url,omitempty"`
	/* Some assets might have the content directly in them as base64
	encoded strings */
	Content    string    `json:"content,omitempty"`
	Type       string    `json:"type,omitempty"`
	Decription string    `json:"description,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

func (a Asset) GetId() string {
	return a.Id
}

func (c *Message) GetId() string {
	return c.Id
}

func (c *Message) GetUpdatedAt() string {
	return c.Id
}

type ByTime []*Message

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByTime) Less(i, j int) bool {
	ti := a[i].CreatedAt
	tj := a[j].CreatedAt

	return ti.Before(tj)
}

type AddMessageRequest struct {
	Message *Message `json:"message"`
}

type GetMessagesRequest struct {
	ThreadId string `json:"threadId"`
}

type GetMessagesResponse struct {
	Messages []*Message `json:"messages"`
	Assets   []*Asset   `json:"assets,omitempty"`
}

type DeleteMessageRequest struct {
	MessageId string `json:"messageId"`
}
