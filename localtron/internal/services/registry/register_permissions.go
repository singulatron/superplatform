/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package registryservice

import (
	"context"
	"fmt"

	registrytypes "github.com/singulatron/singulatron/localtron/internal/services/registry/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

func (ns *RegistryService) registerPermissions() error {
	for _, permission := range registrytypes.RegistryPermissions {
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
		// registrytypes.RoleUser,
	} {
		for _, permission := range registrytypes.RegistryPermissions {
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
