/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package deployservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	deploy "github.com/singulatron/superplatform/server/internal/services/deploy/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// @ID listDeployments
// @Summary List Deployments
// @Description Retrieve a list of deployments.
// @Tags Deploy Svc
// @Accept json
// @Produce json
// @Param body body deploy.ListDeploymentsRequest false "List Deploys Request"
// @Success 200 {object} deploy.ListDeploymentsResponse
// @Failure 400 {object} deploy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} deploy.ErrorResponse "Unauthorized"
// @Failure 500 {object} deploy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /deploy-svc/deployments [post]
func (ns *DeployService) List(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := ns.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", deploy.PermissionDeploymentView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
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

	req := deploy.ListDeploymentsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	deployments, err := ns.listDeployments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := deploy.ListDeploymentsResponse{
		Deployments: deployments,
	}

	bs, _ := json.Marshal(response)
	w.Write(bs)
}

func (ns *DeployService) listDeployments() ([]*deploy.Deployment, error) {
	deploymentIs, err := ns.deploymentStore.Query().Find()
	if err != nil {
		return nil, err
	}

	ret := []*deploy.Deployment{}
	for _, deploymentI := range deploymentIs {
		ret = append(ret, deploymentI.(*deploy.Deployment))
	}

	return ret, err
}
