/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package downloadtypes

import (
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

var PermissionDownloadCreate = usertypes.Permission{
	Id:   "download-svc:download:create",
	Name: "Download Service - Download Create",
}

var PermissionDownloadView = usertypes.Permission{
	Id:   "download-svc:download:view",
	Name: "Download Service - Download View",
}

var PermissionDownloadEdit = usertypes.Permission{
	Id:   "download-svc:download:edit",
	Name: "Download Service - Download Edit",
}

var PermissionDownloadDelete = usertypes.Permission{
	Id:   "download-svc:download:delete",
	Name: "Download Service - Download Delete",
}

var DownloadPermissions = []usertypes.Permission{
	PermissionDownloadCreate,
	PermissionDownloadView,
	PermissionDownloadEdit,
	PermissionDownloadDelete,
}
