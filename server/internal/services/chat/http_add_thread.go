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

	chat "github.com/singulatron/superplatform/server/internal/services/chat/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// AddThread creates a new chat thread
// @ID addThread
// @Summary Add Thread
// @Description Create a new chat thread and add the requesting user to it.
// @Decription
// @Description Requires the `chat-svc:thread:create` permission.
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param request body chat.AddThreadRequest true "Add Thread Request"
// @Success 200 {object} chat.AddThreadResponse "Thread successfully created"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/thread [post]
func (a *ChatService) AddThread(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := a.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", chat.PermissionThreadCreate.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	if !rsp.Authorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	req := chat.AddThreadRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	if req.Thread == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`missing thread`))
		return
	}

	req.Thread.UserIds = append(req.Thread.UserIds, rsp.User.Id)

	thread, err := a.addThread(req.Thread)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(chat.AddThreadResponse{
		Thread: thread,
	})
	w.Write(jsonData)
}
