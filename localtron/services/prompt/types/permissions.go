/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package prompttypes

import (
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

var PermissionPromptCreate = usertypes.Permission{
	Id:   "prompt.create",
	Name: "Prompt Create",
}

var PermissionPromptView = usertypes.Permission{
	Id:   "prompt.view",
	Name: "Prompt View",
}

var PermissionPromptEdit = usertypes.Permission{
	Id:   "prompt.edit",
	Name: "Prompt Edit",
}

var PermissionPromptDelete = usertypes.Permission{
	Id:   "prompt.delete",
	Name: "Prompt Delete",
}

var PermissionPromptStream = usertypes.Permission{
	Id:   "prompt.stream",
	Name: "Prompt Stream",
}

var PromptPermissions = []usertypes.Permission{
	PermissionPromptCreate,
	PermissionPromptView,
	PermissionPromptEdit,
	PermissionPromptDelete,
	PermissionPromptStream,
}
