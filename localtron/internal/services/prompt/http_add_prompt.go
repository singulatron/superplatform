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

	prompt "github.com/singulatron/singulatron/localtron/internal/services/prompt/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// Add a new prompt
// @ID addPrompt
// @Summary Add Prompt
// @Description Adds a new prompt to the prompt queue and either waits for the response (if `sync` is set to true), or returns immediately.
// @Tags Prompt Svc
// @Accept json
// @Produce json
// @Param request body prompt.AddPromptRequest true "Add Prompt Request"
// @Success 200 {object} prompt.AddPromptResponse
// @Failure 400 {object} prompt.ErrorResponse "Invalid JSON"
// @Failure 401 {object} prompt.ErrorResponse "Unauthorized"
// @Failure 500 {object} prompt.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /prompt-svc/prompt [post]
func (p *PromptService) Add(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := p.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", prompt.PermissionPromptCreate.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	if !rsp.Authorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	req := &prompt.AddPromptRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	prsp, err := p.addPrompt(r.Context(), req, rsp.User.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(prsp)
	w.Write(bs)
}
