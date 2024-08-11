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

	// tk, err := usertypes.RegisterUser(router, "someuser", "pw123", "Some name")
	// require.NoError(t, err)
	// router = router.SetBearerToken(tk)

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
		require.Equal(t, claim.RoleIds, byTokenRsp.User.RoleIds)

		//err = router.Post(context.Background(), "user-svc", "/change-password-admin", usertypes.ChangePasswordAdminRequest{
		//	Email:       "singulatron",
		//	NewPassword: "yo",
		//}, nil)
		//require.NoError(t, err)
		//
		//loginReq := usertypes.LoginRequest{
		//	Email:    "singulatron",
		//	Password: "changeme",
		//}
		//err = router.Post(context.Background(), "user-svc", "/login", loginReq, &rsp)
		//require.Error(t, err)
		//
		//loginReq.Password = "yo"
		//err = router.Post(context.Background(), "user-svc", "/login", loginReq, &rsp)
		//require.NoError(t, err)

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
