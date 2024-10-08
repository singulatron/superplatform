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

var PermissionServiceDefinitionCreate = usertypes.Permission{
	Id:   "registry-svc:service-definition:create",
	Name: "Registry Create",
}

var PermissionServiceDefinitionView = usertypes.Permission{
	Id:   "registry-svc:service-definition:view",
	Name: "Registry View",
}

var PermissionServiceDefinitionEdit = usertypes.Permission{
	Id:   "registry-svc:service-definition:edit",
	Name: "Registry Edit",
}

var PermissionServiceDefinitionDelete = usertypes.Permission{
	Id:   "registry-svc:service-definition:delete",
	Name: "Registry Delete",
}

var ServiceDefinitionPermissions = []usertypes.Permission{
	PermissionServiceDefinitionView,
}

var ServiceDefinitionAdminPermissions = []usertypes.Permission{
	PermissionServiceDefinitionCreate,
	PermissionServiceDefinitionView,
	PermissionServiceDefinitionEdit,
	PermissionServiceDefinitionDelete,
}
