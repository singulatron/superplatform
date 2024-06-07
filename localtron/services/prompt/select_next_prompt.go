package promptservice

import (
	"time"

	"github.com/singulatron/singulatron/localtron/lib"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

func selectPrompt(promptsMem *lib.MemoryStore[*prompttypes.Prompt]) *prompttypes.Prompt {
	var currentPrompt *prompttypes.Prompt

	promptsMem.ForeachStop(func(i int, prompt *prompttypes.Prompt) bool {
		if prompt.Status == prompttypes.PromptStatusAbandoned ||
			prompt.Status == prompttypes.PromptStatusCompleted ||
			prompt.Status == prompttypes.PromptStatusCanceled {
			return false
		}

		if prompt.RunCount == 0 ||
			time.Since(prompt.LastRun) >= baseDelay*time.Duration(1<<uint(prompt.RunCount-1)) {
			currentPrompt = prompt

			return true
		}
		return false
	})

	return currentPrompt
}
