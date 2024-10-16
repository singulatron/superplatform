/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package deployservice

import (
	"context"
	"fmt"

	deploytypes "github.com/singulatron/superplatform/server/internal/services/deploy/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

func app(permSlices ...[]usertypes.Permission) []usertypes.Permission {
	ret := []usertypes.Permission{}
	for _, ps := range permSlices {
		ret = append(ret, ps...)
	}
	return ret
}

func (ns *DeployService) registerPermissions() error {
	for _, permission := range app(
		deploytypes.DeployAdminPermissions,
	) {
		rsp := &usertypes.UpserPermissionResponse{}
		err := ns.router.Put(context.Background(), "user-svc", fmt.Sprintf("/permission/%v", permission.Id), &usertypes.UpserPermissionRequest{
			Permission: &usertypes.Permission{
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
	} {
		for _, permission := range app(
			deploytypes.DeployAdminPermissions,
		) {
			rsp := &usertypes.AddPermissionToRoleResponse{}
			err := ns.router.Put(context.Background(), "user-svc",
				fmt.Sprintf("/role/%v/permission/%v", role.Id, permission.Id), &usertypes.AddPermissionToRoleRequest{}, rsp)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
