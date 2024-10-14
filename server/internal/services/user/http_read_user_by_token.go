/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"encoding/json"
	"net/http"

	sdk "github.com/singulatron/superplatform/sdk/go"
	user "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// @ID readUserByToken
// @Summary Read User by Token
// @Description Retrieve user information based on an authentication token.
// @Tags User Svc
// @Accept json
// @Produce json
// @Success 200 {object} user.ReadUserByTokenResponse
// @Failure 400 {object} user.ErrorResponse "Token Missing"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/user/by-token [post]
func (s *UserService) ReadUserByToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token, exists := sdk.TokenFromRequest(r)
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Token Missing`))
		return
	}

	usr, err := s.readUserByToken(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	usr.PasswordHash = ""

	orgs, activeOrgId, err := s.getUserOrganizations(usr.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.ReadUserByTokenResponse{
		User:                 usr,
		Organizations:        orgs,
		ActiveOrganizationId: activeOrgId,
	})
	w.Write(bs)
}
