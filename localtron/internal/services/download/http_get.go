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

// @ID getDownload
// @Summary Get a Download
// @Description Get a download by ID.
// @Description
// @Description Requires the `download-svc:download:view` permission.
// @Tags Download Svc
// @Accept json
// @Produce json
// @Param downloadId path string true "Download ID"
// @Success 200 {object} download.GetDownloadResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /download-svc/download/{downloadId} [get]
func (ds *DownloadService) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	rsp := &usertypes.IsAuthorizedResponse{}
	err := ds.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", download.PermissionDownloadView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
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

	vars := mux.Vars(r)
	did, err := url.PathUnescape(vars["downloadId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	dl, exists := ds.getDownload(did)

	jsonData, _ := json.Marshal(download.GetDownloadResponse{
		Exists:   exists,
		Download: downloadToDownloadDetails(vars["downloadId"], dl),
	})
	w.Write(jsonData)
}
