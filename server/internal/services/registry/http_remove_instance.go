package registryservice

import (
	"fmt"
	"net/http"

	"github.com/singulatron/superplatform/sdk/go/datastore"
	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// Remove an instance
// @ID removeInstance
// @Summary Remove Instance
// @Description Removes a registered instance by ID.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param id path string true "Instance ID"
// @Success 204 "No Content"
// @Failure 400 {object} registry.ErrorResponse "Invalid ID"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 404 {object} registry.ErrorResponse "Service not found"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/instance/{id} [delete]
func (rs *RegistryService) RemoveInstance(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := rs.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", registry.PermissionInstanceDelete.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !rsp.Authorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	instanceID := r.URL.Query().Get("id")
	if instanceID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid instance ID`))
		return
	}

	err = rs.removeInstanceByID(instanceID)
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

func (rs *RegistryService) removeInstanceByID(id string) error {
	return rs.instanceStore.Query(datastore.Id(id)).Delete()
}
