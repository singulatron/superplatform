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
	"strings"

	user "github.com/singulatron/singulatron/localtron/services/user/types"
)

// CreateRole creates a new role
// @ID createRole
// @Summary Create a New Role
// @Description Create a new role.
// @Description <b>The role ID must be prefixed by the callers username (email).</b>
// @Description Eg. if the owner's email/username is `petstore-svc` the role should look like `petstore-svc:admin`.
// @Description The user account who creates the role will become the owner of that role, and only the owner will be able to edit the role.
// @Description
// @Description Requires the `user-svc:role:create` permission.
// @Tags User Service
// @Accept json
// @Produce json
// @Param request body user.CreateRoleRequest true "Create Role Request"
// @Success 200 {object} user.CreateRoleResponse "Role created successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role [post]
func (s *UserService) CreateRole(w http.ResponseWriter, r *http.Request) {
	rsp, err := s.isAuthorized(r, user.PermissionRoleCreate.Id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := user.CreateRoleRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ownerUsername := rsp.Email
	if !strings.HasPrefix(req.Name, ownerUsername) {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}

	role, err := s.createRole(rsp.Id, req.Name, req.Description, req.PermissionIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(user.CreateRoleResponse{
		Role: role,
	})
	w.Write(bs)
}
