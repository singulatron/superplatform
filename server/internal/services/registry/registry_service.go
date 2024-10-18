/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package registryservice

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/datastore"
	"github.com/singulatron/superplatform/sdk/go/lock"
	"github.com/singulatron/superplatform/sdk/go/logger"
	"github.com/singulatron/superplatform/sdk/go/router"
	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
)

type RegistryService struct {
	URL              string
	AvailabilityZone string
	Region           string

	router *router.Router
	lock   lock.DistributedLock

	credentialStore        datastore.DataStore
	serviceDefinitionStore datastore.DataStore
	serviceInstanceStore   datastore.DataStore
	nodeStore              datastore.DataStore
}

func NewRegistryService(
	address string,
	az string,
	region string,
	router *router.Router,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any,
	) (datastore.DataStore, error)) (*RegistryService, error) {

	nodeUrl := address
	var err error

	if nodeUrl == "" {
		nodeUrl, err = os.Hostname()
		if err != nil {
			return nil, err
		}
		nodeUrl = fmt.Sprintf("%v:%v", nodeUrl, "58231")
	}

	credentialStore, err := datastoreFactory("registrySvcCredentials", &sdk.Credential{})
	if err != nil {
		return nil, err
	}
	serviceInstanceStore, err := datastoreFactory("registrySvcServiceInstances", &registry.ServiceInstance{})
	if err != nil {
		return nil, err
	}
	serviceDefinitionStore, err := datastoreFactory("registrySvcServiceInstances", &registry.ServiceDefinition{})
	if err != nil {
		return nil, err
	}
	nodeStore, err := datastoreFactory("registrySvcNodes", &registry.Node{})
	if err != nil {
		return nil, err
	}

	service := &RegistryService{
		URL:                    nodeUrl,
		router:                 router,
		lock:                   lock,
		credentialStore:        credentialStore,
		serviceDefinitionStore: serviceDefinitionStore,
		serviceInstanceStore:   serviceInstanceStore,
		nodeStore:              nodeStore,
		AvailabilityZone:       az,
		Region:                 region,
	}

	return service, nil
}

func (ns *RegistryService) Start() error {
	go ns.nodeHeartbeat()

	ctx := context.Background()
	ns.lock.Acquire(ctx, "registry-svc-start")
	defer ns.lock.Release(ctx, "registry-svc-start")

	token, err := sdk.RegisterService("registry-svc", "Registry Service", ns.router, ns.credentialStore)
	if err != nil {
		return err
	}
	ns.router = ns.router.SetBearerToken(token)

	return ns.registerPermissions()
}

func (ns *RegistryService) nodeHeartbeat() {
	first := true
	for {
		if !first {
			time.Sleep(2 * time.Second)
			first = false
		}

		node := registry.Node{
			URL:              ns.URL,
			AvailabilityZone: ns.AvailabilityZone,
			Region:           ns.Region,
			LastHeartbeat:    time.Now(),
		}

		usage, err := getResourceUsage()
		if err != nil {
			logger.Warn("Failed to get resource usage", slog.Any("error", err))
		}

		node.Usage = usage

		// @todo detect non-nvidia gpus
		outp, err := ns.getNvidiaSmiOutput()
		if err != nil {
			logger.Warn("Failed to get smi output", slog.Any("error", err))
		} else {
			gpus, err := ns.ParseNvidiaSmiOutput(outp)
			if err != nil {
				logger.Warn("Failed to parse smi output", slog.Any("error", err))
			} else {
				node.GPUs = gpus
			}
		}

		err = ns.nodeStore.Upsert(node)
		if err != nil {
			logger.Error("Failed to save node", err)
		}

	}
}

func (ns *RegistryService) ParseNvidiaSmiOutput(output string) ([]*registry.GPU, error) {
	reader := csv.NewReader(strings.NewReader(output))
	reader.TrimLeadingSpace = true
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "reading nvidia-smi output")
	}

	gpus := []*registry.GPU{}

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

		gpu := registry.GPU{
			Id:               fmt.Sprintf("%v:%v", ns.URL, strconv.Itoa(i)),
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

func (ns *RegistryService) getNvidiaSmiOutput() (string, error) {
	cmd := exec.Command("nvidia-smi", "--query-gpu=name,temperature.gpu,utilization.gpu,memory.total,memory.used,power.draw,power.limit,driver_version,pci.bus_id,compute_mode,pstate", "--format=csv,noheader,nounits")
	output, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("executing nvidia-smi command: %v", string(output)))
	}
	return string(output), nil
}

func getResourceUsage() (registry.ResourceUsage, error) {
	var usage registry.ResourceUsage

	// Get CPU usage
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		log.Println("Error getting CPU usage:", err)
		return usage, err
	}
	usage.CPU.Percent = cpuPercent[0]            // Take the first element since it returns a slice
	usage.CPU.Used = uint64(cpuPercent[0]) * 100 // Assume total is 100% for simplification
	usage.CPU.Total = 100                        // This should be replaced with actual total if available

	// Get Memory usage
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Println("Error getting memory usage:", err)
		return usage, err
	}
	usage.Memory.Used = memInfo.Used
	usage.Memory.Total = memInfo.Total
	usage.Memory.Percent = memInfo.UsedPercent

	// Get Disk usage
	diskUsage, err := disk.Usage("/")
	if err != nil {
		log.Println("Error getting disk usage:", err)
		return usage, err
	}
	usage.Disk.Used = diskUsage.Used
	usage.Disk.Total = diskUsage.Total
	usage.Disk.Percent = diskUsage.UsedPercent

	return usage, nil
}
