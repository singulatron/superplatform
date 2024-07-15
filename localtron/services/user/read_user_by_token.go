/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"errors"

	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) ReadUserByToken(token string) (*usertypes.User, error) {
	authTokenI, found, err := s.authTokensStore.Query(
		datastore.Equal(datastore.Field("token"), token),
	).FindOne()
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("token not found")
	}
	authToken := authTokenI.(*usertypes.AuthToken)

	userI, found, err := s.usersStore.Query(
		datastore.Equal(datastore.Field("id"), authToken.UserId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	ret := &usertypes.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return ret, nil
}
