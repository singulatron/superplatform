/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

import (
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

var PermissionRegistryCreate = usertypes.Permission{
	Id:   "node-svc:node:create",
	Name: "Registry Create",
}

var PermissionRegistryView = usertypes.Permission{
	Id:   "node-svc:node:view",
	Name: "Registry View",
}

var PermissionRegistryEdit = usertypes.Permission{
	Id:   "node-svc:node:edit",
	Name: "Registry Edit",
}

var PermissionRegistryDelete = usertypes.Permission{
	Id:   "node-svc:node:delete",
	Name: "Registry Delete",
}

var PermissionRegistryStream = usertypes.Permission{
	Id:   "node-svc:node:stream",
	Name: "Registry Stream",
}

var RegistryPermissions = []usertypes.Permission{
	PermissionRegistryCreate,
	PermissionRegistryView,
	PermissionRegistryEdit,
	PermissionRegistryDelete,
	PermissionRegistryStream,
}
