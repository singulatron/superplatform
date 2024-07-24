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
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// GetThread retrieves details of a specific chat thread
// @Summary Get Thread
// @Description Fetch information about a specific chat thread by its ID
// @Tags chat
// @Accept json
// @Produce json
// @Param request body chattypes.GetThreadRequest true "Get Thread Request"
// @Success 200 {object} chattypes.GetThreadResponse "Thread details successfully retrieved"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /chat/thread [post]
func GetThread(
	w http.ResponseWriter,
	r *http.Request,
	userService usertypes.UserServiceI,
	ds chattypes.ChatServiceI,
) {
	err := userService.IsAuthorized(chattypes.PermissionThreadCreate.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := chattypes.GetThreadRequest{}
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

	jsonData, _ := json.Marshal(chattypes.GetThreadResponse{
		Thread: *thread,
	})
	w.Write(jsonData)
}
