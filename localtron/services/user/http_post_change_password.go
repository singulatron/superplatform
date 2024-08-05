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

// ChangePassword allows a user to update their own password
// @Summary Change User Password
// @Description Allows an authenticated user to change their own password.
// @Tags User Service
// @Accept json
// @Produce json
// @Param request body usertypes.ChangePasswordRequest true "Change Password Request"
// @Success 200 {object} usertypes.ChangePasswordResponse "Password changed successfully"
// @Failure 400 {object} usertypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} usertypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} usertypes.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-service/change-password [post]
func (s *UserService) ChangePassword(w http.ResponseWriter, r *http.Request) {
	_, err := s.isAuthorized(r, usertypes.PermissionUserPasswordChange.Id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := usertypes.ChangePasswordRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = s.changePassword(req.Email, req.CurrentPassword, req.NewPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.ChangePasswordResponse{})
	w.Write(bs)
}
