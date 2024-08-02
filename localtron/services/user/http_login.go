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

// Login handles user authentication
// @Summary Login
// @Description Authenticates a user and returns a token.
// @Tags User Service
// @Accept json
// @Produce json
// @Param request body usertypes.LoginRequest true "Login Request"
// @Success 200 {object} usertypes.LoginResponse "Login successful"
// @Failure 400 {object} usertypes.ErrorResponse "Invalid JSON"
// @Failure 500 {object} usertypes.ErrorResponse "Internal Server Error"
// @Router /user-service/login [post]
func (s *UserService) Login(w http.ResponseWriter, r *http.Request) {
	req := usertypes.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	token, err := s.login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.LoginResponse{
		Token: token,
	})
	w.Write(bs)
}
