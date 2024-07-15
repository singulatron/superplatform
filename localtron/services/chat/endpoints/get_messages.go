/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
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

func GetMessages(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	ds *chatservice.ChatService,
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
