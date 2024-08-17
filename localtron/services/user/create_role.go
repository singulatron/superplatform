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

	"github.com/google/uuid"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

func (s *UserService) createRole(ownerId, name, description string, permissionIds []string) (*usertypes.Role, error) {
	permissions, err := s.permissionsStore.Query(
		datastore.Equal(datastore.Field("id"), permissionIds),
	).Find()
	if err != nil {
		return nil, err
	}
	if len(permissions) < len(permissionIds) {
		return nil, errors.New("nonexistent permissions")
	}

	role := &usertypes.Role{
		Id:          uuid.New().String(),
		Name:        name,
		Description: description,
		OwnerId:     ownerId,
	}

	err = s.rolesStore.Upsert(role)
	if err != nil {
		return nil, err
	}

	for _, permissionId := range permissionIds {
		err = s.addPermissionToRole(ownerId, role.Id, permissionId)
		if err != nil {
			return nil, err
		}
	}

	return role, nil
}
