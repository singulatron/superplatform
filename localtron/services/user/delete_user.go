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
	"fmt"

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

func (s *UserService) deleteUser(userId string) error {
	if userId == "" {
		return errors.New("no user id")
	}
	q := s.usersStore.Query(
		datastore.Id(userId),
	)
	_, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}

	isAdminUser, err := s.isAdmin(userId)
	if err != nil {
		return err
	}

	if isAdminUser {
		adminUsers, err := s.userRoleLinksStore.Query(
			datastore.Equal(datastore.Field("roleId"), usertypes.RoleAdmin.Id),
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

func (s *UserService) isAdmin(userId string) (bool, error) {
	_, isAdminUser, err := s.userRoleLinksStore.Query(
		datastore.Id(fmt.Sprintf("%v:%v", userId, usertypes.RoleAdmin.Id)),
	).FindOne()
	if err != nil {
		return false, err
	}

	return isAdminUser, nil
}
