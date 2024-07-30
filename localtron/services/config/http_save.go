/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package configservice

import (
	"encoding/json"
	"net/http"

	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Get retrieves the current configuration
// @Summary Get
// @Description Fetch the current configuration from the server
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} configtypes.ConfigGetResponse "Current configuration retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /config/get [post]
func (cs *ConfigService) Save(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := cs.router.AsRequestMaker(r).Post(r.Context(), "user", "/is-authorized", &usertypes.IsAuthorizedRequest{
		PermissionId: configtypes.PermissionConfigEdit.Id,
	}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := &configtypes.SaveConfigRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = cs.saveConfig(*req.Config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(configtypes.SaveConfigResponse{})
	w.Write(jsonData)
}
