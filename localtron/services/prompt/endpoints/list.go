/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptendpoints

import (
	"encoding/json"
	"net/http"

	"github.com/singulatron/singulatron/localtron/datastore"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// List lists prompts
// @Summary List prompts
// @Description List prompts with specific criteria
// @Tags prompts
// @Accept json
// @Produce json
// @Param request body prompttypes.ListPromptsRequest true "List Prompts Request"
// @Success 200 {object} prompttypes.ListPromptsResponse
// @Failure 400 {object} prompttypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} prompttypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} prompttypes.ErrorResponse "Internal Server Error"
// @Router /prompt/list [post]
func List(
	w http.ResponseWriter,
	r *http.Request,
	userService usertypes.UserServiceI,
	promptService prompttypes.PromptServiceI,
) {
	err := userService.IsAuthorized(prompttypes.PermissionPromptView.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user, found, err := userService.GetUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !found {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	req := prompttypes.ListPromptsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	options := &prompttypes.ListPromptOptions{
		Query: req.Query,
	}
	if options.Query == nil {
		options.Query = &datastore.Query{}
	}

	if !options.Query.HasFieldCondition("status") {
		options.Query.Conditions = append(options.Query.Conditions, datastore.Equal(datastore.Field("status"), []prompttypes.PromptStatus{
			prompttypes.PromptStatusRunning,
			prompttypes.PromptStatusScheduled,
		}))
	}
	if options.Query.Limit == 0 {
		options.Query.Limit = 20
	}

	prompts, count, err := promptService.ListPrompts(options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := range prompts {
		if prompts[i].UserId != user.Id {
			// do not let users see other peoples promtps,
			// not even if they are admins
			// eg. imagine a sysadmin looking at the CEO's prompt
			prompts[i].Prompt = ""
		}
	}

	response := prompttypes.ListPromptsResponse{
		Prompts: prompts,
		Count:   count,
	}
	if len(prompts) >= 20 {
		response.After = prompts[len(prompts)-1].CreatedAt
	}

	bs, _ := json.Marshal(response)
	w.Write(bs)
}
