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

// @ID saveDeployment
// @Summary Save Deployment
// @Description Save a deployment.
// @Tags Deploy Svc
// @Accept json
// @Produce json
// @Param body body deploy.SaveDeploymentRequest false "Save Deploys Request"
// @Success 200 {object} deploy.SaveDeploymentResponse
// @Failure 400 {object} deploy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} deploy.ErrorResponse "Unauthorized"
// @Failure 500 {object} deploy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /deploy-svc/deployment [put]
func (ns *DeployService) SaveDeployment(
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

	req := deploy.SaveDeploymentRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = ns.saveDeployment(req.Deployment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{}`))
}

func (ns *DeployService) saveDeployment(deployment *deploy.Deployment) error {
	return ns.deploymentStore.Upsert(deployment)
}
