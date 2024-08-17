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
	user "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// SetRolePermissions sets the permissions for a specified role
// @ID setRolePermission
// @Summary Set Role Permissions
// @Description Set permissions for a specified role. The caller can add permissions it owns to any role.
// @Description If the caller tries to add a permission it doesn't own to a role, `StatusBadRequest` will be returned.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param roleId path string true "Role ID"
// @Param body body user.SetRolePermissionsRequest true "Set Role Permissions Request"
// @Success 200 {object} user.SetRolePermissionsResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role/{roleId}/permissions [put]
func (s *UserService) SetRolePermissions(
	w http.ResponseWriter,
	r *http.Request) {
	usr, err := s.isAuthorized(r, user.PermissionRoleEdit.Id, nil, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := user.SetRolePermissionsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	roleId := mux.Vars(r)["roleId"]

	err = s.overwriteRolePermissions(usr.Id, roleId, req.PermissionIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(user.SetRolePermissionsResponse{})
	w.Write(bs)
}
