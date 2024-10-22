package registryservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// List all registered definitions or filter by criteria
// @ID listDefinitions
// @Summary List Definitions
// @Description Retrieves a list of all definitions or filters them by specific criteria.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Success 200 {object} registry.ListDefinitionsResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid filters"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/definitions [get]
func (rs *RegistryService) ListDefinitions(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := rs.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", registry.PermissionDefinitionView.Id), &usertypes.IsAuthorizedRequest{
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

	definitions, err := rs.getDefinitions(DefinitionList{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	bs, _ := json.Marshal(registry.ListDefinitionsResponse{
		Definitions: definitions,
	})
	w.Write(bs)
}

type DefinitionList struct {
}

func (rs *RegistryService) getDefinitions(query DefinitionList) ([]*registry.Definition, error) {
	serviceInstaceIs, err := rs.definitionStore.Query().Find()
	if err != nil {
		return nil, err
	}

	definitions := []*registry.Definition{}
	for _, definitionI := range serviceInstaceIs {
		definitions = append(definitions, definitionI.(*registry.Definition))
	}

	filtered := []*registry.Definition{}
	for _, v := range definitions {
		match := true

		if match {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}
