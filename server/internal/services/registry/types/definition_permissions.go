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

var PermissionDefinitionCreate = usertypes.Permission{
	Id:   "registry-svc:definition:create",
	Name: "Registry Create",
}

var PermissionDefinitionView = usertypes.Permission{
	Id:   "registry-svc:definition:view",
	Name: "Registry View",
}

var PermissionDefinitionEdit = usertypes.Permission{
	Id:   "registry-svc:definition:edit",
	Name: "Registry Edit",
}

var PermissionDefinitionDelete = usertypes.Permission{
	Id:   "registry-svc:definition:delete",
	Name: "Registry Delete",
}

var DefinitionPermissions = []usertypes.Permission{
	PermissionDefinitionView,
}

var DefinitionAdminPermissions = []usertypes.Permission{
	PermissionDefinitionCreate,
	PermissionDefinitionView,
	PermissionDefinitionEdit,
	PermissionDefinitionDelete,
}
