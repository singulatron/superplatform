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
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

// List retrieves a list of download details
// @Summary Retrieve download details
// @Description Fetch a list of all download details
// @Tags download
// @Accept json
// @Produce json
// @Success 200 {object} downloadtypes.DownloadsResponse "List of downloads"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /download/list [post]
func List(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	ds *downloadservice.DownloadService,
) {
	err := userService.IsAuthorized(downloadtypes.PermissionDownloadView.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	details, err := ds.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(downloadtypes.DownloadsResponse{
		Downloads: details,
	})
	w.Write(jsonData)
}
