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
	"fmt"
	"net/http"

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// GetRoles handles the retrieval of all roles.
// @Summary Get all Roles
// @Description Retrieve all roles from the user service.
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} usertypes.GetRolesResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-service/roles [get]
func (s *UserService) GetRoles(
	w http.ResponseWriter,
	r *http.Request) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := s.router.AsRequestMaker(r).Post(r.Context(), "user-service", fmt.Sprintf("/permission/%v/is-authorized", usertypes.PermissionRoleView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	roles, err := s.getRoles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.GetRolesResponse{
		Roles: roles,
	})
	w.Write(bs)
}
