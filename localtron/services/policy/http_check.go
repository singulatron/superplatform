package policyservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	policy "github.com/singulatron/singulatron/localtron/services/policy/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
	"golang.org/x/time/rate"
)

// Check records a resource access and returns if the access is allowed.
// @ID check
// @Summary Check
// @Description Check records a resource access and returns if the access is allowed.
// @Tags Policy Svc
// @Accept json
// @Produce json
// @Param request body policy.CheckRequest true "Check Request"
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

	req := policy.CheckRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	allowed, err := s.check(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(policy.CheckResponse{
		Allowed: allowed,
	})
	w.Write(bs)
}

func (s *PolicyService) check(request *policy.CheckRequest) (bool, error) {
	for _, instance := range s.instances {

		switch string(instance.TemplateId) {
		case policy.RateLimitPolicyTemplate.GetId():

			maxRequests := instance.RateLimitParameters.MaxRequests
			timeWindow, err := time.ParseDuration(instance.RateLimitParameters.TimeWindow)
			if err != nil {
				return false, err
			}

			var limiterKey string
			switch instance.RateLimitParameters.Entity {
			case policy.EntityUserID:
				limiterKey = request.UserId
			case policy.EntityIP:
				limiterKey = request.Ip
			default:
				return false, fmt.Errorf("unknown entity type")
			}

			if instance.RateLimitParameters.Scope == policy.ScopeEndpoint {
				limiterKey += ":" + request.Endpoint
			}

			s.mutex.Lock()
			limiter, exists := s.rateLimiters.Load(limiterKey)
			if !exists {
				limiter = rate.NewLimiter(rate.Every(timeWindow), maxRequests)
				s.rateLimiters.Store(limiterKey, limiter)
			}
			s.mutex.Unlock()

			if !limiter.(*rate.Limiter).Allow() {
				return false, nil
			}
		}
	}

	return true, nil
}
