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

type ListPromptOptions struct {
	CreatedAfter time.Time
	// or relationship
	Statuses     []prompttypes.PromptStatus
	LastRunAfter time.Time
}

func (p *PromptService) ListPrompts(options *ListPromptOptions) ([]*prompttypes.Prompt, error) {
	prompts := p.promptsMem.Filter(func(p *prompttypes.Prompt) bool {
		passes := true
		if len(options.Statuses) > 0 {
			statusMatches := false
			for _, v := range options.Statuses {
				if p.Status == v {
					statusMatches = true
				}
			}
			if !statusMatches {
				passes = false
			}
		}
		if !options.CreatedAfter.IsZero() && p.CreatedAt.After(options.CreatedAfter) {
			passes = false
		}
		if !options.LastRunAfter.IsZero() && p.LastRun.After(options.CreatedAfter) {
			passes = false
		}

		return passes
	})

	return prompts, nil
}
