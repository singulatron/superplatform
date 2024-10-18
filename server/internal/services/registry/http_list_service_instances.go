package registryservice

import (
	"encoding/json"
	"net/http"

	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
)

// List all registered service instances or filter by criteria
// @ID queryServiceInstances
// @Summary List Service Instances
// @Description Retrieves a list of all registered service instances or filters them by specific criteria (e.g., host, IP).
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param scheme query string false "Scheme to filter by"
// @Param ip query string false "IP to filter by"
// @Param host query string false "Host to filter by"
// @Param ip query string false "IP to filter by"
// @Param id query string false "Id to filter by"
// @Success 200 {object} registry.ListServiceInstancesResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid filters"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/services [get]
func (rs *RegistryService) ListServiceInstances(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	host := q.Get("host")
	ip := q.Get("ip")
	serviceSlug := q.Get("serviceSlug")
	id := q.Get("ip")
	path := q.Get("path")

	serviceInstances, err := rs.getServiceInstances(List{
		Id:          id,
		Host:        host,
		ServiceSlug: serviceSlug,
		IP:          ip,
		Path:        path,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	bs, _ := json.Marshal(serviceInstances)
	w.Write(bs)
}

type List struct {
	Id          string
	ServiceSlug string
	Host        string
	IP          string
	Scheme      string
	Path        string
}

func (rs *RegistryService) getServiceInstances(query List) ([]*registry.ServiceInstance, error) {
	serviceInstaceIs, err := rs.serviceInstanceStore.Query().Find()
	if err != nil {
		return nil, err
	}

	serviceInstances := []*registry.ServiceInstance{}
	for _, serviceInstanceI := range serviceInstaceIs {
		serviceInstances = append(serviceInstances, serviceInstanceI.(*registry.ServiceInstance))
	}

	filtered := []*registry.ServiceInstance{}
	for _, v := range serviceInstances {
		match := true

		if query.IP != "" && v.IP != query.IP {
			match = false
		}

		if query.Scheme != "" && v.Scheme != query.Scheme {
			match = false
		}

		if query.Host != "" && v.Host != query.Host {
			match = false
		}

		if query.Path != "" && v.Path != query.Path {
			match = false
		}

		if query.ServiceSlug != "" && v.ServiceSlug != query.ServiceSlug {
			match = false
		}

		if match {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}
