/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

func (us *UserService) registerPermissions() error {
	for _, permission := range append(usertypes.UserPermissions, usertypes.AdminPermissions...) {
		_, err := us.upsertPermission(
			us.serviceUserId,
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
	} {
		for _, permission := range usertypes.AdminPermissions {
			err := us.addPermissionToRole(us.serviceUserId, role.Id, permission.Id)
			if err != nil {
				return err
			}
		}
	}

	for _, role := range []*usertypes.Role{
		usertypes.RoleUser,
	} {
		for _, permission := range usertypes.UserPermissions {
			err := us.addPermissionToRole(us.serviceUserId, role.Id, permission.Id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
