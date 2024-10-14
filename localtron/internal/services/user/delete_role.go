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

	"github.com/singulatron/superplatform/sdk/go/datastore"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

func (s *UserService) deleteRole(roleId string) error {
	q := s.rolesStore.Query(
		datastore.Id(roleId),
	)
	roleI, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}
	role := roleI.(*usertypes.Role)

	if role.Id == usertypes.RoleAdmin.Id {
		return errors.New("cannot delete default role")
	}

	return q.Delete()
}
