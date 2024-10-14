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
	"fmt"

	"github.com/singulatron/superplatform/sdk/go/datastore"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

func (s *UserService) RemoveRole(userId string, roleId string) error {
	query := s.usersStore.Query(
		datastore.Equals(datastore.Field("id"), userId),
	)
	userI, found, err := query.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	return s.userRoleLinksStore.Query(
		datastore.Id(fmt.Sprintf("%v:%v", user.Id, roleId)),
	).Delete()

}
