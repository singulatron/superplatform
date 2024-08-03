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

// UpsertPermission handles the creation or update of a permission
//
// @Summary Upsert a Permission
// @Description Creates or updates a permission.
// @Description
// @Description Requires the `permission.create` permission.
// @Tags User Service
// @Accept json
// @Produce json
// @Param permissionId path string true "Permission ID"
// @Param requestBody body usertypes.UpserPermissionRequest true "Permission Details"
// @Success 200 {object} usertypes.CreateUserResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user-service/permission/{permissionId} [put]
func (s *UserService) UpsertPermission(
	w http.ResponseWriter,
	r *http.Request,
) {
	// @todo add proper permission here
	_, err := s.isAuthorized(r, usertypes.PermissionPermissionCreate.Id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := usertypes.UpserPermissionRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)

	_, err = s.upsertPermission(vars["permissionId"], req.Permission.Name, req.Permission.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.CreateUserResponse{})
	w.Write(bs)
}
