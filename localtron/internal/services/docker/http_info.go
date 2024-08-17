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

	docker "github.com/singulatron/singulatron/localtron/internal/services/docker/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// @ID getInfo
// @Summary      Get Docker Service Information
// @Description  Retrieve detailed information about the Docker service
// @Tags         Docker Svc
// @Accept       json
// @Produce      json
// @Success      200   {object} docker.GetInfoResponse "Service Information"
// @Failure      401   {object} docker.ErrorResponse  "Unauthorized"
// @Failure      500   {object} docker.ErrorResponse  "Internal Server Error"
// @Security BearerAuth
// @Router       /docker-svc/info [get]
func (dm *DockerService) Info(
	w http.ResponseWriter,
	req *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := dm.router.AsRequestMaker(req).Post(req.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", docker.PermissionDockerView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	di, err := dm.info()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(docker.GetInfoResponse{
		Info: di,
	})
	w.Write(jsonData)
}
