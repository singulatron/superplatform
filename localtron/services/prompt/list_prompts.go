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

import prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"

func (p *PromptService) ListPrompts() ([]*prompttypes.Prompt, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.currentPrompt == nil {
		return p.promptsToProcess, nil
	}
	return append([]*prompttypes.Prompt{p.currentPrompt}, p.promptsToProcess...), nil
}
