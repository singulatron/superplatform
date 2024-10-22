package registryservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	registry "github.com/singulatron/superplatform/server/internal/services/registry/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// List all instances or filter by criteria
// @ID listInstances
// @Summary List Service Instances
// @Description Retrieves a list of all instances or filters them by specific criteria (e.g., host, IP).
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param scheme query string false "Scheme to filter by"
// @Param ip query string false "IP to filter by"
// @Param deploymentId query string false "Deployment ID to filter by"
// @Param host query string false "Host to filter by"
// @Param ip query string false "IP to filter by"
// @Param id query string false "Id to filter by"
// @Success 200 {object} registry.ListInstancesResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid filters"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/instances [get]
func (rs *RegistryService) ListInstances(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := rs.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", registry.PermissionInstanceView.Id), &usertypes.IsAuthorizedRequest{
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

	q := r.URL.Query()
	host := q.Get("host")
	ip := q.Get("ip")
	deploymentId := q.Get("deploymentId")
	id := q.Get("ip")
	path := q.Get("path")

	instances, err := rs.getInstances(List{
		Id:           id,
		Host:         host,
		DeploymentId: deploymentId,
		IP:           ip,
		Path:         path,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	bs, _ := json.Marshal(registry.ListInstancesResponse{
		Instances: instances,
	})
	w.Write(bs)
}

type List struct {
	Id           string
	DeploymentId string
	ServiceSlug  string
	Host         string
	IP           string
	Scheme       string
	Path         string
}

func (rs *RegistryService) getInstances(query List) ([]*registry.Instance, error) {
	instanceIs, err := rs.instanceStore.Query().Find()
	if err != nil {
		return nil, err
	}

	instances := []*registry.Instance{}
	for _, instanceI := range instanceIs {
		instances = append(instances, instanceI.(*registry.Instance))
	}

	filtered := []*registry.Instance{}
	for _, v := range instances {
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
		if query.DeploymentId != "" && v.DeploymentId != query.DeploymentId {
			match = false
		}

		if match {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}
