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
	"fmt"

	prompttypes "github.com/singulatron/singulatron/localtron/internal/services/prompt/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

func (p *PromptService) registerPermissions() error {
	for _, permission := range prompttypes.PromptPermissions {
		rsp := &usertypes.UpserPermissionResponse{}
		err := p.router.Put(context.Background(), "user-svc", fmt.Sprintf("/permission/%v", permission.Id), &usertypes.UpserPermissionRequest{
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
		usertypes.RoleUser,
	} {
		for _, permission := range prompttypes.PromptPermissions {
			rsp := &usertypes.AddPermissionToRoleResponse{}
			err := p.router.Put(context.Background(), "user-svc",
				fmt.Sprintf("/role/%v/permission/%v", role.Id, permission.Id), &usertypes.AddPermissionToRoleRequest{}, rsp)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
