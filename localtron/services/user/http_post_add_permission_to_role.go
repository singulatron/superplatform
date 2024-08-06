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

// AddPermissionToRole handles the addition of a permission to a role
//
// @Summary Add Permission to Role
// @Description Adds a specific permission to a role identified by roleId.
// @Description
// @Description Requires the `user-svc:permission:assign` permission.
// @Tags User Service
// @Accept json
// @Produce json
// @Param roleId path string true "Role ID"
// @Param permissionId path string true "Permission ID"
// @Success 200 {object} usertypes.CreateUserResponse
// @Failure 401 {object} usertypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} usertypes.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role/{roleId}/permission/{permissionId} [put]
func (s *UserService) AddPermissionToRole(
	w http.ResponseWriter,
	r *http.Request,
) {
	// @todo add proper permission here
	_, err := s.isAuthorized(r, usertypes.PermissionPermissionAssign.Id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user, err := s.getUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	err = s.addPermissionToRole(user.Id, vars["roleId"], vars["permissionId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.CreateUserResponse{})
	w.Write(bs)
}
