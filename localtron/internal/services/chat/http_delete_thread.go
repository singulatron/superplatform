/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	chattypes "github.com/singulatron/singulatron/localtron/internal/services/chat/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// DeleteThread removes a chat thread
// @ID deleteThread
// @Summary Delete a Thread
// @Description Delete a specific chat thread by its ID
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Success 200 {object} map[string]any "Thread successfully deleted"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/thread/{threadId} [delete]
func (a *ChatService) DeleteThread(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := a.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", chattypes.PermissionThreadCreate.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	err = a.deleteThread(vars["threadId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}
