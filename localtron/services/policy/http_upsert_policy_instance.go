package policyservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/singulatron/singulatron/localtron/datastore"
	policy "github.com/singulatron/singulatron/localtron/services/policy/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// UpsertPolicyInstance allows a user to upsert a new policy instance
// @ID upsertPolicyInstance
// @Summary Upsert a Policy Instance
// @Description Allows user to upsert a new policy instance based on a template.
// @Tags Policy Svc
// @Accept json
// @Produce json
// @Param request body types.UpsertPolicyInstanceRequest true "Upsert Policy Instance Request"
// @Success 200 {object} policy.UpsertPolicyInstanceResponse "Policy instance upsertd successfully"
// @Failure 400 {object} policy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} policy.ErrorResponse "Unauthorized"
// @Failure 500 {object} policy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /policy-svc/instance/{instanceId} [put]
func (s *PolicyService) UpsertInstance(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := s.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", policy.PermissionInstanceEdit.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := policy.UpsertInstanceRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	req.Instance.Id = mux.Vars(r)["instanceId"]

	err = s.upsertInstance(&req.Instance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(policy.UpsertPolicyInstanceResponse{})
	w.Write(bs)
}

func (s *PolicyService) upsertInstance(instance *policy.Instance) error {
	_, found, err := s.templatesStore.Query(
		datastore.Id(instance.TemplateId),
	).FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("policy template not foundu")
	}

	return s.instancesStore.Upsert(instance)
}
