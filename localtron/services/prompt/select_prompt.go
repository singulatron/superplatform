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

	"github.com/singulatron/singulatron/localtron/lib"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

var timeNow = time.Now

func selectPrompt(promptsMem *lib.MemoryStore[*prompttypes.Prompt]) *prompttypes.Prompt {
	var currentPrompt *prompttypes.Prompt

	promptsMem.ForeachStop(func(i int, prompt *prompttypes.Prompt) bool {
		if prompt.Status == prompttypes.PromptStatusAbandoned ||
			prompt.Status == prompttypes.PromptStatusCompleted ||
			prompt.Status == prompttypes.PromptStatusCanceled {
			return false
		}

		if prompt.RunCount == 0 ||
			timeNow().Sub(prompt.LastRun) >= baseDelay*time.Duration(1<<uint(prompt.RunCount-1)) {
			currentPrompt = prompt

			return true
		}
		return false
	})

	return currentPrompt
}
