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
	model "github.com/singulatron/superplatform/server/internal/services/model/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// MakeDefault godoc
// @ID makeDefault
// @Summary Make a Model Default
// @Description Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.
// @Tags Model Svc
// @Accept json
// @Produce json
// @Param modelId path string true "Model ID"
// @Success 200 {object} model.MakeDefaultResponse
// @Failure 400 {object} model.ErrorResponse "Invalid JSON"
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/model/{modelId}/make-default [put]
func (ms *ModelService) MakeDefault(
	w http.ResponseWriter,
	r *http.Request,
) {

	rsp := &usertypes.IsAuthorizedResponse{}
	err := ms.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", model.PermissionModelEdit.Id), &usertypes.IsAuthorizedRequest{}, rsp)
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
	modelId, err := url.PathUnescape(vars["modelId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Model ID in path is not URL encoded"))
		return
	}

	err = ms.makeDefault(modelId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(model.MakeDefaultResponse{})
	w.Write(jsonData)
}
