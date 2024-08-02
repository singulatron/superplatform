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

// @Summary      Check If a Container Is Running
// @Description  Check if a Docker container identified by the hash is running
// @Tags         Docker Service
// @Accept       json
// @Produce      json
// @Param        hash  path      string  true  "Container Hash"
// @Success      200   {object}  dockertypes.ContainerIsRunningResponse
// @Failure      400   {object}  dockertypes.ErrorResponse  "Invalid JSON"
// @Failure      401   {object}  dockertypes.ErrorResponse  "Unauthorized"
// @Failure      500   {object}  dockertypes.ErrorResponse  "Internal Server Error"
// @SecurityDefinitions.bearerAuth BearerAuth
// @Security     BearerAuth
// @Router       /docker-service/container/{hash}/is-running [get]
func (dm *DockerService) HashIsRunning(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := dm.router.AsRequestMaker(r).Post(r.Context(), "user-service", fmt.Sprintf("/permission/%v/is-authorized", dockertypes.PermissionDockerView.Id), &usertypes.IsAuthorizedRequest{
		EmailsGranted: []string{"model"},
	}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := &dockertypes.ContainerIsRunningRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	isRunning, err := dm.hashIsRunning(req.Hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(&dockertypes.ContainerIsRunningResponse{
		IsRunning: isRunning,
	})
	w.Write(jsonData)
}
