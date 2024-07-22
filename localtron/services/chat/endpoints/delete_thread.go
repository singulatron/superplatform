/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package appendpoints

import (
	"encoding/json"
	"net/http"

	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	types "github.com/singulatron/singulatron/localtron/services/chat/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// DeleteThread removes a chat thread
// @Summary Remove a chat thread
// @Description Delete a specific chat thread by its ID
// @Tags chat
// @Accept json
// @Produce json
// @Param request body types.DeleteThreadRequest true "Delete Thread Request"
// @Success 200 {object} map[string]any "Thread successfully deleted"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /chat/thread/delete [post]
func DeleteThread(
	w http.ResponseWriter,
	r *http.Request,
	userService usertypes.UserServiceI,
	ds chattypes.ChatServiceI,
) {
	err := userService.IsAuthorized(types.PermissionThreadCreate.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := types.DeleteThreadRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ds.DeleteThread(req.ThreadId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}
