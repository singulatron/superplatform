/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dockerservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	dockertypes "github.com/singulatron/singulatron/localtron/services/docker/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// @Summary      Get Docker Host
// @Description  Retrieve information about the Docker host
// @Tags         Docker Service
// @Accept       json
// @Produce      json
// @Success      200   {object}  dockertypes.GetDockerHostResponse
// @Failure      401   {object}  dockertypes.ErrorResponse  "Unauthorized"
// @Failure      500   {object}  dockertypes.ErrorResponse  "Internal Server Error"
// @Security BearerAuth
// @Router       /docker-service/host [get]
func (dm *DockerService) Host(
	w http.ResponseWriter,
	req *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := dm.router.AsRequestMaker(req).Post(req.Context(), "user-service", fmt.Sprintf("/permission/%v/is-authorized", dockertypes.PermissionDockerView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	host := dm.getDockerHost()

	jsonData, _ := json.Marshal(dockertypes.GetDockerHostResponse{
		Host: host,
	})
	w.Write(jsonData)
}
