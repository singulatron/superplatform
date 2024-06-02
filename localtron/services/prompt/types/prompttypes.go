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
package prompttypes

// Prompt
// @todo:
// - message and prompt have a lot of overlap, rethink
type Prompt struct {
	Id       string `json:"id"`
	ThreadId string `json:"threadId"`
	// Prompt is the full prompt including template as in
	//    [INST]What's a banana?[/INST]
	Prompt string `json:"prompt"`
	// Message is the prompt without the template wrapper as in
	//    What's a banana?
	Message          string `json:"message"`
	ModelId          string `json:"modelId"`
	Time             string `json:"time"`
	IsBeingProcessed bool   `json:"isBeingProcessed"`
}

type AddPromptRequest struct {
	Prompt Prompt `json:"prompt"`
}

type ListPromptsRequest struct{}

type ListPromptsResponse struct {
	Prompts []*Prompt `json:"prompts"`
}

//
// Events
//

const EventPromptAddedName = "promptAdded"

type EventPromptAdded struct {
	Prompt Prompt `json:"prompt"`
}

func (e EventPromptAdded) Name() string {
	return EventPromptAddedName
}

const EventPromptProcessingStartedName = "promptProcessingStarted"

type EventPromptProcessingStarted struct {
	PromptId string `json:"promptId"`
}

func (e EventPromptProcessingStarted) Name() string {
	return EventPromptProcessingStartedName
}

const EventPromptProcessingFinishedName = "promptProcessingStarted"

type EventPromptProcessingFinished struct {
	PromptId string `json:"promptId"`
	Error    string `json:"error"`
}

func (e EventPromptProcessingFinished) Name() string {
	return EventPromptProcessingFinishedName
}

const EventPromptListChangedName = "promptListChanged"

type EventPromptListChanged struct {
}

func (e EventPromptListChanged) Name() string {
	return EventPromptListChangedName
}
