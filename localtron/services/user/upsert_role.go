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

func (s *UserService) UpsertRole(userId, id, name, description string, permissionIds []string) error {
	permissions, err := s.permissionsStore.Query(
		datastore.Equal(datastore.Field("id"), permissionIds),
	).Find()
	if err != nil {
		return err
	}

	if len(permissions) < len(permissionIds) {
		return errors.New("nonexistent permissions")
	}

	roleI, found, err := s.rolesStore.Query(
		datastore.Equal(datastore.Field("id"), id),
	).FindOne()
	if err != nil {
		return err
	}
	if !found {
		role := &usertypes.Role{
			Id:          id,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Name:        name,
			Description: description,
		}
		err = s.rolesStore.Create(role)
		if err != nil {
			return err
		}
	} else {
		err = s.rolesStore.Query(
			datastore.Equal(datastore.Field("id"), id),
		).Update(roleI)
		if err != nil {
			return err
		}
	}

	for _, permissionId := range permissionIds {
		err = s.addPermissionToRole(userId, roleI.GetId(), permissionId)
		if err != nil {
			return err
		}
	}

	return nil
}
