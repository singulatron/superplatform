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
package promptservice

import (
	"time"

	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

func (p *PromptService) AddPrompt(prompt *prompttypes.Prompt) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if prompt.Time == "" {
		prompt.Time = time.Now().Format(time.RFC3339)
	}
	p.promptsToProcess = append(p.promptsToProcess, prompt)
	p.TriggerPromptProcessing()
	return nil
}
