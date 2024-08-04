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

func (s *UserService) login(email, password string) (*usertypes.AuthToken, error) {
	userI, found, err := s.usersStore.Query(
		datastore.Equal(datastore.Field("email"), email),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("unauthorized")
	}
	user := userI.(*usertypes.User)

	if !checkPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("unauthorized")
	}

	tokens, err := s.authTokensStore.Query(
		datastore.Equal(datastore.Field("userId"), user.Id),
	).OrderBy(datastore.OrderByField("createdAt", true)).Find()
	if err != nil {
		return nil, err
	}

	if len(tokens) > 0 {
		return tokens[0].(*usertypes.AuthToken), nil
	}

	token, err := s.generateAuthToken(user.Id)
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

func (s *UserService) generateAuthToken(userId string) (*usertypes.AuthToken, error) {
	token, err := generateJWT(userId, s.privateKey)
	if err != nil {
		return nil, err
	}

	return &usertypes.AuthToken{
		Id:        uuid.New().String(),
		UserId:    userId,
		Token:     token,
		CreatedAt: time.Now(),
	}, nil
}
