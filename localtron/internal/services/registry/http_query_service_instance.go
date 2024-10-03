package registryservice

import (
	"encoding/json"
	"net/http"

	registry "github.com/singulatron/singulatron/localtron/internal/services/registry/types"
)

// Query all registered services or filter by criteria
// @ID queryServices
// @Summary Query Services
// @Description Retrieves a list of all registered services or filters them by specific criteria (e.g., domain, IP).
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param domain query string false "Domain to filter by"
// @Param ip query string false "IP to filter by"
// @Success 200 {array} registry.QueryServiceResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid filters"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/services [get]
func (rs *RegistryService) QueryServices(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	// Extract filters from query params (if any)
	domain := r.URL.Query().Get("domain")
	ip := r.URL.Query().Get("ip")

	// Call internal function to query services
	services, err := rs.getServices(r.Context(), domain, ip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(registry.ErrorResponse{Message: err.Error()})
		return
	}

	// Return the list of services
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)
}
