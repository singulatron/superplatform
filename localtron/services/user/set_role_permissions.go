/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"fmt"

	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) SetRolePermissions(roleId string, permissionIds []string) error {
	q := s.rolesStore.Query(
		datastore.Id(roleId),
	)
	roleI, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("Cannot find role %v", roleId)
	}
	role := roleI.(*usertypes.Role)

	perms, err := s.permissionsStore.Query(
		datastore.Equal(datastore.Field("id"), permissionIds),
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
