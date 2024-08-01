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

	"github.com/gorilla/mux"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// StartSpecific godoc
// @Summary Start a Model
// @Description Starts a model by ID
// @Tags model
// @Accept json
// @Produce json
// @Param id path string true "Model ID"
// @Success 200 {object} modeltypes.StartResponse
// @Failure 400 {object} modeltypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} modeltypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} modeltypes.ErrorResponse "Internal Server Error"
// @Router /model-service/{modelId}/start [put]
func (ms *ModelService) StartSpecific(
	w http.ResponseWriter,
	r *http.Request,
) {
	v := mux.Vars(r)
	if v["id"] == "" {
		http.Error(w, "Missing model ID in request path", http.StatusUnauthorized)
		return
	}

	rsp := &usertypes.IsAuthorizedResponse{}
	err := ms.router.AsRequestMaker(r).Post(r.Context(), "user-service", fmt.Sprintf("/permission/%v/is-authorized", modeltypes.PermissionModelCreate.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := modeltypes.StartRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ms.start(v["modelId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(modeltypes.StartResponse{})
	w.Write(jsonData)
}
