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

// MakeDefault godoc
// @Summary Make a Model Default
// @Description Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.
// @Tags model
// @Accept json
// @Produce json
// @Param id path string true "Model ID"
// @Success 200 {object} modeltypes.MakeDefaultResponse
// @Failure 400 {object} modeltypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} modeltypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} modeltypes.ErrorResponse "Internal Server Error"
// @Router /model/{modelId}/make-default [put]
func (ms *ModelService) MakeDefault(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := ms.router.AsRequestMaker(r).Post(r.Context(), "user-service", fmt.Sprintf("/permission/%v/is-authorized", modeltypes.PermissionModelEdit.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	err = ms.makeDefault(vars["modelId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(modeltypes.MakeDefaultResponse{})
	w.Write(jsonData)
}
