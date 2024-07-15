/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package downloadtypes

import (
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

var PermissionDownloadCreate = usertypes.Permission{
	Id:   "download.create",
	Name: "Download Create",
}

var PermissionDownloadView = usertypes.Permission{
	Id:   "download.view",
	Name: "Download View",
}

var PermissionDownloadEdit = usertypes.Permission{
	Id:   "download.edit",
	Name: "Download Edit",
}

var PermissionDownloadDelete = usertypes.Permission{
	Id:   "download.delete",
	Name: "Download Delete",
}

var DownloadPermissions = []usertypes.Permission{
	PermissionDownloadCreate,
	PermissionDownloadView,
	PermissionDownloadEdit,
	PermissionDownloadDelete,
}
