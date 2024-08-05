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

// CreateRole creates a new role
// @Summary Create a New Role
// @Description Create a new role.
// @Description
// @Description Requires the `role.create` permission.
// @Tags User Service
// @Accept json
// @Produce json
// @Param request body usertypes.CreateRoleRequest true "Create Role Request"
// @Success 200 {object} usertypes.CreateRoleResponse "Role created successfully"
// @Failure 400 {object} usertypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} usertypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} usertypes.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role [post]
func (s *UserService) CreateRole(w http.ResponseWriter, r *http.Request) {
	_, err := s.isAuthorized(r, usertypes.PermissionRoleCreate.Id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := usertypes.CreateRoleRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	role, err := s.createRole(req.Name, req.Description, req.PermissionIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.CreateRoleResponse{
		Role: role,
	})
	w.Write(bs)
}
