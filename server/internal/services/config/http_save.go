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

	config "github.com/singulatron/superplatform/server/internal/services/config/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// Save saves the configuration
// @Id saveConfig
// @Summary Save Config
// @Description Save the provided configuration to the server
// @Tags Config Svc
// @Accept json
// @Produce json
// @Param request body config.SaveConfigRequest true "Save Config Request"
// @Success 200 {object} config.SaveConfigResponse "Save Config Response"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /config-svc/config [put]
func (cs *ConfigService) Save(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := cs.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", config.PermissionConfigEdit.Id), &usertypes.IsAuthorizedRequest{
		SlugsGranted: []string{"model-svc"},
	}, rsp)
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

	req := &config.SaveConfigRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = cs.saveConfig(*req.Config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(config.SaveConfigResponse{})
	w.Write(jsonData)
}
