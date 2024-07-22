/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package prompttypes

import (
	"context"

	"github.com/singulatron/singulatron/localtron/clients/llm"
)

type PromptServiceI interface {
	AddPrompt(ctx context.Context, prompt *Prompt) (*AddPromptResponse, error)
	ListPrompts(options *ListPromptOptions) ([]*Prompt, int64, error)
	Remove(prompt *Prompt) error
	Subscribe(threadId string, subscriber SubscriberChan)
	Unsubscribe(threadId string, subscriber SubscriberChan)
	Broadcast(threadId string, response *llm.CompletionResponse)
}
