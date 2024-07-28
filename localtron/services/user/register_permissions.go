/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (us *UserService) registerPermissions() error {
	for _, permission := range usertypes.UserPermissions {
		_, err := us.upsertPermission(
			permission.Id,
			permission.Name,
			permission.Description,
		)
		if err != nil {
			return err
		}
	}

	for _, role := range []*usertypes.Role{
		usertypes.RoleAdmin,
		// usertypes.RoleUser,
	} {
		for _, permission := range usertypes.UserPermissions {
			us.addPermissionToRole(role.Id, permission.Id)
		}
	}

	return nil
}
