/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package downloadservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Do initiates a download request
// @Summary Download a File
// @Description Start a download for a specified URL.
// @Description
// @Description Requires the `download.create` permission.
// @Tags Download Service
// @Accept json
// @Produce json
// @Param request body downloadtypes.DownloadRequest true "Download Request"
// @Success 200 {object} map[string]any "Download initiated successfully"
// @Failure 400 {object} downloadtypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} downloadtypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} downloadtypes.ErrorResponse "Internal Server Error"
// @Router /download-service/download [put]
func (ds *DownloadService) Do(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := ds.router.AsRequestMaker(r).Post(r.Context(), "user-service", fmt.Sprintf("/permission/%v/is-authorized", downloadtypes.PermissionDownloadCreate.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := downloadtypes.DownloadRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ds.do(req.URL, req.FolderPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}
