/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package nodeendpoints

import (
	"encoding/json"
	"net/http"

	nodeservice "github.com/singulatron/singulatron/localtron/services/node"
	nodetypes "github.com/singulatron/singulatron/localtron/services/node/types"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

func List(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	nodeService *nodeservice.NodeService,
) {
	err := userService.IsAuthorized(nodetypes.PermissionNodeView.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user, found, err := userService.GetUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !found {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	req := nodetypes.ListNodesRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	nodes, count, err := nodeService.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := range nodes {
		if nodes[i].UserId != user.Id {
			// do not let users see other peoples promtps,
			// not even if they are admins
			// eg. imagine a sysadmin looking at the CEO's node
			nodes[i].Prompt = ""
		}
	}

	response := nodetypes.ListPromptsResponse{
		Prompts: nodes,
		Count:   count,
	}
	if len(nodes) >= 20 {
		response.After = nodes[len(nodes)-1].CreatedAt
	}

	bs, _ := json.Marshal(response)
	w.Write(bs)
}
