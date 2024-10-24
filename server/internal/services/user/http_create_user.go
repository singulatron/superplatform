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

	user "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// CreateUser allows an administrator to create a new user
// @ID createUser
// @Summary Create a New User
// @Description Allows an authenticated administrator to create a new user with specified details.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param request body user.CreateUserRequest true "Create User Request"
// @Success 200 {object} user.CreateUserResponse "User created successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/user [post]
func (s *UserService) CreateUser(
	w http.ResponseWriter,
	r *http.Request) {

	_, err := s.isAuthorized(r, user.PermissionUserCreate.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := user.CreateUserRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = s.createUser(req.User, req.Password, req.RoleIds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.CreateUserResponse{})
	w.Write(bs)
}
