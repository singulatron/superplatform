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

func (s *UserService) SetRolePermissions(
	w http.ResponseWriter,
	r *http.Request) {
	_, err := s.isAuthorized(r, usertypes.PermissionRoleEdit.Id, nil)
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

	err = s.setRolePermissions(req.RoleId, req.PermissionIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.SetRolePermissionsResponse{})
	w.Write(bs)
}