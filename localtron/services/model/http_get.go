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

// Get godoc
// @Summary Get a Model
// @Description Retrieves the details of a model by its ID.
// @Description
// @Description the Requires `model.view` permission.
// @Tags Model Service
// @Accept json
// @Produce json
// @Param modelId path string true "Model ID"
// @Success 200 {object} modeltypes.GetModelResponse
// @Failure 401 {object} modeltypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} modeltypes.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/{modelId} [get]
func (ms *ModelService) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := ms.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", modeltypes.PermissionModelView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	if vars["modelId"] == "" {
		http.Error(w, "no model id", http.StatusBadRequest)
		return
	}

	mid, err := url.PathUnescape(vars["modelId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, found, err := ms.getModel(mid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !found {
		http.Error(w, "model not found", http.StatusBadRequest)
		return
	}

	platform, _, err := ms.getPlatform(model.PlatformId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(modeltypes.GetModelResponse{
		Exists:   found,
		Model:    model,
		Platform: platform,
	})
	w.Write(jsonData)
}
