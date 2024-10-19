/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package deploy_svc

type ErrorResponse struct {
	Error string `json:"error"`
}

type Deployment struct {
	ID            string             `json:"id,omitempty" example:"depl_dbOdi5eLQK"`                      // ID of the deployment (e.g., "depl_dbOdi5eLQK")
	ServiceSlug   string             `json:"serviceSlug,omitempty" example:"user-svc" binding:"required"` // The User Svc slug of the service that is being deployed.
	Replicas      int                `json:"replicas,omitempty"`                                          // Number of container instances to run
	Strategy      DeploymentStrategy `json:"strategy,omitempty"`                                          // Deployment strategy (e.g., rolling update)
	Resources     ResourceLimits     `json:"resources,omitempty"`                                         // Resource requirements for each replica
	AutoScaling   *AutoScalingConfig `json:"autoScaling,omitempty"`                                       // Optional: Auto-scaling rules
	TargetRegions []TargetRegion     `json:"targetRegions,omitempty"`                                     // Target deployment regions or clusters
}

func (d Deployment) GetId() string {
	return d.ID
}

type DeploymentStrategy struct {
	Type           StrategyType `json:"type,omitempty"`           // Deployment strategy type (RollingUpdate, Recreate, etc.)
	MaxUnavailable int          `json:"maxUnavailable,omitempty"` // Max unavailable replicas during update
	MaxSurge       int          `json:"maxSurge,omitempty"`       // Max extra replicas during update
}

type StrategyType string

const (
	StrategyRollingUpdate StrategyType = "RollingUpdate"
	StrategyRecreate      StrategyType = "Recreate"
)

type ResourceLimits struct {
	CPU    string `json:"cpu,omitempty"`    // CPU limit, e.g., "500m" for 0.5 cores
	Memory string `json:"memory,omitempty"` // Memory limit, e.g., "128Mi"
	VRAM   string `json:"vram,omitempty"`   // Optional: GPU VRAM requirement, e.g., "48GB"
}

type AutoScalingConfig struct {
	MinReplicas  int `json:"minReplicas,omitempty"`  // Minimum number of replicas to run
	MaxReplicas  int `json:"maxReplicas,omitempty"`  // Maximum number of replicas to run
	CPUThreshold int `json:"cpuThreshold,omitempty"` // CPU usage threshold for scaling (as a percentage)

	// @todo need to measure and describe the utilization of GPUs
}

type TargetRegion struct {
	Cluster string `json:"cluster,omitempty"` // Cluster or node where service should be deployed (e.g., "us-west1", "local-docker")
	Zone    string `json:"zone,omitempty"`    // Optional: Specific zone for the deployment
}

type ListDeploymentsRequest struct {
}

type ListDeploymentsResponse struct {
	Deployments []*Deployment `json:"deployments,omitempty"`
}

type SaveDeploymentRequest struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

type SaveDeploymentResponse struct {
}
