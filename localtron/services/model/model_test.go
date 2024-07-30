package modelservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/singulatron/singulatron/localtron/di"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func TestModel(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}

	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)
	router := options.Router

	require.NoError(t, starterFunc())

	token, err := usertypes.RegisterUser(router, "someuser", "pw123", "Some name")
	require.NoError(t, err)
	router = router.SetBearerToken(token)

	t.Run("get models", func(t *testing.T) {
		getModelsReq := modeltypes.GetModelsRequest{}
		getModelsRsp := modeltypes.GetModelsResponse{}
		err = router.Post(context.Background(), "model", "/list", getModelsReq, &getModelsRsp)
		require.NoError(t, err)

		require.Equal(t, 1, len(getModelsRsp.Models[0].Assets))
	})

	t.Run("model status is not running, not ready", func(t *testing.T) {
		statusReq := modeltypes.StatusRequest{
			ModelId: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q2_K.gguf",
		}
		statusRsp := modeltypes.StatusResponse{}
		err = router.Post(context.Background(), "model", "/status", statusReq, &statusRsp)
		require.NoError(t, err)

		require.Equal(t, false, statusRsp.Status.Running)
		require.Equal(t, false, statusRsp.Status.AssetsReady)
		require.Equal(t, "127.0.0.1:8001", statusRsp.Status.Address)
	})

	t.Run("default", func(t *testing.T) {
		getConfigReq := configtypes.GetConfigRequest{}
		getConfigRsp := configtypes.GetConfigResponse{}
		err = router.Post(context.Background(), "config", "/get", getConfigReq, &getConfigRsp)
		require.NoError(t, err)
		require.Equal(t, configservice.DefaultModelId, getConfigRsp.Config.Model.CurrentModelId)

		makeDefReq := modeltypes.MakeDefaultRequest{
			Id: "huggingface/TheBloke/codellama-7b.Q3_K_M.gguf",
		}
		makeDefRsp := modeltypes.MakeDefaultResponse{}
		err = router.Post(context.Background(), "model", "/make-default", makeDefReq, &makeDefRsp)
		// errors because it is not downloaded yet
		require.Error(t, err)

		getConfigReq = configtypes.GetConfigRequest{}
		getConfigRsp = configtypes.GetConfigResponse{}
		err = router.Post(context.Background(), "config", "/get", getConfigReq, &getConfigRsp)
		require.NoError(t, err)
		require.Equal(t, configservice.DefaultModelId, getConfigRsp.Config.Model.CurrentModelId)

	})
}
