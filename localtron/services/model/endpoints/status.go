/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package modelendpoints

import (
	"encoding/json"
	"net/http"

	modelservice "github.com/singulatron/singulatron/localtron/services/model"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
)

type StatusRequest struct {
	Url string `json:"url"`
}

func Status(w http.ResponseWriter, r *http.Request, ms *modelservice.ModelService) {
	req := StatusRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	status, err := ms.Status(req.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(modeltypes.StatusResponse{
		Status: status,
	})
	w.Write(jsonData)
}
