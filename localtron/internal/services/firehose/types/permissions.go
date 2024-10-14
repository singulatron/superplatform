/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package firehose_svc

import (
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

var PermissionFirehoseCreate = usertypes.Permission{
	Id:   "firehose-svc:firehose:create",
	Name: "Firehose Create",
}

var PermissionFirehoseView = usertypes.Permission{
	Id:   "firehose-svc:firehose:view",
	Name: "Firehose View",
}

var PermissionFirehoseEdit = usertypes.Permission{
	Id:   "firehose-svc:firehose:edit",
	Name: "Firehose Edit",
}

var PermissionFirehoseDelete = usertypes.Permission{
	Id:   "firehose-svc:firehose:delete",
	Name: "Firehose Delete",
}

var PermissionFirehoseStream = usertypes.Permission{
	Id:   "firehose-svc:firehose:stream",
	Name: "Firehose Stream",
}

var FirehosePermissions = []usertypes.Permission{
	PermissionFirehoseCreate,
	PermissionFirehoseView,
	PermissionFirehoseEdit,
	PermissionFirehoseDelete,
	PermissionFirehoseStream,
}
