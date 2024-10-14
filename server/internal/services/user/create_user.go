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

	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/datastore"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
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
		datastore.Equals(datastore.Field("id"), user.Contacts[0]),
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

	roleIdAnys := []any{}
	for _, roleId := range roleIds {
		roleIdAnys = append(roleIdAnys, roleId)
	}

	roles, err := s.rolesStore.Query(
		datastore.IsInList(datastore.Field("id"), roleIdAnys...),
	).Find()
	if err != nil {
		return err
	}
	if len(roles) == 0 {
		return errors.New("no roles found")
	}
	if len(roles) < len(roleIds) {
		return errors.New("some roles are not found")
	}

	user.PasswordHash = passwordHash
	if user.Id == "" {
		user.Id = sdk.Id("usr")
	}

	now := time.Now()
	user.UpdatedAt = now
	user.CreatedAt = now

	err = s.usersStore.Create(user)
	if err != nil {
		return nil
	}

	for _, roleId := range roleIds {
		err = s.addRoleToUser(user.Id, roleId)
		if err != nil {
			return err
		}
	}

	return nil
}
