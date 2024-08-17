/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package nodeservice

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	nodetypes "github.com/singulatron/singulatron/localtron/internal/services/node/types"
	"github.com/singulatron/singulatron/localtron/router"
)

type NodeService struct {
	Hostname string

	router *router.Router
}

func NewNodeService(router *router.Router) (*NodeService, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	service := &NodeService{
		Hostname: hostname,
		router:   router,
	}

	return service, nil
}

func (ns *NodeService) Start() error {
	return ns.registerPermissions()
}

func (ns *NodeService) listNodes() ([]*nodetypes.Node, error) {
	outp, err := ns.GetNvidiaSmiOutput()
	if err != nil {
		return nil, err
	}
	gpus, err := ns.ParseNvidiaSmiOutput(outp)
	if err != nil {
		return nil, err
	}

	return []*nodetypes.Node{
		{
			Hostname: ns.Hostname,
			GPUs:     gpus,
		},
	}, nil
}

func (ns *NodeService) ParseNvidiaSmiOutput(output string) ([]*nodetypes.GPU, error) {
	reader := csv.NewReader(strings.NewReader(output))
	reader.TrimLeadingSpace = true
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "reading nvidia-smi output")
	}

	gpus := []*nodetypes.GPU{}

	for i, record := range records {
		if len(record) < 10 {
			continue
		}

		temperature, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, errors.Wrap(err, "parsing GPU temperature")
		}
		utilization, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, errors.Wrap(err, "parsing GPU utilization")
		}
		memoryTotal, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, errors.Wrap(err, "parsing GPU memory total")
		}
		memoryUsed, err := strconv.Atoi(record[4])
		if err != nil {
			return nil, errors.Wrap(err, "parsing GPU memory used")
		}
		powerUsage, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return nil, errors.Wrap(err, "parsing GPU power usage")
		}
		powerCap, err := strconv.ParseFloat(record[6], 64)
		if err != nil {
			return nil, errors.Wrap(err, "parsing GPU power cap")
		}

		gpu := nodetypes.GPU{
			Id:               fmt.Sprintf("%v:%v", ns.Hostname, strconv.Itoa(i)),
			IntraNodeId:      i,
			Name:             record[0],
			BusId:            record[8],
			Temperature:      temperature,
			PerformanceState: record[10],
			PowerUsage:       powerUsage,
			PowerCap:         powerCap,
			MemoryUsage:      memoryUsed,
			MemoryTotal:      memoryTotal,
			GPUUtilization:   utilization,
			ComputeMode:      record[9],
		}

		gpus = append(gpus, &gpu)
	}

	return gpus, nil
}

func (ns *NodeService) GetNvidiaSmiOutput() (string, error) {
	cmd := exec.Command("nvidia-smi", "--query-gpu=name,temperature.gpu,utilization.gpu,memory.total,memory.used,power.draw,power.limit,driver_version,pci.bus_id,compute_mode,pstate", "--format=csv,noheader,nounits")
	output, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("executing nvidia-smi command: %v", string(output)))
	}
	return string(output), nil
}
