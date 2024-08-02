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

	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// StartDefault godoc
// @Summary Start the Default Model
// @Description Starts The Default Model.
// @Description
// @Description Requires the `model.create` permission.
// @Tags Model Service
// @Accept json
// @Produce json
// @Success 200 {object} modeltypes.StartResponse
// @Failure 400 {object} modeltypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} modeltypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} modeltypes.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-service/default/start [put]
func (ms *ModelService) StartDefault(
	w http.ResponseWriter,
	r *http.Request,
) {
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

	err = ms.start("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(modeltypes.StartResponse{})
	w.Write(jsonData)
}
