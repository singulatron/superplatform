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

var PermissionServiceInstanceCreate = usertypes.Permission{
	Id:   "registry-svc:service-instance:create",
	Name: " Create",
}

var PermissionServiceInstanceView = usertypes.Permission{
	Id:   "registry-svc:service-instance:view",
	Name: " View",
}

var PermissionServiceInstanceEdit = usertypes.Permission{
	Id:   "registry-svc:service-instance:edit",
	Name: " Edit",
}

var PermissionServiceInstanceDelete = usertypes.Permission{
	Id:   "registry-svc:service-instance:delete",
	Name: " Delete",
}

var ServiceInstancePermissions = []usertypes.Permission{
	PermissionServiceInstanceView,
}

var ServiceInstanceAdminPermissions = []usertypes.Permission{
	PermissionServiceInstanceView,
	PermissionServiceInstanceCreate,
	PermissionServiceInstanceEdit,
	PermissionServiceInstanceDelete,
}
