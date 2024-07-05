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
	"errors"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) CreateRole(name, description string, permissionIds []string) (*usertypes.Role, error) {
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
		Id:            uuid.New().String(),
		Name:          name,
		Description:   description,
		PermissionIds: permissionIds,
	}

	return role, s.rolesStore.Create(role)
}
