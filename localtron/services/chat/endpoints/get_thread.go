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
	types "github.com/singulatron/singulatron/localtron/services/chat/types"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

// GetThread retrieves details of a specific chat thread
// @Summary Retrieve details of a chat thread
// @Description Fetch information about a specific chat thread by its ID
// @Tags chat
// @Accept json
// @Produce json
// @Param request body types.GetThreadRequest true "Get Thread Request"
// @Success 200 {object} types.GetThreadResponse "Thread details successfully retrieved"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /chat/thread [post]
func GetThread(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	ds *chatservice.ChatService,
) {
	err := userService.IsAuthorized(types.PermissionThreadCreate.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := types.GetThreadRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	thread, _, err := ds.GetThread(req.ThreadId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(types.GetThreadResponse{
		Thread: *thread,
	})
	w.Write(jsonData)
}
