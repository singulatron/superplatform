package registryservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	registry "github.com/singulatron/singulatron/localtron/internal/services/registry/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// Register a new service instance
// @ID registerServiceInstance
// @Summary Register Service Instance
// @Description Registers a new service instance, associating an service instance address with a slug acquired from the bearer token.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param request body registry.RegisterServiceInstanceRequest true "Register Service Instance Request"
// @Success 201 {object} registry.RegisterServiceInstanceResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid JSON"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/service-instance [post]
func (rs *RegistryService) RegisterServiceInstance(
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

	err = rs.registerService(req, authRsp.User.Slug)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{}`))
}

func (rs *RegistryService) registerService(req *registry.RegisterServiceInstanceRequest, userSlug string) error {
	if req.URL == "" {
		if req.Scheme == "" {
			return fmt.Errorf("scheme is mandatory when full URL is not provided")
		}
		if req.Host == "" && req.IP == "" {
			return fmt.Errorf("host or IP is mandatory when full URL is not provided")
		}

		host := req.Host
		if host == "" {
			host = req.IP
		}

		req.URL = fmt.Sprintf("%s://%s", req.Scheme, host)

		if req.Port != 0 {
			req.URL = fmt.Sprintf("%s:%d", req.URL, req.Port)
		}
		if req.Path != "" {
			req.URL = fmt.Sprintf("%s%s", req.URL, req.Path)
		}
	}

	inst := &registry.ServiceInstance{
		URL:    req.URL,
		Scheme: req.Scheme,
		Host:   req.Host,
		IP:     req.IP,
		Path:   req.Path,
		Slug:   req.Slug,
	}
	inst.ID = inst.DeriveID()

	return rs.serviceInstanceStore.Upsert(inst)
}
