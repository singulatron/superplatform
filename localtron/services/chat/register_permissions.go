/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package chatservice

import (
	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (p *ChatService) registerPermissions() error {
	for _, permission := range chattypes.ThreadPermissions {
		_, err := p.userService.UpsertPermission(
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
		usertypes.RoleUser,
	} {
		for _, permission := range chattypes.ThreadPermissions {
			p.userService.AddPermissionToRole(role.Id, permission.Id)
		}
	}

	for _, permission := range chattypes.MessagePermissions {
		_, err := p.userService.UpsertPermission(
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
		usertypes.RoleUser,
	} {
		for _, permission := range chattypes.MessagePermissions {
			p.userService.AddPermissionToRole(role.Id, permission.Id)
		}
	}

	return nil
}
