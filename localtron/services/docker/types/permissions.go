/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package docker_svc

import (
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

var PermissionDockerCreate = usertypes.Permission{
	Id:   "docker-svc:container:create",
	Name: "Docker Service - Container Create",
}

var PermissionDockerView = usertypes.Permission{
	Id:   "docker-svc:container:view",
	Name: "Docker Service - Container View",
}

var PermissionDockerEdit = usertypes.Permission{
	Id:   "docker-svc:container:.edit",
	Name: "Docker Service - Container Edit",
}

var DockerPermissions = []usertypes.Permission{
	PermissionDockerView,
	PermissionDockerEdit,
}
