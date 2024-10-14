package userservice_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/test"
	"github.com/singulatron/superplatform/server/internal/di"

	clients "github.com/singulatron/superplatform/clients/go"
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

	userSvc := test.Client(server.URL).UserSvcAPI

	adminClient, adminToken, err := test.AdminClient(server.URL)
	require.NoError(t, err)

	publicKeyRsp, _, err := userSvc.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	t.Run("user password change", func(t *testing.T) {
		claim, err := sdk.DecodeJWT(adminToken, *publicKeyRsp.PublicKey)
		require.NoError(t, err)

		byTokenRsp, _, err := adminClient.UserSvcAPI.ReadUserByToken(context.Background()).Execute()
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

	adminClient, adminToken, err := test.AdminClient(server.URL)
	require.NoError(t, err)

	manyClients, err := test.MakeClients(options.Router, 2)
	require.NoError(t, err)
	otherClient := manyClients[0]
	thirdClient := manyClients[1]

	publicKeyRsp, _, err := test.Client(server.URL).UserSvcAPI.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	orgId1 := sdk.Id("org")

	t.Run("claim contains new organization admin role after creating organization", func(t *testing.T) {
		createOrgReq := clients.UserSvcCreateOrganizationRequest{
			Id:   clients.PtrString(orgId1),
			Slug: clients.PtrString("test-org"),
			Name: clients.PtrString("Test Org"),
		}
		_, _, err := adminClient.UserSvcAPI.CreateOrganization(context.Background()).Request(createOrgReq).Execute()
		require.NoError(t, err)

		claim, err := sdk.DecodeJWT(adminToken, *publicKeyRsp.PublicKey)
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
		require.Contains(t, claim.RoleIds, fmt.Sprintf("user-svc:org:{%v}:admin", orgId1), claim.RoleIds)

		tokenRsp, _, err := adminClient.UserSvcAPI.ReadUserByToken(context.Background()).Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(tokenRsp.Organizations))
		require.Equal(t, clients.PtrString(orgId1), tokenRsp.ActiveOrganizationId)
	})

	t.Run("assign org to user", func(t *testing.T) {
		byTokenRsp, _, err := otherClient.UserSvcAPI.ReadUserByToken(context.Background()).Execute()
		require.NoError(t, err)

		addUserReq := clients.UserSvcAddUserToOrganizationRequest{
			UserId: byTokenRsp.User.Id,
		}
		_, _, err = adminClient.UserSvcAPI.AddUserToOrganization(context.Background(), orgId1).Request(addUserReq).Execute()
		require.NoError(t, err)

		loginReq := clients.UserSvcLoginRequest{
			Slug:     clients.PtrString("test-user-slug-0"),
			Password: clients.PtrString("testUserPassword0"),
		}
		// log in again and see the claim
		loginRsp, _, err := otherClient.UserSvcAPI.Login(context.Background()).Request(loginReq).Execute()
		require.NoError(t, err)

		claim, err := sdk.DecodeJWT(*loginRsp.Token.Token, *publicKeyRsp.PublicKey)

		require.NoError(t, err)
		require.NotNil(t, claim)
		require.Equal(t, 2, len(claim.RoleIds), claim.RoleIds)
		require.Contains(t, claim.RoleIds, fmt.Sprintf("user-svc:org:{%v}:user", orgId1))

		_, _, err = thirdClient.UserSvcAPI.RemoveUserFromOrganization(context.Background(), orgId1, *byTokenRsp.User.Id).Execute()
		// third user cannot remove the second from the org of the first
		require.Error(t, err)

		_, _, err = adminClient.UserSvcAPI.RemoveUserFromOrganization(context.Background(), orgId1, *byTokenRsp.User.Id).Execute()
		require.NoError(t, err)
	})
}
