/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package chattypes

import (
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

var PermissionThreadCreate = usertypes.Permission{
	Id:   "thread.create",
	Name: "Thread Create",
}

var PermissionThreadView = usertypes.Permission{
	Id:   "thread.view",
	Name: "Thread View",
}

var PermissionThreadEdit = usertypes.Permission{
	Id:   "thread.edit",
	Name: "Thread Edit",
}

var PermissionThreadDelete = usertypes.Permission{
	Id:   "thread.delete",
	Name: "Thread Delete",
}

var PermissionThreadStream = usertypes.Permission{
	Id:   "thread.stream",
	Name: "Thread Stream",
}

var ThreadPermissions = []usertypes.Permission{
	PermissionThreadCreate,
	PermissionThreadView,
	PermissionThreadEdit,
	PermissionThreadDelete,
	PermissionThreadStream,
}
