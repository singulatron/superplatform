package registryservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	registry "github.com/singulatron/singulatron/localtron/internal/services/registry/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// Register a new service
// @ID registerService
// @Summary Register Service. Idempotent.
// @Description Registers a new service, associating it with a domain or IP, identified from the bearer token.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param request body registry.RegisterServiceRequest true "Register Service Request"
// @Success 201 {object} registry.RegisterServiceResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid JSON"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/service-instance [post]
func (rs *RegistryService) RegisterService(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	authRsp := &usertypes.IsAuthorizedResponse{}
	err := rs.router.AsRequestMaker(r).Post(r.Context(), "user-svc", "/permission/register-service", &usertypes.IsAuthorizedRequest{}, authRsp)
	if err != nil || !authRsp.Authorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	req := &registry.RegisterServiceInstanceRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err := rs.registerService(r.Context(), req, authRsp.User.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (rs *RegistryService) registerService(req *registry.RegisterServiceRequest) error {
	if req.URL == "" {
		if req.Scheme == "" {
			return fmt.Errorf("scheme is mandatory when full URL is not provided")
		}
		if req.Host == "" && req.IP == "" {
			return fmt.Errorf("host or IP is mandatory when full URL is not provided")
		}
	}

	rs.serviceInstanceStore.Upsert()
}
