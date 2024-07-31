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
	"net/http"

	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Remove removes a prompt
// @Summary Remove Prompt
// @Description Remove a prompt by ID.
// @Tags prompts
// @Accept json
// @Produce json
// @Param request body prompttypes.RemovePromptRequest true "Remove Prompt Request"
// @Success 200 {object} prompttypes.RemovePromptResponse "{}"
// @Failure 400 {object} prompttypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} prompttypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} prompttypes.ErrorResponse "Internal Server Error"
// @Router /prompt/remove [post]
func (p *PromptService) RemovePrompt(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := p.router.AsRequestMaker(r).Post(r.Context(), "user", "/is-authorized", &usertypes.IsAuthorizedRequest{
		PermissionId: prompttypes.PermissionPromptCreate.Id,
	}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := &prompttypes.RemovePromptRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = p.removePrompt(req.PromptId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{}`))
}
