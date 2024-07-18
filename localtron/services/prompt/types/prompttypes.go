/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package prompttypes

import (
	"time"

	"github.com/singulatron/singulatron/localtron/datastore"
)

type PromptStatus string

const (
	PromptStatusScheduled PromptStatus = "scheduled"
	PromptStatusRunning   PromptStatus = "running"
	PromptStatusCompleted PromptStatus = "completed"
	// Errored means it will be still retried
	PromptStatusErrored   PromptStatus = "errored"
	PromptStatusAbandoned PromptStatus = "abandone"
	PromptStatusCanceled  PromptStatus = "canceled"
)

// Prompt
// @todo:
// - message and prompt have a lot of overlap, rethink
type Prompt struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	ThreadId string `json:"threadId"`
	UserId   string `json:"userId"`
	// Prompt is the message itself
	//    What's a banana?
	Prompt string `json:"prompt"`
	// Prompt template. Optional. Might be derived from ModelId
	//    [INST]{prompt}[/INST]
	Template string       `json:"template"`
	ModelId  string       `json:"modelId,omitempty"`
	Status   PromptStatus `json:"status,omitempty"`
	LastRun  time.Time    `json:"lastRun,omitempty"`
	// how many times this was ran
	// (retries are due to errors)
	RunCount   int    `json:"runCount,omitempty"`
	Error      string `json:"error,omitempty"`
	MaxRetries int    `json:"maxRetries,omitempty"`
	Sync       bool   `json:"sync"`
}

func (c *Prompt) GetId() string {
	return c.Id
}

func (c *Prompt) GetUpdatedAt() string {
	return c.Id
}

type AddPromptRequest struct {
	Prompt *Prompt `json:"prompt"`
}

type AddPromptResponse struct {
	Prompt *Prompt `json:"prompt"`
	Answer string  `json:"answer"`
}

type ListPromptsRequest struct {
	Query *datastore.Query `json:"query"`
}

type ListPromptsResponse struct {
	Prompts []*Prompt `json:"prompts"`
	After   time.Time `json:"after,omitempty"`
	Count   int64     `json:"count"`
}

type RemovePromptRequest struct {
	Prompt *Prompt `json:"prompt"`
}
