/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package userservice

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Register(email, password, name string, roleIds []string) (*usertypes.AuthToken, error) {
	passwordHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	_, alreadyExists, err := s.usersStore.Query(
		datastore.Equal("email", email),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if alreadyExists {
		return nil, errors.New("email already exists")
	}

	user := &usertypes.User{
		Id:           uuid.New().String(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		RoleIds:      roleIds,
	}

	err = s.usersStore.Create(user)
	if err != nil {
		return nil, err
	}

	token := generateAuthToken(user.Id)
	user.AuthTokenIds = append(user.AuthTokenIds, token.Id)

	return token, s.authTokensStore.Create(token)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
