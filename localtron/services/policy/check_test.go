package policyservice_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	clients "github.com/singulatron/singulatron/clients/go"
	"github.com/singulatron/singulatron/localtron/di"
	policytypes "github.com/singulatron/singulatron/localtron/services/policy/types"
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
			TemplateId: string(policytypes.RateLimitPolicyTemplate.Id),
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
			_, rsp, err := policySvc.Check(context.Background()).Request(clients.PolicySvcCheckRequest{
				Endpoint: clients.PtrString("/test-endpoint"),
				Method:   clients.PtrString("GET"),
				Ip:       clients.PtrString("127.0.0.1"),
				UserId:   clients.PtrString("user-1"),
			}).Execute()
			require.NoError(t, err)
			require.Equal(t, 200, rsp.StatusCode)
		}
	})

	t.Run("exceeding the limit", func(t *testing.T) {
		_, rsp, err := policySvc.Check(context.Background()).Request(clients.PolicySvcCheckRequest{
			Endpoint: clients.PtrString("/test-endpoint"),
			Method:   clients.PtrString("GET"),
			Ip:       clients.PtrString("127.0.0.1"),
			UserId:   clients.PtrString("user-1"),
		}).Execute()
		require.NoError(t, err)
		require.Equal(t, 429, rsp.StatusCode) // Expecting a 429 Too Many Requests status
	})

	t.Run("reset after time window", func(t *testing.T) {
		time.Sleep(1 * time.Minute) // Wait for the time window to expire

		_, rsp, err := policySvc.Check(context.Background()).Request(clients.PolicySvcCheckRequest{
			Endpoint: clients.PtrString("/test-endpoint"),
			Method:   clients.PtrString("GET"),
			Ip:       clients.PtrString("127.0.0.1"),
			UserId:   clients.PtrString("user-1"),
		}).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, rsp.StatusCode) // Should be allowed again after the time window
	})
}
