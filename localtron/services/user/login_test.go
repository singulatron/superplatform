package userservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/singulatron/singulatron/localtron/di"
	sdk "github.com/singulatron/singulatron/localtron/sdk/go"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
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
	router := options.Router

	err = starterFunc()
	require.NoError(t, err)

	t.Run("user password change", func(t *testing.T) {
		req := usertypes.LoginRequest{
			Slug:     "singulatron",
			Password: "changeme",
		}
		rsp := usertypes.LoginResponse{}
		err := router.Post(context.Background(), "user-svc", "/login", req, &rsp)
		require.NoError(t, err)

		pkrsp := usertypes.GetPublicKeyResponse{}
		err = router.Get(context.Background(), "user-svc", "/public-key", nil, &pkrsp)
		require.NoError(t, err)

		claim, err := sdk.DecodeJWT(rsp.Token.Token, pkrsp.PublicKey)
		require.NoError(t, err)

		byTokenRsp := usertypes.ReadUserByTokenResponse{}
		err = router.Post(context.Background(), "user-svc", "/user/by-token", usertypes.ReadUserByTokenRequest{
			Token: rsp.Token.Token,
		}, &byTokenRsp)
		require.NoError(t, err)

		require.Equal(t, "singulatron", byTokenRsp.User.Slug)
		require.Equal(t, "", byTokenRsp.User.PasswordHash)

		require.Equal(t, claim.UserId, byTokenRsp.User.Id)

		changePassReq := usertypes.ChangePasswordRequest{
			Slug:            "singulatron",
			CurrentPassword: "changeme",
			NewPassword:     "yo",
		}
		// not logged in router should not be able to change pasword
		err = router.Post(context.Background(), "user-svc", "/change-password", changePassReq, nil)
		require.Error(t, err)

		loggedInRouter := router.SetBearerToken(rsp.Token.Token)

		err = loggedInRouter.Post(context.Background(), "user-svc", "/change-password", changePassReq, nil)
		require.NoError(t, err)

		// changing with wrong password should error
		changePassReq.CurrentPassword = "yoWRONG"
		changePassReq.NewPassword = "yo1"
		err = loggedInRouter.Post(context.Background(), "user-svc", "/change-password", changePassReq, nil)
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
	router := options.Router

	err = starterFunc()
	require.NoError(t, err)

	t.Run("claim contains new organization admin role after creating organization", func(t *testing.T) {
		req := usertypes.LoginRequest{
			Slug:     "singulatron",
			Password: "changeme",
		}
		rsp := usertypes.LoginResponse{}
		err := router.Post(context.Background(), "user-svc", "/login", req, &rsp)
		require.NoError(t, err)

		loggedInRouter := router.SetBearerToken(rsp.Token.Token)

		orgCreateRsp := &usertypes.CreateOrganizationResponse{}
		err = loggedInRouter.Post(context.Background(), "user-svc", "/organization", usertypes.CreateOrganizationRequest{
			Slug: "test-org",
			Name: "Test Org",
		}, orgCreateRsp)
		require.NoError(t, err)

		pkrsp := usertypes.GetPublicKeyResponse{}
		err = router.Get(context.Background(), "user-svc", "/public-key", nil, &pkrsp)
		require.NoError(t, err)

		claim, err := sdk.DecodeJWT(rsp.Token.Token, pkrsp.PublicKey)
		require.NoError(t, err)
		require.NotNil(t, claim)
		require.Equal(t, 1, len(claim.RoleIds), claim.RoleIds)

		err = router.Post(context.Background(), "user-svc", "/login", req, &rsp)
		require.NoError(t, err)

		claim, err = sdk.DecodeJWT(rsp.Token.Token, pkrsp.PublicKey)
		require.NoError(t, err)
		require.NotNil(t, claim)
		require.Equal(t, 2, len(claim.RoleIds), claim.RoleIds)
		require.Contains(t, claim.RoleIds, "user-svc:org:test-org:admin", claim.RoleIds)
	})
}
