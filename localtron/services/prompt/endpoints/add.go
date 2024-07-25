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

	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Add adds a new prompt
// @Summary Add Prompt
// @Description Adds a new prompt to the prompt queue and either waits for the response (if `sync` is set to true), or returns immediately.
// @Tags prompts
// @Accept json
// @Produce json
// @Param request body prompttypes.AddPromptRequest true "Add Prompt Request"
// @Success 200 {object} prompttypes.AddPromptResponse
// @Failure 400 {object} prompttypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} prompttypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} prompttypes.ErrorResponse "Internal Server Error"
// @Router /prompt/add [post]
func Add(
	w http.ResponseWriter,
	r *http.Request,
	userService usertypes.UserServiceI,
	promptService prompttypes.PromptServiceI,
) {
	err := userService.IsAuthorized(prompttypes.PermissionPromptCreate.Id, r)
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

	req := &prompttypes.AddPromptRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	rsp, err := promptService.AddPrompt(r.Context(), req, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(rsp)
	w.Write(bs)
}
