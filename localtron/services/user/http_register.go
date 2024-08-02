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

// @Summary Register a New User
// @Description Register a new user with a name, email, and password.
// @Tags User Service
// @Accept json
// @Produce json
// @Param body body usertypes.RegisterRequest true "Register Request"
// @Success 200 {object} usertypes.RegisterResponse
// @Failure 400 {object} usertypes.ErrorResponse "Invalid JSON"
// @Failure 500 {object} usertypes.ErrorResponse "Internal Server Error"
// @Router /user-service/register [post]
func (s *UserService) Register(w http.ResponseWriter, r *http.Request) {
	req := usertypes.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = s.createUser(&usertypes.User{
		Name:  req.Name,
		Email: req.Email,
	}, req.Password, []string{usertypes.RoleUser.Id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.RegisterResponse{})
	w.Write(bs)
}
