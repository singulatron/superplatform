/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package dynamic_svc

import (
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

var PermissionGenericCreate = usertypes.Permission{
	Id:   "dynamic-svc:object:create",
	Name: "Dynamic Svc - Object Create",
}

var PermissionGenericView = usertypes.Permission{
	Id:   "dynamic-svc:object:view",
	Name: "Dynamic Svc - Object View",
}

var PermissionGenericEdit = usertypes.Permission{
	Id:   "dynamic-svc:object:edit",
	Name: "Dynamic Svc - Object Edit",
}

var PermissionGenericDelete = usertypes.Permission{
	Id:   "dynamic-svc:object:delete",
	Name: "Dynamic Svc - Object Delete",
}

var PermissionGenericStream = usertypes.Permission{
	Id:   "dynamic-svc:object:stream",
	Name: "Dynamic Svc - Object Stream",
}

var GenericPermissions = []usertypes.Permission{
	PermissionGenericCreate,
	PermissionGenericView,
	PermissionGenericEdit,
	PermissionGenericDelete,
	PermissionGenericStream,
}
