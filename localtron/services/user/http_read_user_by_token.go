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

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// @Summary Read User by Token
// @Description Retrieve user information based on an authentication token.
// @Tags User Service
// @Accept json
// @Produce json
// @Param body body usertypes.ReadUserByTokenRequest true "Read User By Token Request"
// @Success 200 {object} usertypes.ReadUserByTokenResponse
// @Failure 400 {object} usertypes.ErrorResponse "Invalid JSON"
// @Failure 500 {object} usertypes.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/user/by-token [post]
func (s *UserService) ReadUserByToken(w http.ResponseWriter, r *http.Request) {
	req := usertypes.ReadUserByTokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, err := s.readUserByToken(req.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.PasswordHash = ""

	bs, _ := json.Marshal(usertypes.ReadUserByTokenResponse{
		User: user,
	})
	w.Write(bs)
}
