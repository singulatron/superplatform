/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package prompttypes

import (
	"time"

	"github.com/singulatron/singulatron/localtron/clients/llm"
	"github.com/singulatron/singulatron/localtron/datastore"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SubscriberChan chan *llm.CompletionResponse

type PromptStatus string

const (
	PromptStatusScheduled PromptStatus = "scheduled"
	PromptStatusRunning   PromptStatus = "running"
	PromptStatusCompleted PromptStatus = "completed"
	// Errored means it will be still retried
	PromptStatusErrored   PromptStatus = "errored"
	PromptStatusAbandoned PromptStatus = "abandoned"
	PromptStatusCanceled  PromptStatus = "canceled"
)

type Prompt struct {
	PromptCreateFields

	// CreatedAt is the time of the prompt creation.
	CreatedAt time.Time `json:"createdAt"`

	// UpdatedAt is the last time the prompt was updated.
	UpdatedAt time.Time `json:"updatedAt"`

	// Status of the prompt.
	Status PromptStatus `json:"status,omitempty"`

	// LastRun is the time of the last prompt run.
	LastRun time.Time `json:"lastRun,omitempty"`

	// RunCount is the number of times the prompt was retried due to errors
	RunCount int `json:"runCount,omitempty"`

	// Error that arose during prompt execution, if any.
	Error string `json:"error,omitempty"`

	// UserId contains the ID of the user who submitted the prompt.
	UserId string `json:"userId"`
}

func (c *Prompt) GetId() string {
	return c.Id
}

func (c *Prompt) GetUpdatedAt() string {
	return c.Id
}

type PromptCreateFields struct {
	// Id is the unique ID of the prompt.
	Id string `json:"id"`

	// Prompt is the message itself eg. "What's a banana?
	Prompt string `json:"prompt" example:"What's a banana?" binding:"required"`

	// Sync drives whether prompt add request should wait and hang until
	// the prompt is done executing. By default the prompt just gets put on a queue
	// and the client will just subscribe to a Thread Stream.
	// For quick and dirty scripting however it's often times easier to do things syncronously.
	// In those cases set Sync to true.
	Sync bool `json:"sync"`

	// ThreadId is the ID of the thread a prompt belongs to.
	// Clients subscribe to Thread Streams to see the answer to a prompt,
	// or set `prompt.sync` to true for a blocking answer.
	ThreadId string `json:"threadId"`

	// Template of the prompt. Optional. If not present it's derived from ModelId.
	Template string `json:"template" example:"[INST]{prompt}[/INST]"`

	// ModelId is just the Singulatron internal ID of the model.
	ModelId string `json:"modelId,omitempty" example:"huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf"`

	// MaxRetries specified how many times the system should retry a prompt when it keeps erroring.
	MaxRetries int `json:"maxRetries,omitempty" example:"10"`
}

type AddPromptRequest struct {
	PromptCreateFields
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
	After   any       `json:"after,omitempty"`
	Count   int64     `json:"count"`
}

type RemovePromptRequest struct {
	PromptId string `json:"promptId"`
}

type RemovePromptResponse struct{}

type ListPromptOptions struct {
	Query *datastore.Query `json:"query"`
}
