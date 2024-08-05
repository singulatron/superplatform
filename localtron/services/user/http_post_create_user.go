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

// CreateUser allows an administrator to create a new user
// @Summary Create a New User
// @Description Allows an authenticated administrator to create a new user with specified details.
// @Tags User Service
// @Accept json
// @Produce json
// @Param request body usertypes.CreateUserRequest true "Create User Request"
// @Success 200 {object} usertypes.CreateUserResponse "User created successfully"
// @Failure 400 {object} usertypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} usertypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} usertypes.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-service/user [post]
func (s *UserService) CreateUser(
	w http.ResponseWriter,
	r *http.Request) {
	_, err := s.isAuthorized(r, usertypes.PermissionUserCreate.Id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := usertypes.CreateUserRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = s.createUser(req.User, req.Password, req.RoleIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.CreateUserResponse{})
	w.Write(bs)
}
