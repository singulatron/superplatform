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

func (s *UserService) GetRoles() ([]*usertypes.Role, error) {
	rolesI, err := s.rolesStore.Query(
		datastore.All(),
	).OrderBy("name", false).Find()

	if err != nil {
		return nil, err
	}

	roles := []*usertypes.Role{}
	for _, roleI := range rolesI {
		roles = append(roles, roleI.(*usertypes.Role))
	}

	return roles, err
}
