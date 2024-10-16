/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package deploy_svc

import (
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

var PermissionDeploymentCreate = usertypes.Permission{
	Id:   "deploy-svc:deployment:create",
	Name: "Deploy Svc - Deployment Create",
}

var PermissionDeploymentView = usertypes.Permission{
	Id:   "deploy-svc:deployment:view",
	Name: "Deploy Svc - Deployment View",
}

var PermissionDeploymentEdit = usertypes.Permission{
	Id:   "deploy-svc:deployment:create",
	Name: "Deploy Svc - Deployment Create",
}

var PermissionDeploymentDelete = usertypes.Permission{
	Id:   "deploy-svc:deployment:delete",
	Name: "Deploy Svc - Deployment Delete",
}

var DeployAdminPermissions = []usertypes.Permission{
	PermissionDeploymentCreate,
	PermissionDeploymentEdit,
	PermissionDeploymentView,
	PermissionDeploymentDelete,
}
