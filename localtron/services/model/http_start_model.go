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
// @Param StartRequest body modeltypes.StartRequest false "Model start request"
// @Success 200 {object} modeltypes.StartResponse
// @Failure 400 {string} string "invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /model/{id}/start [post]
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
	err := ms.router.AsRequestMaker(r).Post(r.Context(), "user", "/is-authorized", &usertypes.IsAuthorizedRequest{
		PermissionId: modeltypes.PermissionModelCreate.Id,
	}, rsp)
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

	err = ms.start(v["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(modeltypes.StartResponse{})
	w.Write(jsonData)
}
