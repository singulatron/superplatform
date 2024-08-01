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

	"github.com/gorilla/mux"
	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// @Summary Get Download
// @Description Get a download by ID.
// @Tags download
// @Accept json
// @Produce json
// @Success 200 {object} downloadtypes.GetDownloadResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /download-service/download/{downloadId} [get]
func (ds *DownloadService) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := ds.router.AsRequestMaker(r).Post(r.Context(), "user-service", fmt.Sprintf("/permission/%v/is-authorized", downloadtypes.PermissionDownloadView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	dl, exists := ds.getDownload(vars["downloadId"])

	jsonData, _ := json.Marshal(downloadtypes.GetDownloadResponse{
		Exists:   exists,
		Download: downloadToDownloadDetails(vars["downloadId"], dl),
	})
	w.Write(jsonData)
}
