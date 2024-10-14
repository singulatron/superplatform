package registryservice

import (
	"encoding/json"
	"net/http"

	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
)

// Query all registered service instances or filter by criteria
// @ID queryServiceInstances
// @Summary Query Service Instances
// @Description Retrieves a list of all registered service instances or filters them by specific criteria (e.g., host, IP).
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param scheme query string false "Scheme to filter by"
// @Param ip query string false "IP to filter by"
// @Param host query string false "Host to filter by"
// @Param ip query string false "IP to filter by"
// @Param id query string false "Id to filter by"
// @Success 200 {array} registry.QueryServiceInstancesResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid filters"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/services [get]
func (rs *RegistryService) QueryServiceInstances(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	host := r.URL.Query().Get("host")
	ip := r.URL.Query().Get("ip")
	slug := r.URL.Query().Get("slug")
	id := r.URL.Query().Get("ip")
	path := r.URL.Query().Get("path")

	serviceInstances, err := rs.getServiceInstances(Query{
		Id:   id,
		Host: host,
		Slug: slug,
		IP:   ip,
		Path: path,
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

type Query struct {
	Id     string
	Slug   string
	Host   string
	IP     string
	Scheme string
	Path   string
}

func (rs *RegistryService) getServiceInstances(query Query) ([]*registry.ServiceInstance, error) {
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

		if query.Slug != "" && v.Slug != query.Slug {
			match = false
		}

		if match {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}
