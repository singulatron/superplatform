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
	"fmt"
	"net/http"

	config "github.com/singulatron/singulatron/localtron/services/config/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Get retrieves the current configuration
// @Id getConfig
// @Summary Get Config
// @Description Fetch the current configuration from the server
// @Tags Config Svc
// @Accept json
// @Produce json
// @Param request body config.GetConfigRequest true "Get Config Request"
// @Success 200 {object} config.GetConfigResponse "Current configuration retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /config-svc/get [post]
func (cs *ConfigService) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := cs.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", config.PermissionConfigView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conf, err := cs.getConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(config.GetConfigResponse{
		Config: &conf,
	})
	w.Write(jsonData)
}
