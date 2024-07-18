/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userendpoints

import (
	"encoding/json"
	"net/http"

	userservice "github.com/singulatron/singulatron/localtron/services/user"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func SetRolePermissions(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService) {
	err := userService.IsAuthorized(usertypes.PermissionRoleEdit.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := usertypes.SetRolePermissionsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = userService.SetRolePermissions(req.RoleId, req.PermissionIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.SetRolePermissionsResponse{})
	w.Write(bs)
}
