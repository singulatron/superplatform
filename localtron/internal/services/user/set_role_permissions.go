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

	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

func (s *UserService) overwriteRolePermissions(userId, roleId string, permissionIds []string) error {
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
	if role.OwnerId != userId {
		return fmt.Errorf("cannot add permission to unowned role")
	}

	perms, err := s.permissionsStore.Query(
		datastore.Equals(datastore.Field("id"), permissionIds),
	).Find()
	if err != nil {
		return err
	}
	if len(perms) < len(permissionIds) {
		return fmt.Errorf("cannot find some permissions")
	}

	err = s.userRoleLinksStore.Query(
		datastore.Equals(datastore.Field("roleId"), roleId),
	).Delete()
	if err != nil {
		return err
	}

	for _, permissionId := range permissionIds {
		err = s.addPermissionToRole(userId, roleId, permissionId)
		if err != nil {
			return err
		}
	}

	return nil
}
