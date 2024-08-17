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

	user "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// GetPermissions handles the retrieval of permissions based on the role ID.
// @ID getPermissionsByRole
// @Summary Get Permissions by Role
// @Description Retrieve permissions associated with a specific role ID.
// @Tags User Svc
// @Accept  json
// @Produce  json
// @Param   roleId     path    string     true        "Role ID"
// @Success 200 {object} user.GetPermissionsResponse
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role/{roleId}/permissions [get]
func (s *UserService) GetPermissions(
	w http.ResponseWriter,
	r *http.Request) {
	_, err := s.isAuthorized(r, user.PermissionRoleView.Id, nil, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	permissions, err := s.getPermissions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(user.GetPermissionsResponse{
		Permissions: permissions,
	})
	w.Write(bs)
}
