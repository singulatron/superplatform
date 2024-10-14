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

	user "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// GetRoles handles the retrieval of all roles.
// @ID getRoles
// @Summary Get all Roles
// @Description Retrieve all roles from the user service.
// @Tags User Svc
// @Accept json
// @Produce json
// @Success 200 {object} user.GetRolesResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/roles [get]
func (s *UserService) GetRoles(
	w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &user.IsAuthorizedResponse{}
	err := s.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", user.PermissionRoleView.Id), &user.IsAuthorizedRequest{}, rsp)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	if !rsp.Authorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	roles, err := s.getRoles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.GetRolesResponse{
		Roles: roles,
	})
	w.Write(bs)
}
