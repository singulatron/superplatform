/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package nodeservice

import (
	nodetypes "github.com/singulatron/singulatron/localtron/services/node/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (ns *NodeService) registerPermissions() error {
	for _, permission := range nodetypes.NodePermissions {
		_, err := ns.userService.UpsertPermission(
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
		// nodetypes.RoleUser,
	} {
		for _, permission := range nodetypes.NodePermissions {
			err := ns.userService.AddPermissionToRole(role.Id, permission.Id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
