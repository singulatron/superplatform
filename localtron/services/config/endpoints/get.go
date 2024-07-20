/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package configendpoints

import (
	"encoding/json"
	"net/http"

	configservice "github.com/singulatron/singulatron/localtron/services/config"
	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"

	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

// Get retrieves the current configuration
// @Summary Retrieve the current configuration
// @Description Fetch the current configuration from the server
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} configtypes.ConfigGetResponse "Current configuration retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /config/get [post]
func Get(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	cs *configservice.ConfigService,
) {
	err := userService.IsAuthorized(configtypes.PermissionConfigView.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	conf, err := cs.GetConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(configtypes.ConfigGetResponse{
		Config: &conf,
	})
	w.Write(jsonData)
}
