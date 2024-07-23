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

// GetMessages retrieves messages from a chat thread
// @Summary Retrieve messages from a chat thread
// @Description Fetch messages for a specific chat thread and associated assets
// @Tags chat
// @Accept json
// @Produce json
// @Param request body types.GetMessagesRequest true "Get Messages Request"
// @Success 200 {object} types.GetMessagesResponse "Messages and assets successfully retrieved"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /chat/messages [post]
func GetMessages(
	w http.ResponseWriter,
	r *http.Request,
	userService usertypes.UserServiceI,
	ds chattypes.ChatServiceI,
) {
	err := userService.IsAuthorized(types.PermissionMessageView.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := types.GetMessagesRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	messages, err := ds.GetMessages(req.ThreadId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	assetIds := []string{}
	for _, v := range messages {
		assetIds = append(assetIds, v.AssetIds...)
	}
	assets, err := ds.GetAssets(assetIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(types.GetMessagesResponse{
		Messages: messages,
		Assets:   assets,
	})
	w.Write(jsonData)
}
