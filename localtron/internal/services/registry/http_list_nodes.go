/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package registryservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	registry "github.com/singulatron/singulatron/localtron/internal/services/registry/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// @ID listNodes
// @Summary List Nodes
// @Description Retrieve a list of nodes.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param body body registry.ListNodesRequest false "List Registrys Request"
// @Success 200 {object} registry.ListNodesResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid JSON"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/nodes [post]
func (ns *RegistryService) List(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := ns.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", registry.PermissionNodeView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
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

	req := registry.ListNodesRequest{}
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

	response := registry.ListNodesResponse{
		Nodes: nodes,
	}

	bs, _ := json.Marshal(response)
	w.Write(bs)
}

func (ns *RegistryService) listNodes() ([]*registry.Node, error) {
	nodeIs, err := ns.nodeStore.Query().Find()
	if err != nil {
		return nil, err
	}

	ret := []*registry.Node{}
	for _, nodeI := range nodeIs {
		ret = append(ret, nodeI.(*registry.Node))
	}

	return ret, err
}
