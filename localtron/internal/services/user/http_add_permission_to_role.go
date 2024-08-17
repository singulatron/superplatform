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

// AddPermissionToRole handles the addition of a permission to a role
// @ID addPermissionToRole
// @Summary Add Permission to Role
// @Description Adds a specific permission to a role identified by roleId.
// @Description
// @Description Requires the `user-svc:permission:assign` permission.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param roleId path string true "Role ID"
// @Param permissionId path string true "Permission ID"
// @Success 200 {object} user.CreateUserResponse
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role/{roleId}/permission/{permissionId} [put]
func (s *UserService) AddPermissionToRole(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	// @todo add proper permission here
	_, err := s.isAuthorized(r, user.PermissionPermissionAssign.Id, nil, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	usr, err := s.getUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	err = s.addPermissionToRole(usr.Id, vars["roleId"], vars["permissionId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(user.CreateUserResponse{})
	w.Write(bs)
}
