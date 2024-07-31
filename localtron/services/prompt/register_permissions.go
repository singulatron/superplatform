/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package promptservice

import (
	"context"

	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (p *PromptService) registerPermissions() error {
	for _, permission := range prompttypes.PromptPermissions {
		rsp := &usertypes.UpserPermissionResponse{}
		err := p.router.Post(context.Background(), "user", "/upsert-permission", &usertypes.UpserPermissionRequest{
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
		usertypes.RoleUser,
	} {
		for _, permission := range prompttypes.PromptPermissions {
			rsp := &usertypes.AddPermissionToRoleResponse{}
			err := p.router.Post(context.Background(), "user", "/add-permission-to-role", &usertypes.AddPermissionToRoleRequest{
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
