/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"errors"
	"fmt"
	"time"

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

func (s *UserService) addPermissionToRole(userId, roleId, permissionId string) error {
	roleQ := s.rolesStore.Query(
		datastore.Id(roleId),
	)
	roleI, found, err := roleQ.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("cannot find role %v", roleId)
	}
	role := roleI.(*usertypes.Role)

	permQ := s.permissionsStore.Query(
		datastore.Id(permissionId),
	)
	permissionI, found, err := permQ.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("cannot find permission %v", permissionId)
	}
	permission := permissionI.(*usertypes.Permission)

	if permission.OwnerId != "" && permission.OwnerId != userId {
		return errors.New("not an owner of the permission")
	}

	return s.permissionRoleLinksStore.Upsert(&usertypes.PermissionRoleLink{
		Id:        fmt.Sprintf("%v:%v", permission.Id, role.Id),
		CreatedAt: time.Now(),

		RoleId:       roleId,
		PermissionId: permissionId,
	})
}
