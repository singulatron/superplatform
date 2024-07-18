/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package configtypes

import (
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

var PermissionConfigCreate = usertypes.Permission{
	Id:   "config.create",
	Name: "Config Create",
}

var PermissionConfigView = usertypes.Permission{
	Id:   "config.view",
	Name: "Config View",
}

var PermissionConfigEdit = usertypes.Permission{
	Id:   "config.edit",
	Name: "Config Edit",
}

var PermissionConfigDelete = usertypes.Permission{
	Id:   "config.delete",
	Name: "Config Delete",
}

var PermissionConfigStream = usertypes.Permission{
	Id:   "config.stream",
	Name: "Config Stream",
}

var ConfigPermissions = []usertypes.Permission{
	PermissionConfigCreate,
	PermissionConfigView,
	PermissionConfigEdit,
	PermissionConfigDelete,
	PermissionConfigStream,
}
