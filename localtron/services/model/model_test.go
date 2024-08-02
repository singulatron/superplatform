package modelservice_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"net/url"
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
		getModelsRsp := modeltypes.ListResponse{}
		err = router.Get(context.Background(), "model-service", "/models", nil, &getModelsRsp)
		require.NoError(t, err)

		require.Equal(t, 1, len(getModelsRsp.Models[0].Assets))
	})

	t.Run("model status is not running, not ready", func(t *testing.T) {
		// statusReq := modeltypes.StatusRequest{}
		statusRsp := modeltypes.StatusResponse{}
		err = router.Get(context.Background(), "model-service", fmt.Sprintf("/model/%v/status", url.PathEscape("huggingface/TheBloke/mistral-7b-instruct-v0.2.Q2_K.gguf")), nil, &statusRsp)
		require.NoError(t, err)

		require.Equal(t, false, statusRsp.Status.Running)
		require.Equal(t, false, statusRsp.Status.AssetsReady)
		require.Equal(t, "127.0.0.1:8001", statusRsp.Status.Address)
	})

	t.Run("default", func(t *testing.T) {
		//getConfigReq := configtypes.GetConfigRequest{}
		getConfigRsp := configtypes.GetConfigResponse{}
		err = router.Get(context.Background(), "config-service", "/config", nil, &getConfigRsp)
		require.NoError(t, err)
		require.Equal(t, configservice.DefaultModelId, getConfigRsp.Config.Model.CurrentModelId)

		makeDefReq := modeltypes.MakeDefaultRequest{}
		makeDefRsp := modeltypes.MakeDefaultResponse{}
		err = router.Post(context.Background(), "model-service", fmt.Sprintf("/%v/make-default", url.PathEscape("huggingface/TheBloke/codellama-7b.Q3_K_M.gguf")), makeDefReq, &makeDefRsp)
		// errors because it is not downloaded yet
		require.Error(t, err)

		//getConfigReq = configtypes.GetConfigRequest{}
		getConfigRsp = configtypes.GetConfigResponse{}
		err = router.Get(context.Background(), "config-service", "/config", nil, &getConfigRsp)
		require.NoError(t, err)
		require.Equal(t, configservice.DefaultModelId, getConfigRsp.Config.Model.CurrentModelId)

	})
}
