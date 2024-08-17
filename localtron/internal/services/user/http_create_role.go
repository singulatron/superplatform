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

	user "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// CreateRole creates a new role
// @ID createRole
// @Summary Create a New Role
// @Description Create a new role.
// @Description <b>The role ID must be prefixed by the callers username (email).</b>
// @Description Eg. if the owner's slug is `petstore-svc` the role should look like `petstore-svc:admin`.
// @Description The user account who creates the role will become the owner of that role, and only the owner will be able to edit the role.
// @Description
// @Description Requires the `user-svc:role:create` permission.
// @Tags User Svc
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
	w.Header().Set("Content-Type", "application/json")

	rsp, err := s.isAuthorized(r, user.PermissionRoleCreate.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := user.CreateRoleRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	ownerUsername := rsp.Slug
	if !strings.HasPrefix(req.Name, ownerUsername) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid prefix`))
		return
	}

	role, err := s.createRole(rsp.Id, req.Name, req.Description, req.PermissionIds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.CreateRoleResponse{
		Role: role,
	})
	w.Write(bs)
}
