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

// DeleteUser handles the deletion of a user by user ID.
// @Summary Delete a User
// @Description Delete a user based on the user ID.
// @Tags users
// @Accept  json
// @Produce  json
// @Param   userId     path    string  true  "User ID"
// @Success 200 {object} usertypes.DeleteUserResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-service/user/{userId} [delete]
func (s *UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, err := s.isAuthorized(r, usertypes.PermissionUserDelete.Id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userId := mux.Vars(r)["userId"]

	err = s.deleteUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.DeleteUserResponse{})
	w.Write(bs)
}
