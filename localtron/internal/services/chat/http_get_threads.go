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

	chat "github.com/singulatron/singulatron/localtron/internal/services/chat/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// GetThreads retrieves a list of chat threads for a user
// @ID getThreads
// @Summary Get Threads
// @Description Fetch all chat threads associated with a specific user
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param request body chat.GetThreadsRequest false "Get Threads Request"
// @Success 200 {object} chat.GetThreadsResponse "Threads successfully retrieved"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/threads [post]
func (a *ChatService) GetThreads(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := a.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", chat.PermissionThreadView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	threads, err := a.getThreads(rsp.User.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(chat.GetThreadsResponse{
		Threads: threads,
	})
	w.Write(jsonData)
}
