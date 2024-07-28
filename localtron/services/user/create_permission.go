/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) createPermission(id, name, description string) (*usertypes.Permission, error) {
	permission := &usertypes.Permission{
		Id:          id,
		Name:        name,
		Description: description,
	}

	return permission, s.permissionsStore.Create(permission)
}
