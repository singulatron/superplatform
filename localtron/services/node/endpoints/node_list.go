/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
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

	req := nodetypes.ListNodesRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	nodes, err := nodeService.ListNodes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := nodetypes.ListNodesResponse{
		Nodes: nodes,
	}

	bs, _ := json.Marshal(response)
	w.Write(bs)
}
