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
	"net/url"

	"github.com/gorilla/mux"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Status godoc
// @Summary Get Model Status
// @Description Retrieves the status of a model by ID.
// @Description
// @Description Requires the `model.view` permission.
// @Tags model
// @Accept json
// @Produce json
// @Param id path string true "Model ID" // Path parameter for the model ID
// @Success 200 {object} modeltypes.StatusResponse "Model status retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /model/{id}/status [get]
func (ms *ModelService) Status(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := ms.router.AsRequestMaker(r).Post(r.Context(), "user-service", fmt.Sprintf("/permission/%v/is-authorized", modeltypes.PermissionModelView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	modelId := mux.Vars(r)["id"]
	unesc, err := url.PathUnescape(modelId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := ms.status(unesc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(modeltypes.StatusResponse{
		Status: status,
	})
	w.Write(jsonData)
}
