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

	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) DeleteUser(userId string) error {
	if userId == "" {
		return errors.New("no user id")
	}
	q := s.usersStore.Query(
		datastore.Id(userId),
	)
	user, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}

	isAdminUser := false
	for _, roleId := range user.RoleIds {
		if roleId == usertypes.RoleAdmin.Id {
			isAdminUser = true
		}
	}

	if isAdminUser {
		adminUsers, err := s.usersStore.Query(
			datastore.Equal("roleIds", usertypes.RoleAdmin.Id),
		).Find()
		if err != nil {
			return err
		}
		if len(adminUsers) == 0 {
			return errors.New("cannot detect number of admin users")
		}
		if len(adminUsers) == 1 {
			return errors.New("Cannot delete last admin user")
		}
	}

	return q.Delete()
}
