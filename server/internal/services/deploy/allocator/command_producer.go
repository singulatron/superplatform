/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package allocator

import (
	openapi "github.com/singulatron/superplatform/clients/go"
	deploy "github.com/singulatron/superplatform/server/internal/services/deploy/types"
)

func GenerateCommands(
	nodes []openapi.RegistrySvcNode,
	serviceInstances []openapi.RegistrySvcServiceInstance,
	deployments []*deploy.Deployment) []*deploy.Command {

	commands := []*deploy.Command{}

	for _, deployment := range deployments {
		commands = append(commands, scaleDeployment(deployment, nodes, serviceInstances)...)
	}

	for _, instance := range serviceInstances {
		commands = append(commands, checkHealthAndKill(instance)...)
	}

	return commands
}

func scaleDeployment(
	deployment *deploy.Deployment,
	nodes []openapi.RegistrySvcNode,
	serviceInstances []openapi.RegistrySvcServiceInstance,
) []*deploy.Command {
	commands := []*deploy.Command{}
	activeInstances := 0
	assignedNodes := map[string]bool{} // Tracks nodes assigned for this service

	// Count active instances and track nodes assigned to this service
	for _, instance := range serviceInstances {
		if instance.ServiceSlug == deployment.ServiceSlug {
			activeInstances++
			if instance.NodeUrl != nil {
				assignedNodes[*instance.NodeUrl] = true
			}
		}
	}

	// Scale up: Add instances if replicas required > active instances
	if activeInstances < deployment.Replicas {
		for i := activeInstances; i < deployment.Replicas; i++ {
			node := findAvailableNode(nodes, assignedNodes)
			if node != nil {
				commands = append(commands, &deploy.Command{
					Action:      "START",
					ServiceSlug: deployment.ServiceSlug,
					Node:        node.Url,
				})
				assignedNodes[*node.Url] = true // Mark this node as assigned for this service
			}
		}
	}

	// Scale down: Remove instances if replicas required < active instances
	if activeInstances > deployment.Replicas {
		for i := deployment.Replicas; i < activeInstances; i++ {
			commands = append(commands, &deploy.Command{
				Action:      "KILL",
				ServiceSlug: deployment.ServiceSlug,
			})
		}
	}

	return commands
}

func checkHealthAndKill(instance openapi.RegistrySvcServiceInstance) []*deploy.Command {
	commands := []*deploy.Command{}

	if instance.LastHeartbeat == nil {
		commands = append(commands, &deploy.Command{
			Action:      "KILL",
			ServiceSlug: instance.ServiceSlug,
			InstanceId:  &instance.Id,
		})
	}

	return commands
}

func findAvailableNode(nodes []openapi.RegistrySvcNode, assignedNodes map[string]bool) *openapi.RegistrySvcNode {
	for _, node := range nodes {
		// Check if the node is not assigned to this service and has available CPU capacity
		if !assignedNodes[*node.Url] && node.Usage != nil && node.Usage.Cpu != nil && *node.Usage.Cpu.Used < 80 {
			return &node
		}
	}
	return nil
}
