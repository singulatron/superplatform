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

// @Summary Launch a Docker Container
// @Description Launches a Docker container with the specified parameters.
// @Description
// @Description Requires the `docker.create` permission.
// @Tags Docker Service
// @Accept json
// @Produce json
// @Param request body dockertypes.LaunchContainerRequest true "Launch Container Request"
// @Success 200 {object} dockertypes.LaunchContainerResponse
// @Failure 400 {object} dockertypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} dockertypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} dockertypes.ErrorResponse "Internal Server Error"
// @Router /docker-service/container [put]
func (dm *DockerService) LaunchContainer(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := dm.router.AsRequestMaker(r).Post(r.Context(), "user-service", fmt.Sprintf("/permission/%v/is-authorized", dockertypes.PermissionDockerCreate.Id), &usertypes.IsAuthorizedRequest{
		EmailsGranted: []string{"model-service"},
	}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := &dockertypes.LaunchContainerRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	di, err := dm.launchContainer(req.Image, req.Port, req.HostPort, req.Options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(&dockertypes.LaunchContainerResponse{
		Info: di,
	})
	w.Write(jsonData)
}
