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
package appendpoints

import (
	"encoding/json"
	"net/http"

	appservice "github.com/singulatron/singulatron/localtron/services/app"
	types "github.com/singulatron/singulatron/localtron/services/app/types"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

func DeleteChatMessage(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	ds *appservice.AppService,
) {
	err := userService.IsAuthorized(types.PermissionMessageDelete.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := types.DeleteChatMessageRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ds.DeleteChatMessage(req.MessageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}
