package registryservice

import (
	"net/http"

	"github.com/singulatron/superplatform/sdk/go/datastore"
	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
)

// Delete a service definition
// @ID deleteServiceDefinition
// @Summary Delete Service Definition
// @Description Deletes a registered service definition based on the service ID.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param id path string true "Service Definition ID"
// @Success 204 "No Content"
// @Failure 400 {object} registry.ErrorResponse "Invalid ID"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 404 {object} registry.ErrorResponse "Service not found"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/service-definition/{id} [delete]
func (rs *RegistryService) DeleteServiceDefinition(
	w http.ResponseWriter,
	r *http.Request,
) {
	serviceID := r.URL.Query().Get("id")
	if serviceID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid service ID`))
		return
	}

	err := rs.deleteServiceDefinitionByID(serviceID)
	if err != nil {
		if err == registry.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`Service not found`))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rs *RegistryService) deleteServiceDefinitionByID(id string) error {
	return rs.serviceDefinitionStore.Query(datastore.Id(id)).Delete()
}
