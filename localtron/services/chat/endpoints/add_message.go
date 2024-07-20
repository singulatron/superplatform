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

	chatservice "github.com/singulatron/singulatron/localtron/services/chat"
	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

// AddMessage sends a new message to a chat thread
// @Summary Send a new message to a chat thread
// @Description Add a new message to a specific chat thread
// @Tags chat
// @Accept json
// @Produce json
// @Param request body chattypes.AddMessageRequest true "Add Message Request"
// @Success 200 {object} map[string]any "Message successfully added"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /chat/message/add [post]
func AddMessage(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	ds *chatservice.ChatService,
) {
	err := userService.IsAuthorized(chattypes.PermissionMessageCreate.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := chattypes.AddMessageRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ds.AddMessage(req.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}
