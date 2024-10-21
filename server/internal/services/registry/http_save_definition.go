package registryservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// Register a new definition
// @ID saveDefinition
// @Summary Register a Definition
// @Description Registers a new definition, associating an definition address with a slug acquired from the bearer token.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param request body registry.SaveDefinitionRequest true "Register Service Definition Request"
// @Success 201 {object} registry.SaveDefinitionResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid JSON"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/definition [put]
func (rs *RegistryService) SaveDefinition(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := rs.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", registry.PermissionNodeView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
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

	req := &registry.SaveDefinitionRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = rs.saveServiceDefinition(req, rsp.User.Slug)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{}`))
}

func (rs *RegistryService) saveServiceDefinition(req *registry.SaveDefinitionRequest, userSlug string) error {
	return rs.definitionStore.Upsert(req.Definition)
}
