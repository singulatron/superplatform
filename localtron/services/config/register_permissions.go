/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package configservice

import (
	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// this is called in the Start not in constructor to avoid import cycles
func (p *ConfigService) registerPermissions() error {
	for _, permission := range configtypes.ConfigPermissions {
		_, err := p.UpsertPermission(
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
		for _, permission := range configtypes.ConfigPermissions {
			p.AddPermissionToRole(role.Id, permission.Id)
		}
	}

	return nil
}
