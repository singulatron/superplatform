package registryservice

import (
	"encoding/json"
	"net/http"

	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
)

// List all registered service definitions or filter by criteria
// @ID queryServiceDefinitions
// @Summary List Service Definitions
// @Description Retrieves a list of all registered service definitions or filters them by specific criteria (e.g., host, IP).
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Success 200 {object} registry.ListServiceDefinitionsResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid filters"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/services [get]
func (rs *RegistryService) ListServiceDefinitions(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	serviceDefinitions, err := rs.getServiceDefinitions(ServiceDefinitionList{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	bs, _ := json.Marshal(serviceDefinitions)
	w.Write(bs)
}

type ServiceDefinitionList struct {
}

func (rs *RegistryService) getServiceDefinitions(query ServiceDefinitionList) ([]*registry.ServiceDefinition, error) {
	serviceInstaceIs, err := rs.serviceDefinitionStore.Query().Find()
	if err != nil {
		return nil, err
	}

	serviceDefinitions := []*registry.ServiceDefinition{}
	for _, serviceDefinitionI := range serviceInstaceIs {
		serviceDefinitions = append(serviceDefinitions, serviceDefinitionI.(*registry.ServiceDefinition))
	}

	filtered := []*registry.ServiceDefinition{}
	for _, v := range serviceDefinitions {
		match := true

		if match {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}
