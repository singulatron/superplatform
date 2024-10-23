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

	"github.com/gorilla/mux"
	docker "github.com/singulatron/superplatform/server/internal/services/docker/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// @ID isRunning
// @Summary      Check If a Container Is Running
// @Description  Check if a Docker container identified by the hash is running
// @Tags         Docker Svc
// @Accept       json
// @Produce      json
// @Param        hash  path      string  true  "Container Hash"
// @Success      200   {object}  docker.ContainerIsRunningResponse
// @Failure      400   {object}  docker.ErrorResponse  "Invalid JSON"
// @Failure      401   {object}  docker.ErrorResponse  "Unauthorized"
// @Failure      500   {object}  docker.ErrorResponse  "Internal Server Error"
// @SecurityDefinitions.bearerAuth BearerAuth
// @Security     BearerAuth
// @Router       /docker-svc/container/{hash}/is-running [get]
func (dm *DockerService) HashIsRunning(
	w http.ResponseWriter,
	r *http.Request,
) {

	rsp := &usertypes.IsAuthorizedResponse{}
	err := dm.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", docker.PermissionDockerView.Id), &usertypes.IsAuthorizedRequest{
		SlugsGranted: []string{"model-svc"},
	}, rsp)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	if !rsp.Authorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	vars := mux.Vars(r)

	isRunning, err := dm.hashIsRunning(vars["hash"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(&docker.ContainerIsRunningResponse{
		IsRunning: isRunning,
	})
	w.Write(jsonData)
}
