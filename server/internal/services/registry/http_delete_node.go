package registryservice

import (
	"net/http"

	"github.com/singulatron/superplatform/sdk/go/datastore"
	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
)

// Delete an node
// @ID removeNode
// @Summary Remove Node
// @Description Removes a registered node by node URL. This endpoint is useful when a node is no longer available but it's still present in the database.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param id path string true "Node URL"
// @Success 204 "No Content"
// @Failure 400 {object} registry.ErrorResponse "Invalid ID"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 404 {object} registry.ErrorResponse "Service not found"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/node/{id} [delete]
func (rs *RegistryService) RemoveNode(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	nodeID := r.URL.Query().Get("id")
	if nodeID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid node ID`))
		return
	}

	err := rs.removeNodeByID(nodeID)
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

func (rs *RegistryService) removeNodeByID(id string) error {
	return rs.nodeStore.Query(datastore.Id(id)).Delete()
}
