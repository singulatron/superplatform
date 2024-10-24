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
	"strconv"

	"github.com/gorilla/mux"
	docker "github.com/singulatron/superplatform/server/internal/services/docker/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// @ID getContainerSummary
// @Summary      Get Container Summary
// @Description  Get a summary of the Docker container identified by the hash, limited to a specified number of lines
// @Tags         Docker Svc
// @Accept       json
// @Produce      json
// @Param        hash           path     string  true  "Container Hash"
// @Param        numberOfLines  path     int     true  "Number of Lines"
// @Success      200            {object} docker.GetContainerSummaryResponse
// @Failure      400            {object} docker.ErrorResponse  "Invalid JSON"
// @Failure      401            {object} docker.ErrorResponse  "Unauthorized"
// @Failure      500            {object} docker.ErrorResponse  "Internal Server Error"
// @Security BearerAuth
// @Router       /docker-svc/container/{hash}/summary/{numberOfLines} [get]
func (dm *DockerService) Summary(
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
	hash := vars["hash"]
	lines, err := strconv.ParseInt(vars["numberOfLines"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	summary, err := dm.getContainerLogsAndStatus(hash, int(lines))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(&docker.GetContainerSummaryResponse{
		Summary: summary,
	})
	w.Write(jsonData)
}
