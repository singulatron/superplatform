package registryservice

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/singulatron/superplatform/sdk/go/datastore"
	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// Delete an node
// @ID deleteNode
// @Summary Delete Node
// @Description Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it's still present in the database.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param url path string true "Node URL"
// @Success 204 "No Content"
// @Failure 400 {object} registry.ErrorResponse "Invalid ID"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 404 {object} registry.ErrorResponse "Service not found"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/node/{url} [delete]
func (rs *RegistryService) DeleteNode(
	w http.ResponseWriter,
	r *http.Request,
) {

	rsp := &usertypes.IsAuthorizedResponse{}
	err := rs.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", registry.PermissionNodeDelete.Id), &usertypes.IsAuthorizedRequest{
		SlugsGranted: []string{"deploy-svc"},
	}, rsp)
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

	vars := mux.Vars(r)
	nodeURL, err := url.PathUnescape(vars["url"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if nodeURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid node ID`))
		return
	}

	err = rs.deleteNodeByURL(nodeURL)
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

func (rs *RegistryService) deleteNodeByURL(ur string) error {
	return rs.nodeStore.Query(datastore.Equals([]string{"url"}, ur)).Delete()
}