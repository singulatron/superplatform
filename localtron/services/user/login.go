/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"time"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
	"golang.org/x/crypto/bcrypt"

	"github.com/pkg/errors"
)

func (s *UserService) login(slug, password string) (*usertypes.AuthToken, error) {
	userI, found, err := s.usersStore.Query(
		datastore.Equal(datastore.Field("slug"), slug),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("slug not found")
	}
	user := userI.(*usertypes.User)

	if !checkPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("unauthorized")
	}

	token, err := s.generateAuthToken(user)
	if err != nil {
		return nil, err
	}

	err = s.authTokensStore.Create(token)
	if err != nil {
		return nil, errors.Wrap(err, "error creating token")
	}

	return token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UserService) generateAuthToken(user *usertypes.User) (*usertypes.AuthToken, error) {
	roleLinks, err := s.userRoleLinksStore.Query(
		datastore.Equal(datastore.Field("userId"), user.Id),
	).Find()
	if err != nil {
		return nil, err
	}
	roleIds := []string{}
	for _, roleLink := range roleLinks {
		roleIds = append(roleIds, roleLink.(*usertypes.UserRoleLink).RoleId)
	}

	token, err := generateJWT(user, roleIds, s.privateKey)
	if err != nil {
		return nil, err
	}

	return &usertypes.AuthToken{
		Id:        uuid.New().String(),
		UserId:    user.Id,
		Token:     token,
		CreatedAt: time.Now(),
	}, nil
}
