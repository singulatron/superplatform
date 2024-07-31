/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package nodeservice

import (
	"context"

	nodetypes "github.com/singulatron/singulatron/localtron/services/node/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (ns *NodeService) registerPermissions() error {
	for _, permission := range nodetypes.NodePermissions {
		rsp := &usertypes.UpserPermissionResponse{}
		err := ns.router.Post(context.Background(), "user", "/upsert-permission", &usertypes.UpserPermissionRequest{
			Permission: &usertypes.Permission{
				Id:          permission.Id,
				Name:        permission.Name,
				Description: permission.Description,
			},
		}, rsp)
		if err != nil {
			return err
		}
	}

	for _, role := range []*usertypes.Role{
		usertypes.RoleAdmin,
		// nodetypes.RoleUser,
	} {
		for _, permission := range nodetypes.NodePermissions {
			rsp := &usertypes.AddPermissionToRoleResponse{}
			err := ns.router.Post(context.Background(), "user", "/add-permission-to-role", &usertypes.AddPermissionToRoleRequest{
				RoleId:       role.Id,
				PermissionId: permission.Id,
			}, rsp)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
