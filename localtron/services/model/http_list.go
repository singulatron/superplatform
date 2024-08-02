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

// List godoc
// @Summary List Models
// @Description Retrieves a list of models after checking authorization
// @Description Requires "model.view" permission.
// @Tags model
// @Accept json
// @Produce json
// @Success 200 {object} modeltypes.ListResponse
// @Failure 401 {object} modeltypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} modeltypes.ErrorResponse "Internal Server Error"
// @Router /model-service/models [get]
func (ms *ModelService) List(
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

	models, err := ms.getModels()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(modeltypes.ListResponse{
		Models: models,
	})
	w.Write(jsonData)
}
