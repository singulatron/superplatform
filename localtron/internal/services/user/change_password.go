/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"errors"
	"time"

	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

func (s *UserService) changePassword(slug, currentPassword, newPassword string) error {
	q := s.usersStore.Query(
		datastore.Equals(datastore.Field("slug"), slug),
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

func (s *UserService) changePasswordAdmin(slug, newPassword string) error {
	q := s.usersStore.Query(
		datastore.Equals(datastore.Field("slug"), slug),
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
