/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package deployservice

import (
	"context"
	"encoding/json"
	"net/http"

	sdk "github.com/singulatron/superplatform/sdk/go"
	deploy "github.com/singulatron/superplatform/server/internal/services/deploy/types"
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
func (ns *DeployService) ListDeployments(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	isAuthRsp, _, err := ns.clientFactory.Client(sdk.WithTokenFromRequest(r)).UserSvcAPI.IsAuthorized(context.Background(), deploy.PermissionDeploymentView.Id).Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAuthRsp.GetAuthorized() {
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
