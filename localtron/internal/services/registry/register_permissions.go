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

func app(permSlices ...[]usertypes.Permission) []usertypes.Permission {
	ret := []usertypes.Permission{}
	for _, ps := range permSlices {
		ret = append(ret, ps...)
	}
	return ret
}

func (ns *RegistryService) registerPermissions() error {
	for _, permission := range app(
		registrytypes.NodeAdminPermissions,
		registrytypes.ServiceInstancePermissions,
		registrytypes.ServiceInstanceAdminPermissions,
		registrytypes.ServiceDefinitionPermissions,
		registrytypes.ServiceDefinitionAdminPermissions,
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
			registrytypes.NodeAdminPermissions,
			registrytypes.ServiceInstanceAdminPermissions,
			registrytypes.ServiceDefinitionAdminPermissions,
		) {
			rsp := &usertypes.AddPermissionToRoleResponse{}
			err := ns.router.Put(context.Background(), "user-svc",
				fmt.Sprintf("/role/%v/permission/%v", role.Id, permission.Id), &usertypes.AddPermissionToRoleRequest{}, rsp)
			if err != nil {
				return err
			}
		}
	}

	for _, role := range []*usertypes.Role{
		usertypes.RoleUser,
	} {
		for _, permission := range app(
			registrytypes.ServiceInstancePermissions,
			registrytypes.ServiceDefinitionPermissions,
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
