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
	"testing"
	"time"

	"github.com/singulatron/singulatron/localtron/datastore/localstore"

	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	"github.com/stretchr/testify/assert"
)

func TestSelectPrompt(t *testing.T) {
	fixedTime := time.Date(2023, 6, 1, 12, 0, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return fixedTime
	}

	tests := []struct {
		name           string
		prompts        []*prompttypes.Prompt
		expectedPrompt *prompttypes.Prompt
	}{
		{
			name:           "No prompts",
			prompts:        []*prompttypes.Prompt{},
			expectedPrompt: nil,
		},
		{
			name: "Prompt with RunCount 0",
			prompts: []*prompttypes.Prompt{
				{Status: prompttypes.PromptStatusScheduled, RunCount: 0},
			},
			expectedPrompt: &prompttypes.Prompt{Status: prompttypes.PromptStatusScheduled, RunCount: 0},
		},
		{
			name: "Prompt not due yet",
			prompts: []*prompttypes.Prompt{
				{Status: prompttypes.PromptStatusScheduled, RunCount: 1, LastRun: fixedTime.Add(-baseDelay / 2)},
			},
			expectedPrompt: nil,
		},
		{
			name: "Prompt due",
			prompts: []*prompttypes.Prompt{
				{Status: prompttypes.PromptStatusScheduled, RunCount: 1, LastRun: fixedTime.Add(-baseDelay)},
			},
			expectedPrompt: &prompttypes.Prompt{Status: prompttypes.PromptStatusScheduled, RunCount: 1, LastRun: fixedTime.Add(-baseDelay)},
		},
		{
			name: "Abandoned prompt",
			prompts: []*prompttypes.Prompt{
				{Status: prompttypes.PromptStatusAbandoned, RunCount: 0},
			},
			expectedPrompt: nil,
		},
		{
			name: "Completed prompt",
			prompts: []*prompttypes.Prompt{
				{Status: prompttypes.PromptStatusCompleted, RunCount: 0},
			},
			expectedPrompt: nil,
		},
		{
			name: "Canceled prompt",
			prompts: []*prompttypes.Prompt{
				{Status: prompttypes.PromptStatusCanceled, RunCount: 0},
			},
			expectedPrompt: nil,
		},
		{
			name: "Prompt with RunCount greater than 1, not due yet",
			prompts: []*prompttypes.Prompt{
				{Status: prompttypes.PromptStatusScheduled, RunCount: 2, LastRun: fixedTime.Add(-baseDelay)},
			},
			expectedPrompt: nil,
		},
		{
			name: "Prompt with RunCount greater than 1, due",
			prompts: []*prompttypes.Prompt{
				{Status: prompttypes.PromptStatusScheduled, RunCount: 2, LastRun: fixedTime.Add(-baseDelay * 2)},
			},
			expectedPrompt: &prompttypes.Prompt{Status: prompttypes.PromptStatusScheduled, RunCount: 2, LastRun: fixedTime.Add(-baseDelay * 2)},
		},
		{
			name: "Prompt with RunCount greater than 1, off by one",
			prompts: []*prompttypes.Prompt{
				{Status: prompttypes.PromptStatusScheduled, RunCount: 2, LastRun: fixedTime.Add(-baseDelay*2 + time.Second)},
			},
			expectedPrompt: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memStore := localstore.NewLocalStore[*prompttypes.Prompt]("")
			err := memStore.UpsertMany(tt.prompts)
			assert.NoError(t, err)
			actualPrompt, err := selectPrompt(memStore)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedPrompt, actualPrompt)
		})
	}
}
