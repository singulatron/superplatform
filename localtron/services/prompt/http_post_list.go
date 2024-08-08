/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/singulatron/singulatron/localtron/datastore"
	prompt "github.com/singulatron/singulatron/localtron/services/prompt/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// List lists prompts
// @ID getPrompts
// @Summary List Prompts
// @Description List prompts that satisfy a query.
// @Tags Prompt Svc
// @Accept json
// @Produce json
// @Param request body prompt.ListPromptsRequest false "List Prompts Request"
// @Success 200 {object} prompt.ListPromptsResponse
// @Failure 400 {object} prompt.ErrorResponse "Invalid JSON"
// @Failure 401 {object} prompt.ErrorResponse "Unauthorized"
// @Failure 500 {object} prompt.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /prompt-svc/prompts [post]
func (p *PromptService) GetPrompts(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := p.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", prompt.PermissionPromptView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := prompt.ListPromptsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	options := &prompt.ListPromptOptions{
		Query: req.Query,
	}
	if options.Query == nil {
		options.Query = &datastore.Query{}
	}

	if !options.Query.HasFieldCondition("status") {
		options.Query.Conditions = append(options.Query.Conditions, datastore.Equal(datastore.Field("status"), []prompt.PromptStatus{
			prompt.PromptStatusRunning,
			prompt.PromptStatusScheduled,
		}))
	}
	if options.Query.Limit == 0 {
		options.Query.Limit = 20
	}

	prompts, count, err := p.listPrompts(options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := range prompts {
		if prompts[i].UserId != rsp.User.Id {
			// do not let users see other peoples promtps,
			// not even if they are admins
			// eg. imagine a sysadmin looking at the CEO's prompt
			prompts[i].Prompt = ""
		}
	}

	response := prompt.ListPromptsResponse{
		Prompts: prompts,
		Count:   count,
	}
	if len(prompts) >= 20 {
		response.After = prompts[len(prompts)-1].CreatedAt
	}

	bs, _ := json.Marshal(response)
	w.Write(bs)
}
