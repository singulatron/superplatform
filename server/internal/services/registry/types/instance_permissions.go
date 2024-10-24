/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

import (
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

var PermissionInstanceCreate = usertypes.Permission{
	Id:   "registry-svc:instance:create",
	Name: " Create",
}

var PermissionInstanceView = usertypes.Permission{
	Id:   "registry-svc:instance:view",
	Name: " View",
}

var PermissionInstanceEdit = usertypes.Permission{
	Id:   "registry-svc:instance:edit",
	Name: " Edit",
}

var PermissionInstanceDelete = usertypes.Permission{
	Id:   "registry-svc:instance:delete",
	Name: " Delete",
}

var InstancePermissions = []usertypes.Permission{
	PermissionInstanceView,
}

var InstanceAdminPermissions = []usertypes.Permission{
	PermissionInstanceView,
	PermissionInstanceCreate,
	PermissionInstanceEdit,
	PermissionInstanceDelete,
}
