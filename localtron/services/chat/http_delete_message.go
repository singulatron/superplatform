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
	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// DeleteMessage removes a message from a chat thread
// @Summary Delete Message
// @Description Delete a specific message from a chat thread by its ID
// @Tags Chat Service
// @Accept json
// @Produce json
// @Param messageId path string true "Message ID"
// @Success 200 {object} map[string]any "Message successfully deleted"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /chat-service/message/{messageId} [delete]
func (a *ChatService) DeleteMessage(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := a.router.AsRequestMaker(r).Post(r.Context(), "user-service", fmt.Sprintf("/permission/%v/is-authorized", chattypes.PermissionMessageDelete.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	threadId := vars["threadId"]

	err = a.deleteMessage(threadId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}
