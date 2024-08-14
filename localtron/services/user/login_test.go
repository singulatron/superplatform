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

//func TestOrganization(t *testing.T) {
//	hs := &di.HandlerSwitcher{}
//	server := httptest.NewServer(hs)
//	defer server.Close()
//
//	options := &di.Options{
//		Test: true,
//		Url:  server.URL,
//	}
//	universe, starterFunc, err := di.BigBang(options)
//	require.NoError(t, err)
//
//	hs.UpdateHandler(universe)
//	router := options.Router
//
//	err = starterFunc()
//	require.NoError(t, err)
//
//	token, err := sdk.RegisterUser(router, "someuser", "pw123", "Some name")
//	require.NoError(t, err)
//	loggedInRouter := router.SetBearerToken(token)
//
//	pkrsp := usertypes.GetPublicKeyResponse{}
//	err = router.Get(context.Background(), "user-svc", "/public-key", nil, &pkrsp)
//	require.NoError(t, err)
//
//	t.Run("claim contains new organization admin role after creating organization", func(t *testing.T) {
//		orgCreateRsp := &usertypes.CreateOrganizationResponse{}
//		err = loggedInRouter.Post(context.Background(), "user-svc", "/organization", usertypes.CreateOrganizationRequest{
//			Id:   "torgid1",
//			Slug: "test-org",
//			Name: "Test Org",
//		}, orgCreateRsp)
//		require.NoError(t, err)
//
//		claim, err := sdk.DecodeJWT(token, pkrsp.PublicKey)
//		require.NoError(t, err)
//		require.NotNil(t, claim)
//		require.Equal(t, 1, len(claim.RoleIds), claim.RoleIds)
//
//		req := usertypes.LoginRequest{
//			Slug:     "someuser",
//			Password: "pw123",
//		}
//		rsp := usertypes.LoginResponse{}
//		err = router.Post(context.Background(), "user-svc", "/login", req, &rsp)
//		require.NoError(t, err)
//
//		claim, err = sdk.DecodeJWT(rsp.Token.Token, pkrsp.PublicKey)
//		require.NoError(t, err)
//		require.NotNil(t, claim)
//		require.Equal(t, 2, len(claim.RoleIds), claim.RoleIds)
//		require.Contains(t, claim.RoleIds, "user-svc:org:test-org:admin", claim.RoleIds)
//	})
//
//	otherToken, err := sdk.RegisterUser(router, "someotheruser", "pw1233", "Some other name")
//	require.NoError(t, err)
//	otherRouter := router.SetBearerToken(otherToken)
//
//	thirdToken, err := sdk.RegisterUser(router, "someotherthirduser", "thirdpw1233", "Some third other name")
//	require.NoError(t, err)
//	thirdRouter := router.SetBearerToken(thirdToken)
//
//	t.Run("assign org to user", func(t *testing.T) {
//		claim, err := sdk.DecodeJWT(otherToken, pkrsp.PublicKey)
//		require.NoError(t, err)
//		require.NotNil(t, claim)
//		require.Equal(t, 1, len(claim.RoleIds), claim.RoleIds)
//
//		byTokenRsp := usertypes.ReadUserByTokenResponse{}
//		err = router.Post(context.Background(), "user-svc", "/user/by-token", usertypes.ReadUserByTokenRequest{
//			Token: otherToken,
//		}, &byTokenRsp)
//		require.NoError(t, err)
//
//		req := usertypes.AddUserToOrganizationRequest{
//			UserId: byTokenRsp.User.Id,
//		}
//		rsp := usertypes.AddUserToOrganizationResponse{}
//		err = loggedInRouter.Post(context.Background(), "user-svc", "/organization/torgid1/user", req, &rsp)
//		require.NoError(t, err)
//
//		loginReq := usertypes.LoginRequest{
//			Slug:     "someotheruser",
//			Password: "pw1233",
//		}
//		loginRsp := usertypes.LoginResponse{}
//		err = otherRouter.Post(context.Background(), "user-svc", "/login", loginReq, &loginRsp)
//		require.NoError(t, err)
//
//		claim, err = sdk.DecodeJWT(loginRsp.Token.Token, pkrsp.PublicKey)
//		require.NoError(t, err)
//		require.NotNil(t, claim)
//		require.Equal(t, 2, len(claim.RoleIds), claim.RoleIds)
//
//		ruReq := usertypes.RemoveUserFromOrganizationRequest{}
//		ruRsp := usertypes.RemoveUserFromOrganizationResponse{}
//		err = otherRouter.Delete(context.Background(), "user-svc", "/organization/torgid1/user/"+byTokenRsp.User.Id, ruReq, &ruRsp)
//		require.NoError(t, err)
//	})
//}
