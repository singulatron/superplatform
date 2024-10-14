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
	chat "github.com/singulatron/superplatform/server/internal/services/chat/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// GetMessages retrieves messages from a chat thread
// @ID getMessages
// @Summary List Messages
// @Description Fetch messages (and associated assets) for a specific chat thread.
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Success 200 {object} chat.GetMessagesResponse "Messages and assets successfully retrieved"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/thread/{threadId}/messages [post]
func (a *ChatService) GetMessages(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := a.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", chat.PermissionMessageView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	threadId := mux.Vars(r)["threadId"]

	messages, err := a.getMessages(threadId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	assetIds := []string{}
	for _, v := range messages {
		assetIds = append(assetIds, v.AssetIds...)
	}
	assets, err := a.getAssets(assetIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(chat.GetMessagesResponse{
		Messages: messages,
		Assets:   assets,
	})
	w.Write(jsonData)
}
