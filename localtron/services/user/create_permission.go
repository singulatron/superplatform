/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package userservice

import (
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) CreatePermission(id, name, description string) (*usertypes.Permission, error) {
	permission := &usertypes.Permission{
		Id:          id,
		Name:        name,
		Description: description,
	}

	s.permissionsMem.Add(permission)
	s.permissionsFile.MarkChanged()

	return permission, nil
}
