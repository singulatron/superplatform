/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package nodeservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	node "github.com/singulatron/singulatron/localtron/internal/services/node/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// @ID listNodes
// @Summary List Nodes
// @Description Retrieve a list of nodes.
// @Tags Node Svc
// @Accept json
// @Produce json
// @Param body body node.ListNodesRequest false "List Nodes Request"
// @Success 200 {object} node.ListNodesResponse
// @Failure 400 {object} node.ErrorResponse "Invalid JSON"
// @Failure 401 {object} node.ErrorResponse "Unauthorized"
// @Failure 500 {object} node.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /node-svc/nodes [post]
func (ns *NodeService) List(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := ns.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", node.PermissionNodeView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !rsp.Authorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	req := node.ListNodesRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	nodes, err := ns.listNodes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := node.ListNodesResponse{
		Nodes: nodes,
	}

	bs, _ := json.Marshal(response)
	w.Write(bs)
}
