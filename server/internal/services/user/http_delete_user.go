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
	user "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// DeleteUser handles the deletion of a user by user ID.
// @ID deleteUser
// @Summary Delete a User
// @Description Delete a user based on the user ID.
// @Tags User Svc
// @Accept  json
// @Produce  json
// @Param   userId     path    string  true  "User ID"
// @Success 200 {object} user.DeleteUserResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/user/{userId} [delete]
func (s *UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	usr, err := s.isAuthorized(r, user.PermissionUserDelete.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	callerUserId := usr.Id
	isAdmin, err := s.isAdmin(callerUserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAdmin {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	userId := mux.Vars(r)["userId"]

	err = s.deleteUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.DeleteUserResponse{})
	w.Write(bs)
}
