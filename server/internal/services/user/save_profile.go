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

	"github.com/singulatron/superplatform/sdk/go/datastore"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

func (s *UserService) saveProfile(email, newName string) error {
	query := s.usersStore.Query(
		datastore.Equals(datastore.Field("email"), email),
	)

	userI, found, err := query.FindOne()
	if err != nil {
		return err
	}

	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	user.Name = newName
	user.UpdatedAt = time.Now()

	query.Update(user)

	return nil
}
