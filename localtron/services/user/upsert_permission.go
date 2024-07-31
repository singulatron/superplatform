/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) upsertPermission(id, name, description string) (*usertypes.Permission, error) {
	query := s.permissionsStore.Query(
		datastore.Equal(datastore.Field("id"), id),
	)

	permI, found, err := query.FindOne()
	if err != nil {
		return nil, err
	}

	if found {
		perm := permI.(*usertypes.Permission)

		perm.Name = name
		perm.Description = description
		query.Update(perm)
		return perm, nil
	}

	permission := &usertypes.Permission{
		Id:          id,
		Name:        name,
		Description: description,
	}

	s.permissionsStore.Create(permission)

	return permission, nil
}
