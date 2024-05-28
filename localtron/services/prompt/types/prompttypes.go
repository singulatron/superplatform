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

type Prompt struct {
	Id       string `json:"id"`
	ThreadId string `json:"threadId"`
	Prompt   string `json:"prompt"`
	ModelId  string `json:"modelId"`
}

type AddPromptRequest struct {
	Prompt Prompt `json:"prompt"`
}

type ListPromptsRequest struct{}

type ListPromptsResponse struct {
	Prompts []*Prompt `json:"prompts"`
}
