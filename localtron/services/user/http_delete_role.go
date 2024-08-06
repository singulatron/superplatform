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
	user "github.com/singulatron/singulatron/localtron/services/user/types"
)

// DeleteRole handles the deletion of a role by role ID.
// @ID deleteRole
// @Summary Delete a Role
// @Description Delete a role based on the role ID.
// @Tags User Service
// @Accept  json
// @Produce  json
// @Param   roleId     path    string  true  "Role ID"
// @Success 200 {object} user.DeleteRoleResponse
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role/{roleId} [delete]
func (s *UserService) DeleteRole(w http.ResponseWriter, r *http.Request) {
	_, err := s.isAuthorized(r, user.PermissionRoleDelete.Id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	roleId := mux.Vars(r)["roleId"]

	err = s.deleteRole(roleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(user.DeleteRoleResponse{})
	w.Write(bs)
}
