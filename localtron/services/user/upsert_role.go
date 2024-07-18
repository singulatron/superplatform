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
	"time"

	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) UpsertRole(id, name, description string, permissionIds []string) (*usertypes.Role, error) {
	permissions, err := s.permissionsStore.Query(
		datastore.Equal(datastore.Field("id"), permissionIds),
	).Find()
	if err != nil {
		return nil, err
	}

	if len(permissions) < len(permissionIds) {
		return nil, errors.New("nonexistent permissions")
	}

	roleI, found, err := s.rolesStore.Query(
		datastore.Equal(datastore.Field("id"), id),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		role := &usertypes.Role{
			Id:            id,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			Name:          name,
			Description:   description,
			PermissionIds: permissionIds,
		}
		return role, s.rolesStore.Create(role)
	}

	role := roleI.(*usertypes.Role)

	existingPermissionIdIndex := map[string]struct{}{}
	for _, permissionId := range role.PermissionIds {
		existingPermissionIdIndex[permissionId] = struct{}{}
	}

	for _, permissionId := range permissionIds {
		_, ok := existingPermissionIdIndex[permissionId]
		if !ok {
			role.PermissionIds = append(role.PermissionIds, permissionId)
		}
	}

	return role, s.rolesStore.Query(
		datastore.Equal(datastore.Field("id"), id),
	).Update(role)
}
