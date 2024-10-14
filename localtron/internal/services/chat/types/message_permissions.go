/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package chat_svc

import (
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

var PermissionMessageCreate = usertypes.Permission{
	Id:   "chat-svc:message:create",
	Name: "Message Create",
}

var PermissionMessageView = usertypes.Permission{
	Id:   "chat-svc:message:view",
	Name: "Message View",
}

var PermissionMessageEdit = usertypes.Permission{
	Id:   "chat-svc:message:edit",
	Name: "Message Edit",
}

var PermissionMessageDelete = usertypes.Permission{
	Id:   "chat-svc:message:delete",
	Name: "Message Delete",
}

var PermissionMessageStream = usertypes.Permission{
	Id:   "chat-svc:message:stream",
	Name: "Message Stream",
}

var MessagePermissions = []usertypes.Permission{
	PermissionMessageCreate,
	PermissionMessageView,
	PermissionMessageEdit,
	PermissionMessageDelete,
	PermissionMessageStream,
}
