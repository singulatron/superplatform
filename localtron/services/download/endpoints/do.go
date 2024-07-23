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
	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Do initiates a download request
// @Summary Initiate a file download
// @Description Start a download for a specified URL and folder path
// @Tags download
// @Accept json
// @Produce json
// @Param request body downloadtypes.DownloadRequest true "Download Request"
// @Success 200 {object} map[string]any "Download initiated successfully"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /download/do [post]
func Do(
	w http.ResponseWriter,
	r *http.Request,
	userService usertypes.UserServiceI,
	ds *downloadservice.DownloadService,
) {
	err := userService.IsAuthorized(downloadtypes.PermissionDownloadCreate.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := downloadtypes.DownloadRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ds.Do(req.URL, req.FolderPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}
