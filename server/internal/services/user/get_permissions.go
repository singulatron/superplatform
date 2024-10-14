/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"github.com/singulatron/superplatform/sdk/go/datastore"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

func (s *UserService) getPermissions() ([]*usertypes.Permission, error) {
	permissionsI, err := s.permissionsStore.Query().OrderBy(datastore.OrderByField("name", false)).Find()

	if err != nil {
		return nil, err
	}

	permissions := []*usertypes.Permission{}
	for _, permissionI := range permissionsI {
		permissions = append(permissions, permissionI.(*usertypes.Permission))
	}

	return permissions, nil
}
