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

	"github.com/pkg/errors"
	openapi "github.com/singulatron/superplatform/clients/go"
	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/datastore"
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

	registry := ns.clientFactory.Client(sdk.WithToken(ns.token)).RegistrySvcAPI

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
		return errors.Wrap(err, "Error calling list nodes")
	}

	listInstancesRsp, _, err := registry.ListInstances(ctx).Execute()
	if err != nil {
		return errors.Wrap(err, "Error calling list service instances")
	}

	listDefinitionsRsp, _, err := registry.ListDefinitions(ctx).Execute()
	if err != nil {
		return errors.Wrap(err, "Error calling list service definitions")
	}

	commands := allocator.GenerateCommands(listNodesRsp.Nodes, listInstancesRsp.Instances, deployments)

	for _, command := range commands {
		var node *openapi.RegistrySvcNode
		var definition *openapi.RegistrySvcDefinition

		for _, v := range listNodesRsp.Nodes {
			if v.Url == command.NodeUrl {
				node = &v
			}
		}

		for _, v := range listDefinitionsRsp.Definitions {
			if v.Id == command.DeploymentId {
				definition = &v
			}
		}

		err := ns.processCommand(ctx, command, node, definition)
		if err != nil {
			logger.Error("Error processing deploy command", slog.Any("error", err))
		}
	}

	return nil
}

func (ns *DeployService) processCommand(
	ctx context.Context,
	command *deploy.Command,
	node *openapi.RegistrySvcNode,
	definition *openapi.RegistrySvcDefinition,
) error {
	deployment, err := ns.getDeploymentById(command.DeploymentId)
	if err != nil {
		return err
	}
	if deployment.Status == deploy.StatusPending {
		deployment.Status = deploy.StatusDeploying
		err = ns.deploymentStore.Upsert(deployment)
		if err != nil {
			return err
		}
	}

	switch command.Action {
	case deploy.CommandTypeStart:
		logger.Info("Executing start command", slog.String("deploymentId", deployment.Id))

		_, _, err := ns.clientFactory.Client(sdk.WithAddress(*command.NodeUrl), sdk.WithToken(ns.token)).DockerSvcAPI.LaunchContainer(ctx).Request(
			openapi.DockerSvcLaunchContainerRequest{
				Image: definition.Image.Name,
				Port:  definition.Image.Port,
			},
		).Execute()
		err = sdk.OpenAPIError(err)

		if err != nil {
			logger.Info("Error executing start command", slog.Any("error", err))
			deployment, readErr := ns.getDeploymentById(command.DeploymentId)
			if readErr != nil {
				return readErr
			}
			deployment.Status = deploy.DeploymentStatus(openapi.StatusError)
			deployment.Details = err.Error()

			writeErr := ns.deploymentStore.Upsert(deployment)
			if writeErr != nil {
				return writeErr
			}
		}
	case deploy.CommandTypeScale:
	case deploy.CommandTypeKill:
	}

	return nil
}

func (ns *DeployService) getDeploymentById(deploymentId string) (*deploy.Deployment, error) {
	deploymentIs, err := ns.deploymentStore.Query(datastore.Id(deploymentId)).Find()
	if err != nil {
		return nil, err
	}

	deployment := deploymentIs[0].(*deploy.Deployment)

	return deployment, err
}
