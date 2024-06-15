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
	"math"
	"time"

	"github.com/singulatron/singulatron/localtron/datastore"

	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

var timeNow = time.Now

func selectPrompt(promptsMem datastore.DataStore[*prompttypes.Prompt]) (*prompttypes.Prompt, error) {
	prompts, err := promptsMem.Query(
		datastore.All(),
	).Find()
	if err != nil {
		return nil, err
	}

	for _, prompt := range prompts {
		if prompt.Status == prompttypes.PromptStatusAbandoned ||
			prompt.Status == prompttypes.PromptStatusCompleted ||
			prompt.Status == prompttypes.PromptStatusCanceled {
			return nil, nil
		}

		runCount := prompt.RunCount
		if prompt.RunCount == 0 {
			// otherwise backoff is 0s
			runCount = 1
		}
		cappedRunCount := math.Min(float64(runCount), 10)
		backoff := baseDelay * time.Duration(math.Pow(2, cappedRunCount-1))

		if prompt.RunCount == 0 ||
			timeNow().Sub(prompt.LastRun) >= backoff {
			return prompt, nil
		}
	}

	return nil, nil
}
