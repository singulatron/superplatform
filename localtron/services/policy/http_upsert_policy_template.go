package policyservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	policy "github.com/singulatron/singulatron/localtron/services/policy/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// UpsertPolicyTemplate allows a user to upsert a new policy template
// @ID upsertPolicyTemplate
// @Summary Upsert a Policy Template
// @Description Allows user to upsert a new policy template.
// @Tags Policy Svc
// @Accept json
// @Produce json
// @Param request body types.UpsertPolicyTemplateRequest true "Upsert Policy Template Request"
// @Success 200 {object} policy.UpsertPolicyTemplateResponse "Policy template upsertd successfully"
// @Failure 400 {object} policy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} policy.ErrorResponse "Unauthorized"
// @Failure 500 {object} policy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /policy-svc/template/{templateId} [put]
func (s *PolicyService) UpsertTemplate(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := s.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", policy.PermissionTemplateEdit.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := policy.UpsertTemplateRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	req.Template.Id = mux.Vars(r)["templateId"]

	err = s.upsertTemplate(&req.Template)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(policy.UpsertTemplateResponse{})
	w.Write(bs)
}

func (s *PolicyService) upsertTemplate(template *policy.Template) error {
	return s.templatesStore.Upsert(template)
}
