/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package modelservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/singulatron/singulatron/localtron/internal/services/model/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// DefaultStatus godoc
// @ID getDefaultModelStatus
// @Summary Get Default Model Status
// @Description Retrieves the status of the default model.
// @Description
// @Description Requires the `model-svc:model:view` permission.
// @Tags Model Svc
// @Accept json
// @Produce json
// @Success 200 {object} model.StatusResponse "Model status retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/default-model/status [get]
func (ms *ModelService) DefaultStatus(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := ms.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", model.PermissionModelView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
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

	status, err := ms.status("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(model.StatusResponse{
		Status: status,
	})
	w.Write(jsonData)
}
