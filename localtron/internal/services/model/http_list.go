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

// List godoc
// @ID listModels
// @Summary List Models
// @Description Retrieves a list of models.
// @Description
// @Description Requires `model-svc:model:view` permission.
// @Tags Model Svc
// @Accept json
// @Produce json
// @Success 200 {object} model.ListResponse
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/models [post]
func (ms *ModelService) List(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := ms.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", model.PermissionModelView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
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

	jsonData, _ := json.Marshal(model.ListResponse{
		Models: models,
	})
	w.Write(jsonData)
}
