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

// StartSpecific godoc
// @ID startModel
// @Summary Start a Model
// @Description Starts a model by ID
// @Tags Model Svc
// @Accept json
// @Produce json
// @Param modelId path string true "Model ID"
// @Success 200 {object} model.StartResponse
// @Failure 400 {object} model.ErrorResponse "Invalid JSON"
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/model/{modelId}/start [put]
func (ms *ModelService) StartSpecific(
	w http.ResponseWriter,
	r *http.Request,
) {

	v := mux.Vars(r)
	if v["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing model ID in request path"))
		return
	}

	rsp := &usertypes.IsAuthorizedResponse{}
	err := ms.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", model.PermissionModelCreate.Id), &usertypes.IsAuthorizedRequest{}, rsp)
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

	req := model.StartRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	modelId, err := url.PathUnescape(v["modelId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = ms.start(modelId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(model.StartResponse{})
	w.Write(jsonData)
}
