package policyservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	policy "github.com/singulatron/singulatron/localtron/services/policy/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Check records a resource access and returns if the access is allowed.
// @ID check
// @Summary Check
// @Description Check records a resource access and returns if the access is allowed.
// @Tags Policy Svc
// @Accept json
// @Produce json
// @Param request body types.CheckRequest true "Check Request"
// @Success 200 {object} policy.CheckResponse "Checked successfully"
// @Failure 400 {object} policy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} policy.ErrorResponse "Unauthorized"
// @Failure 500 {object} policy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /policy-svc/check [post]
func (s *PolicyService) Check(
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

func (s *PolicyService) check(template *policy.Template) error {
	return s.templatesStore.Upsert(template)
}
