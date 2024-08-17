package policyservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	clients "github.com/singulatron/singulatron/clients/go"
	"github.com/singulatron/singulatron/localtron/di"
	policytypes "github.com/singulatron/singulatron/localtron/internal/services/policy/types"
)

func TestRateLimiting(t *testing.T) {
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

	err = starterFunc()
	require.NoError(t, err)

	client := clients.NewAPIClient(&clients.Configuration{
		Servers: clients.ServerConfigurations{
			{
				URL:         server.URL,
				Description: "Default server",
			},
		},
	})

	adminLoginRsp, _, err := client.UserSvcAPI.Login(context.Background()).Request(clients.UserSvcLoginRequest{
		Slug:     clients.PtrString("singulatron"),
		Password: clients.PtrString("changeme"),
	}).Execute()
	require.NoError(t, err)

	client = clients.NewAPIClient(&clients.Configuration{
		Servers: clients.ServerConfigurations{
			{
				URL:         server.URL,
				Description: "Default server",
			},
		},
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + *adminLoginRsp.Token.Token,
		},
	})

	policySvc := client.PolicySvcAPI
	instanceId := "instance-1"

	// Create a rate limit policy instance
	rateLimitReq := clients.PolicySvcUpsertInstanceRequest{
		Instance: &clients.PolicySvcInstance{
			Id:         &instanceId,
			Endpoint:   clients.PtrString("/test-endpoint"),
			TemplateId: clients.PolicySvcTemplateId(policytypes.RateLimitPolicyTemplate.Id),
			RateLimitParameters: &clients.PolicySvcRateLimitParameters{
				MaxRequests: clients.PtrInt32(5),
				TimeWindow:  clients.PtrString("1m"),
				Entity:      clients.EntityUserID.Ptr(),
				Scope:       clients.ScopeEndpoint.Ptr(),
			},
		},
	}
	_, _, err = policySvc.UpsertInstance(context.Background(), instanceId).Request(rateLimitReq).Execute()
	require.NoError(t, err)

	t.Run("allow up to the limit", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			checkRsp, _, err := policySvc.Check(context.Background()).Request(clients.PolicySvcCheckRequest{
				Endpoint: clients.PtrString("/test-endpoint"),
				Method:   clients.PtrString("GET"),
				Ip:       clients.PtrString("127.0.0.1"),
				UserId:   clients.PtrString("user-1"),
			}).Execute()
			require.NoError(t, err)
			require.Equal(t, true, checkRsp.Allowed)
		}
	})

	t.Run("exceeding the limit", func(t *testing.T) {
		checkRsp, _, err := policySvc.Check(context.Background()).Request(clients.PolicySvcCheckRequest{
			Endpoint: clients.PtrString("/test-endpoint"),
			Method:   clients.PtrString("GET"),
			Ip:       clients.PtrString("127.0.0.1"),
			UserId:   clients.PtrString("user-1"),
		}).Execute()
		require.NoError(t, err)
		require.Equal(t, false, checkRsp.Allowed)
	})
}
