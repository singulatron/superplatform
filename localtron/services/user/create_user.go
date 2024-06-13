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
	"time"

	"github.com/google/uuid"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) CreateUser(user *usertypes.User, password string, roleIds []string) error {
	emailExists := s.usersMem.ForeachStop(func(i int, u *usertypes.User) bool {
		return u.Email == user.Email
	})
	if emailExists {
		return errors.New("email already exists")
	}

	passwordHash, err := hashPassword(password)
	if err != nil {
		return err
	}

	roles := s.rolesMem.FindByIds(roleIds)
	if len(roles) < len(roleIds) {
		return errors.New("some roles are not found")
	}

	user.PasswordHash = passwordHash
	user.RoleIds = roleIds
	if user.Id == "" {
		user.Id = uuid.NewString()
	}

	now := time.Now()
	user.UpdatedAt = now
	user.CreatedAt = now

	s.usersMem.Add(user)
	s.usersFile.MarkChanged()

	return nil
}
