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

	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) ChangePassword(email, currentPassword, newPassword string) error {
	q := s.usersStore.Query(
		datastore.Equal(datastore.Field("email"), email),
	)
	userI, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	if !checkPasswordHash(currentPassword, user.PasswordHash) {
		return errors.New("current password is incorrect")
	}

	newPasswordHash, err := hashPassword(newPassword)
	if err != nil {
		return err
	}
	user.PasswordHash = newPasswordHash
	user.UpdatedAt = time.Now()

	return q.Update(user)
}

func (s *UserService) ChangePasswordAdmin(email, newPassword string) error {
	q := s.usersStore.Query(
		datastore.Equal(datastore.Field("email"), email),
	)
	userI, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	newPasswordHash, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	user.PasswordHash = newPasswordHash
	user.UpdatedAt = time.Now()

	return q.Update(user)
}
