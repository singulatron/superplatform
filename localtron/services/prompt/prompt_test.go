package promptservice_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/singulatron/singulatron/localtron/di"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

func TestMessageCreatesThread(t *testing.T) {
	univ, err := di.BigBang(di.UniverseOptions{
		Test: true,
	})
	require.NoError(t, err)
	ps := univ.PromptService

	promptId := uuid.New().String()
	threadId := uuid.New().String()
	t.Run("add prompt", func(t *testing.T) {
		_, err := ps.AddPrompt(context.Background(), &prompttypes.Prompt{
			Id:       promptId,
			Prompt:   "hi",
			ThreadId: threadId,
		})
		require.NoError(t, err)
	})
}
