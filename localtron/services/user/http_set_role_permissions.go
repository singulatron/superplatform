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

	"github.com/gorilla/mux"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// SetRolePermissions sets the permissions for a specified role
// @Summary Set Role Permissions
// @Description Set permissions for a specified role. The caller can add permissions it owns to any role.
// @Description If the caller tries to add a permission it doesn't own to a role, `StatusBadRequest` will be returned.
// @Tags User Service
// @Accept json
// @Produce json
// @Param roleId path string true "Role ID"
// @Param body body usertypes.SetRolePermissionsRequest true "Set Role Permissions Request"
// @Success 200 {object} usertypes.SetRolePermissionsResponse
// @Failure 400 {object} usertypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} usertypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} usertypes.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role/{roleId}/permissions [put]
func (s *UserService) SetRolePermissions(
	w http.ResponseWriter,
	r *http.Request) {
	user, err := s.isAuthorized(r, usertypes.PermissionRoleEdit.Id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := usertypes.SetRolePermissionsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	roleId := mux.Vars(r)["roleId"]

	err = s.setRolePermissions(user.Id, roleId, req.PermissionIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.SetRolePermissionsResponse{})
	w.Write(bs)
}
