/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package nodetypes

type Cluster struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Hostname string `json:"hostname"`
	GPUs     []*GPU `json:"gpus"`
}

type GPU struct {
	// Id Node.Hostname + IntraNodeId
	Id               string    `json:"id"`
	IntraNodeId      int       `json:"intraNodeId"`
	Name             string    `json:"name"`
	BusId            string    `json:"busId"`
	Temperature      float64   `json:"temperature"`
	PerformanceState string    `json:"performanceState"`
	PowerUsage       float64   `json:"powerUsage"`
	PowerCap         float64   `json:"powerCap"`
	MemoryUsage      int       `json:"memoryUsage"`
	MemoryTotal      int       `json:"memoryTotal"`
	GPUUtilization   float64   `json:"gpuUtilization"`
	ComputeMode      string    `json:"computeMode"`
	ProcessDetails   []Process `json:"processDetails,omitempty"`
}

type Process struct {
	Pid         int    `json:"pid"`
	ProcessName string `json:"processName"`
	MemoryUsage int    `json:"memoryUsage"`
}

type ListNodesRequest struct {
}

type ListNodesResponse struct {
	Nodes []*Node `json:"nodes"`
}
