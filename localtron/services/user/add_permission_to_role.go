/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"fmt"

	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) AddPermissionToRole(roleId, permissionId string) error {
	q := s.rolesStore.Query(
		datastore.Id(roleId),
	)
	roleI, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("cannot find role %v", roleId)
	}
	role := roleI.(*usertypes.Role)

	exists := false
	for _, v := range role.PermissionIds {
		if v == permissionId {
			exists = true
		}
	}

	if exists {
		return nil
	}

	_, found, err = s.permissionsStore.Query(
		datastore.Id(permissionId),
	).FindOne()
	if !found {
		return fmt.Errorf("cannot find permission %v", permissionId)
	}

	role.PermissionIds = append(role.PermissionIds, permissionId)

	return q.Update(role)
}
