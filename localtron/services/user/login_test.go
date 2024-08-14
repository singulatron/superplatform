package userservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/singulatron/singulatron/localtron/di"
	sdk "github.com/singulatron/singulatron/localtron/sdk/go"

	clients "github.com/singulatron/singulatron/clients/go"
)

func TestRegistration(t *testing.T) {
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
	userSvc := client.UserSvcAPI

	adminLoginRsp, _, err := userSvc.Login(context.Background()).Request(clients.UserSvcLoginRequest{
		Slug:     clients.PtrString("singulatron"),
		Password: clients.PtrString("changeme"),
	}).Execute()
	require.NoError(t, err)

	publicKeyRsp, _, err := userSvc.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	adminClient := clients.NewAPIClient(&clients.Configuration{
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

	t.Run("user password change", func(t *testing.T) {
		claim, err := sdk.DecodeJWT(*adminLoginRsp.Token.Token, *publicKeyRsp.PublicKey)
		require.NoError(t, err)

		byTokenRsp, _, err := userSvc.ReadUserByToken(context.Background()).Body(clients.UserSvcReadUserByTokenRequest{
			Token: adminLoginRsp.Token.Token,
		}).Execute()
		require.NoError(t, err)

		require.Equal(t, "singulatron", *byTokenRsp.User.Slug)
		require.True(t, nil == byTokenRsp.User.PasswordHash)

		require.Equal(t, &claim.UserId, byTokenRsp.User.Id)

		changePassReq := clients.UserSvcChangePasswordRequest{
			Slug:            clients.PtrString("singulatron"),
			CurrentPassword: clients.PtrString("changeme"),
			NewPassword:     clients.PtrString("yo"),
		}
		_, _, err = userSvc.ChangePassword(context.Background()).Request(changePassReq).Execute()
		require.Error(t, err)

		_, _, err = adminClient.UserSvcAPI.ChangePassword(context.Background()).Request(changePassReq).Execute()
		require.NoError(t, err)

		// changing with wrong password should error
		changePassReq.CurrentPassword = clients.PtrString("yoWRONG")
		changePassReq.NewPassword = clients.PtrString("yo1")

		_, _, err = adminClient.UserSvcAPI.ChangePassword(context.Background()).Request(changePassReq).Execute()
		require.Error(t, err)
	})
}

func TestOrganization(t *testing.T) {
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
	userSvc := client.UserSvcAPI

	adminLoginRsp, _, err := userSvc.Login(context.Background()).Request(clients.UserSvcLoginRequest{
		Slug:     clients.PtrString("singulatron"),
		Password: clients.PtrString("changeme"),
	}).Execute()
	require.NoError(t, err)

	publicKeyRsp, _, err := userSvc.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	adminClient := clients.NewAPIClient(&clients.Configuration{
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

	// Register additional users
	_, _, err = userSvc.Register(context.Background()).Body(clients.UserSvcRegisterRequest{
		Name:     clients.PtrString("Some User"),
		Slug:     clients.PtrString("someotheruser"),
		Password: clients.PtrString("pw1233"),
	}).Execute()

	require.NoError(t, err)
	otherLoginRsp, _, err := userSvc.Login(context.Background()).Request(clients.UserSvcLoginRequest{
		Slug:     clients.PtrString("someotheruser"),
		Password: clients.PtrString("pw1233"),
	}).Execute()
	require.NoError(t, err)

	otherClient := clients.NewAPIClient(&clients.Configuration{
		Servers: clients.ServerConfigurations{
			{
				URL:         server.URL,
				Description: "Default server",
			},
		},
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + *otherLoginRsp.Token.Token,
		},
	})

	_, _, err = userSvc.Register(context.Background()).Body(clients.UserSvcRegisterRequest{
		Name:     clients.PtrString("Some Other User"),
		Slug:     clients.PtrString("someotherthirduser"),
		Password: clients.PtrString("thirdpw1233"),
	}).Execute()
	require.NoError(t, err)
	thirdLoginRsp, _, err := userSvc.Login(context.Background()).Request(clients.UserSvcLoginRequest{
		Slug:     clients.PtrString("someotherthirduser"),
		Password: clients.PtrString("thirdpw1233"),
	}).Execute()
	require.NoError(t, err)

	thirdClient := clients.NewAPIClient(&clients.Configuration{
		Servers: clients.ServerConfigurations{
			{
				URL:         server.URL,
				Description: "Default server",
			},
		},
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + *thirdLoginRsp.Token.Token,
		},
	})

	t.Run("claim contains new organization admin role after creating organization", func(t *testing.T) {
		createOrgReq := clients.UserSvcCreateOrganizationRequest{
			Id:   clients.PtrString("torgid1"),
			Slug: clients.PtrString("test-org"),
			Name: clients.PtrString("Test Org"),
		}
		_, _, err := adminClient.UserSvcAPI.CreateOrganization(context.Background()).Request(createOrgReq).Execute()
		require.NoError(t, err)

		claim, err := sdk.DecodeJWT(*adminLoginRsp.Token.Token, *publicKeyRsp.PublicKey)
		require.NoError(t, err)
		require.NotNil(t, claim)
		require.Equal(t, 1, len(claim.RoleIds), claim.RoleIds)

		loginReq := clients.UserSvcLoginRequest{
			Slug:     clients.PtrString("singulatron"),
			Password: clients.PtrString("changeme"),
		}
		loginRsp, _, err := adminClient.UserSvcAPI.Login(context.Background()).Request(loginReq).Execute()
		require.NoError(t, err)

		claim, err = sdk.DecodeJWT(*loginRsp.Token.Token, *publicKeyRsp.PublicKey)
		require.NoError(t, err)
		require.NotNil(t, claim)
		require.Equal(t, 2, len(claim.RoleIds), claim.RoleIds)
		require.Contains(t, claim.RoleIds, "user-svc:org:test-org:admin", claim.RoleIds)
	})

	t.Run("assign org to user", func(t *testing.T) {
		byTokenRsp, _, err := otherClient.UserSvcAPI.ReadUserByToken(context.Background()).Body(clients.UserSvcReadUserByTokenRequest{
			Token: otherLoginRsp.Token.Token,
		}).Execute()
		require.NoError(t, err)

		addUserReq := clients.UserSvcAddUserToOrganizationRequest{
			UserId: byTokenRsp.User.Id,
		}
		_, _, err = adminClient.UserSvcAPI.AddUserToOrganization(context.Background(), "torgid1").Request(addUserReq).Execute()
		require.NoError(t, err)

		loginReq := clients.UserSvcLoginRequest{
			Slug:     clients.PtrString("someotheruser"),
			Password: clients.PtrString("pw1233"),
		}
		// log in again and see the claim
		loginRsp, _, err := otherClient.UserSvcAPI.Login(context.Background()).Request(loginReq).Execute()
		require.NoError(t, err)

		claim, err := sdk.DecodeJWT(*loginRsp.Token.Token, *publicKeyRsp.PublicKey)
		require.NoError(t, err)
		require.NotNil(t, claim)
		require.Equal(t, 2, len(claim.RoleIds), claim.RoleIds)

		_, _, err = thirdClient.UserSvcAPI.RemoveUserFromOrganization(context.Background(), "torgid1", *byTokenRsp.User.Id).Execute()
		// third user cannot remove the second from the org of the first
		require.Error(t, err)

		_, _, err = adminClient.UserSvcAPI.RemoveUserFromOrganization(context.Background(), "torgid1", *byTokenRsp.User.Id).Execute()
		require.NoError(t, err)
	})
}
