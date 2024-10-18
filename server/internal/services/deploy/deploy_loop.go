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
	"log/slog"
	"time"

	"github.com/singulatron/superplatform/sdk/go/logger"
	"github.com/singulatron/superplatform/server/internal/services/deploy/allocator"
	deploy "github.com/singulatron/superplatform/server/internal/services/deploy/types"
)

func (ns *DeployService) loop() {
	for {
		func() {
			if r := recover(); r != nil {
				logger.Error("Deploy cycle panic", slog.Any("panic", r))
			}

			err := ns.cycle()
			if err != nil {
				logger.Error("Deploy cycle error", slog.Any("error", err))
			}
			time.Sleep(5 * time.Second)
		}()
	}
}

func (ns *DeployService) cycle() error {
	ctx := context.Background()

	ns.lock.Acquire(ctx, "deploy-svc-deploy")
	defer ns.lock.Release(ctx, "deploy-svc-deploy")

	registry := ns.clientFactory.Client().RegistrySvcAPI

	deploymentIs, err := ns.deploymentStore.Query().Find()
	if err != nil {
		return err
	}

	deployments := []*deploy.Deployment{}

	for _, deploymentI := range deploymentIs {
		deployment := deploymentI.(*deploy.Deployment)
		deployments = append(deployments, deployment)
	}

	listNodesRsp, _, err := registry.ListNodes(ctx).Execute()
	if err != nil {
		return err
	}

	serviceInstances, _, err := registry.QueryServiceInstances(ctx).Execute()
	if err != nil {
		return err
	}

	allocator.GenerateCommands(listNodesRsp.Nodes, serviceInstances)

	return nil
}
