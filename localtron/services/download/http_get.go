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
	"net/http"

	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// @Summary Get Download
// @Description Get a download by URL.
// @Tags download
// @Accept json
// @Produce json
// @Success 200 {object} downloadtypes.GetDownloadResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /download/get [post]
func (ds *DownloadService) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := ds.router.AsRequestMaker(r).Post(r.Context(), "user", "/is-authorized", &usertypes.IsAuthorizedRequest{
		PermissionId: downloadtypes.PermissionDownloadView.Id,
	}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := downloadtypes.GetDownloadRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	dl, exists := ds.getDownload(req.Url)

	jsonData, _ := json.Marshal(downloadtypes.GetDownloadResponse{
		Exists:   exists,
		Download: downloadToDownloadDetails(req.Url, dl),
	})
	w.Write(jsonData)
}
