package userservice_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/singulatron/singulatron/localtron/di"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func TestRegistration(t *testing.T) {
	univ, err := di.BigBang(di.UniverseOptions{
		Test: true,
	})
	require.NoError(t, err)
	us := univ.UserService

	var token *usertypes.AuthToken
	var user *usertypes.User
	t.Run("user password change", func(t *testing.T) {
		var err error
		token, err = us.Login("singulatron", "changeme")
		require.NoError(t, err)

		user, err = us.ReadUserByToken(token.Token)
		require.NoError(t, err)

		require.Equal(t, "singulatron", user.Email)
		require.NotEqual(t, "changeme", user.PasswordHash)

		err = us.ChangePasswordAdmin("singulatron", "yo")
		require.NoError(t, err)

		token, err = us.Login("singulatron", "changeme")
		require.Error(t, err)

		token, err = us.Login("singulatron", "yo")
		require.NoError(t, err)

		err = us.ChangePassword("singulatron", "yo1", "yo1")
		require.Error(t, err)

		err = us.ChangePassword("singulatron", "yo", "yo1")
		require.NoError(t, err)

		token, err = us.Login("singulatron", "yo")
		require.Error(t, err)

		token, err = us.Login("singulatron", "yo1")
		require.NoError(t, err)
	})
}
