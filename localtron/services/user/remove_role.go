/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"errors"
	"time"

	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) RemoveRole(userId string, roleId string) error {
	query := s.usersStore.Query(
		datastore.Equal(datastore.Field("id"), userId),
	)
	userI, found, err := query.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	for i, existingRoleId := range user.RoleIds {
		if existingRoleId == roleId {
			user.RoleIds = append(user.RoleIds[:i], user.RoleIds[i+1:]...)
			user.UpdatedAt = time.Now()
		}
	}

	return query.Update(user)
}
