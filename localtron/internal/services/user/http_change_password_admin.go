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

// ChangePasswordAdmin updates a user's password by an administrator
// @ID changePasswordAdmin
// @Summary Change User Password (Admin)
// @Description Allows an administrator to change a user's password.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param request body user.ChangePasswordAdminRequest true "Change Password Request"
// @Success 200 {object} user.ChangePasswordAdminResponse "Password changed successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/change-password-admin [post]
func (s *UserService) ChangePasswordAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := s.isAuthorized(r, user.PermissionUserPasswordChange.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := user.ChangePasswordAdminRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = s.changePasswordAdmin(req.Slug, req.NewPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.ChangePasswordAdminResponse{})
	w.Write(bs)
}
