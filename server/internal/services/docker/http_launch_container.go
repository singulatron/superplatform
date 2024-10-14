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

	docker "github.com/singulatron/superplatform/server/internal/services/docker/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// @ID launchContainer
// @Summary Launch a Container
// @Description Launches a Docker container with the specified parameters.
// @Description
// @Description Requires the `docker-svc:docker:create` permission.
// @Tags Docker Svc
// @Accept json
// @Produce json
// @Param request body docker.LaunchContainerRequest true "Launch Container Request"
// @Success 200 {object} docker.LaunchContainerResponse
// @Failure 400 {object} docker.ErrorResponse "Invalid JSON"
// @Failure 401 {object} docker.ErrorResponse "Unauthorized"
// @Failure 500 {object} docker.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /docker-svc/container [put]
func (dm *DockerService) LaunchContainer(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := dm.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", docker.PermissionDockerCreate.Id), &usertypes.IsAuthorizedRequest{
		SlugsGranted: []string{"model-svc"},
	}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := &docker.LaunchContainerRequest{}
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

	jsonData, _ := json.Marshal(&docker.LaunchContainerResponse{
		Info: di,
	})
	w.Write(jsonData)
}
