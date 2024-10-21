package registryservice

import (
	"net/http"

	"github.com/singulatron/superplatform/sdk/go/datastore"
	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
)

// Delete a service definition
// @ID deleteDefinition
// @Summary Delete Definition
// @Description Deletes a registered definition by ID.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param id path string true "Definition ID"
// @Success 204 "No Content"
// @Failure 400 {object} registry.ErrorResponse "Invalid ID"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 404 {object} registry.ErrorResponse "Service not found"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/definition/{id} [delete]
func (rs *RegistryService) DeleteDefinition(
	w http.ResponseWriter,
	r *http.Request,
) {
	serviceID := r.URL.Query().Get("id")
	if serviceID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid definition ID`))
		return
	}

	err := rs.deleteDefinitionByID(serviceID)
	if err != nil {
		if err == registry.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`Definition not found`))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rs *RegistryService) deleteDefinitionByID(id string) error {
	return rs.definitionStore.Query(datastore.Id(id)).Delete()
}
