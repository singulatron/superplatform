package registryservice

import (
	"net/http"

	registry "github.com/singulatron/singulatron/localtron/internal/services/registry/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

// Remove a registered service instance
// @ID removeServiceInstance
// @Summary Remove Service Instance
// @Description Removes a registered service instance based on the service ID.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param id path string true "Service Instance ID"
// @Success 204 "No Content"
// @Failure 400 {object} registry.ErrorResponse "Invalid ID"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 404 {object} registry.ErrorResponse "Service not found"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/service-instance/{id} [delete]
func (rs *RegistryService) RemoveServiceInstance(
	w http.ResponseWriter,
	r *http.Request,
) {
	serviceID := r.URL.Query().Get("id")
	if serviceID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid service ID`))
		return
	}

	err := rs.removeServiceInstanceByID(serviceID)
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

func (rs *RegistryService) removeServiceInstanceByID(id string) error {
	return rs.serviceInstanceStore.Query(datastore.Id(id)).Delete()
}
