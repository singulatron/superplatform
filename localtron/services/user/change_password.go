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

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) ChangePassword(email, currentPassword, newPassword string) error {
	s.runMutex.Lock()
	defer s.runMutex.Unlock()

	var errRet error
	changed := s.usersMem.ForeachStop(func(i int, user *usertypes.User) bool {
		if user.Email == email {
			if !checkPasswordHash(currentPassword, user.PasswordHash) {
				errRet = errors.New("current password is incorrect")
				return true
			}

			newPasswordHash, err := hashPassword(newPassword)
			errRet = err

			user.PasswordHash = newPasswordHash
			user.UpdatedAt = time.Now()

			return true
		}
		return false
	})

	if changed {
		s.usersFile.MarkChanged()
	}

	return errRet
}
