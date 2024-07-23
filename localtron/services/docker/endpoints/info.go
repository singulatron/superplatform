/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dockerendpoints

import (
	"encoding/json"
	"net/http"

	dockertypes "github.com/singulatron/singulatron/localtron/services/docker/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func Info(
	w http.ResponseWriter,
	req *http.Request,
	userService usertypes.UserServiceI,
	dm dockertypes.DockerServiceI,
) {
	err := userService.IsAuthorized(dockertypes.PermissionDockerView.Id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	di, err := dm.Info()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{
		"info": di,
	})
	w.Write(jsonData)
}
