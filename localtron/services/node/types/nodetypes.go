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
	ProcessDetails   []Process `json:"processDetails"`
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
