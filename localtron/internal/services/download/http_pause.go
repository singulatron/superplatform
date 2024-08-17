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
	"net/url"

	"github.com/gorilla/mux"
	download "github.com/singulatron/singulatron/localtron/internal/services/download/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// Pause pauses an ongoing download
// @ID pause
// @Summary Pause a Download
// @Description Pause a download that is currently in progress.
// @Description
// @Description Requires the `download-svc:download:edit` permission.
// @Tags Download Svc
// @Accept json
// @Produce json
// @Param downloadId path string true "Download ID"
// @Success 200 {object} map[string]any "Success response"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /download-svc/download/{downloadId}/pause [put]
func (ds *DownloadService) Pause(
	w http.ResponseWriter,
	r *http.Request,

) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := ds.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", download.PermissionDownloadEdit.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !rsp.Authorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	downloadId, err := url.PathUnescape(mux.Vars(r)["downloadId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Download ID in path is not URL encoded"))
		return
	}

	err = ds.pause(downloadId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}
