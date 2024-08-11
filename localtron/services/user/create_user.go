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

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) createUser(user *usertypes.User, password string, roleIds []string) error {
	if user.Name == "" {
		return errors.New("username missing")
	}
	if len(user.Contacts) == 0 {
		return errors.New("contact missing")
	}
	if len(user.Contacts) > 1 {
		return errors.New("more than one contact")
	}
	if password == "" {
		return errors.New("password missing")
	}

	_, contactExists, err := s.contactsStore.Query(
		datastore.Equal(datastore.Field("id"), user.Contacts[0]),
	).FindOne()
	if err != nil {
		return err
	}

	if contactExists {
		return errors.New("contact already exists")
	}

	passwordHash, err := hashPassword(password)
	if err != nil {
		return err
	}

	roles, err := s.rolesStore.Query(
		datastore.Equal(datastore.Field("id"), roleIds),
	).Find()
	if err != nil {
		return err
	}
	if len(roles) < len(roleIds) {
		return errors.New("some roles are not found")
	}

	user.PasswordHash = passwordHash
	user.RoleIds = roleIds
	if user.Id == "" {
		user.Id = uuid.NewString()
	}

	now := time.Now()
	user.UpdatedAt = now
	user.CreatedAt = now

	return s.usersStore.Create(user)
}
