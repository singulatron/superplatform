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
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

func Start(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	ms *modelservice.ModelService,
) {
	err := userService.IsAuthorized(modeltypes.PermissionModelCreate.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := modeltypes.StartRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ms.Start(req.ModelId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(modeltypes.StartResponse{})
	w.Write(jsonData)
}
