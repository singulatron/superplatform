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
	"github.com/google/uuid"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) CreateRole(name, description string, permissions []*usertypes.Permission) (*usertypes.Role, error) {
	s.runMutex.Lock()
	defer s.runMutex.Unlock()

	role := &usertypes.Role{
		Id:          uuid.New().String(),
		Name:        name,
		Description: description,
		Permissions: permissions,
	}

	s.rolesMem.Add(role)
	s.rolesFile.MarkChanged()

	return role, nil
}
