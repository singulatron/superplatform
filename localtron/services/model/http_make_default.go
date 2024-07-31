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

	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// MakeDefault godoc
// @Summary Make Default
// @Description Sets a model as the default model - when prompts are sent without a Model ID, this model will be used.
// @Tags model
// @Accept json
// @Produce json
// @Param MakeDefaultRequest body modeltypes.MakeDefaultRequest true "Make default request"
// @Success 200 {object} modeltypes.MakeDefaultResponse
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /model/{{id}}/make-default [post]
func (ms *ModelService) MakeDefault(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := ms.router.AsRequestMaker(r).Post(r.Context(), "user", "/is-authorized", &usertypes.IsAuthorizedRequest{
		PermissionId: modeltypes.PermissionModelEdit.Id,
	}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := modeltypes.MakeDefaultRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ms.makeDefault(req.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(modeltypes.MakeDefaultResponse{})
	w.Write(jsonData)
}
