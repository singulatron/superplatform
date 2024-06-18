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
	"fmt"

	"github.com/singulatron/singulatron/localtron/datastore"
)

func (s *UserService) SetRolePermissions(roleId string, permissionIds []string) error {
	q := s.rolesStore.Query(
		datastore.Id(roleId),
	)
	role, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("Cannot find role %v", roleId)
	}

	perms, err := s.permissionsStore.Query(
		datastore.Equal("id", permissionIds),
	).Find()
	if err != nil {
		return err
	}
	if len(perms) < len(permissionIds) {
		return fmt.Errorf("cannot find some permissions")
	}

	role.PermissionIds = permissionIds

	return q.Update(role)
}
