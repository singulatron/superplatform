/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package chattypes

import (
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

var PermissionMessageCreate = usertypes.Permission{
	Id:   "message.create",
	Name: "Message Create",
}

var PermissionMessageView = usertypes.Permission{
	Id:   "message.view",
	Name: "Message View",
}

var PermissionMessageEdit = usertypes.Permission{
	Id:   "message.edit",
	Name: "Message Edit",
}

var PermissionMessageDelete = usertypes.Permission{
	Id:   "message.delete",
	Name: "Message Delete",
}

var PermissionMessageStream = usertypes.Permission{
	Id:   "message.stream",
	Name: "Message Stream",
}

var MessagePermissions = []usertypes.Permission{
	PermissionMessageCreate,
	PermissionMessageView,
	PermissionMessageEdit,
	PermissionMessageDelete,
	PermissionMessageStream,
}
