/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package downloadendpoints

import (
	"encoding/json"
	"net/http"

	downloadservice "github.com/singulatron/singulatron/localtron/services/download"
	types "github.com/singulatron/singulatron/localtron/services/download/types"

	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

// Pause pauses an ongoing download
// @Summary Pause an ongoing download
// @Description Pause a download that is currently in progress
// @Tags download
// @Accept json
// @Produce json
// @Param body body types.DownloadRequest true "Download request payload"
// @Success 200 {object} map[string]any "Success response"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /download/pause [post]
func Pause(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	ds *downloadservice.DownloadService,
) {
	err := userService.IsAuthorized(types.PermissionDownloadEdit.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := types.DownloadRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ds.Pause(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}
