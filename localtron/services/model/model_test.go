package modelservice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/singulatron/singulatron/localtron/di"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
)

func TestModel(t *testing.T) {
	univ, err := di.BigBang(di.UniverseOptions{
		Test: true,
	})
	require.NoError(t, err)
	ms := univ.ModelService
	cs := univ.ConfigService

	t.Run("get models", func(t *testing.T) {
		models, err := ms.GetModels()
		require.NoError(t, err)
		require.Equal(t, 1, len(models[0].Assets))
	})

	t.Run("model status", func(t *testing.T) {
		stat, err := ms.Status("huggingface/TheBloke/mistral-7b-instruct-v0.2.Q2_K.gguf")
		require.NoError(t, err)
		require.Equal(t, false, stat.Running)
		require.Equal(t, false, stat.AssetsReady)
		require.Equal(t, "127.0.0.1:8001", stat.Address)
	})

	t.Run("default", func(t *testing.T) {
		conf, err := cs.GetConfig()
		require.NoError(t, err)
		require.Equal(t, configservice.DefaultModelId, conf.Model.CurrentModelId)

		newModel := "huggingface/TheBloke/codellama-7b.Q3_K_M.gguf"
		err = ms.MakeDefault(newModel)
		// errors because it is not downloaded yet
		assert.Error(t, err)

		conf, err = cs.GetConfig()
		require.NoError(t, err)
		require.Equal(t, configservice.DefaultModelId, conf.Model.CurrentModelId)
	})
}
