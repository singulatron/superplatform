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

	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// GetUsers retrieves a list of users based on query parameters
// @Summary List Users
// @Description Fetches a list of users with optional query filters and pagination.
// @Tags User Service
// @Accept json
// @Produce json
// @Param request body usertypes.GetUsersRequest false "Get Users Request"
// @Success 200 {object} usertypes.GetUsersResponse "List of users retrieved successfully"
// @Failure 400 {object} usertypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} usertypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} usertypes.ErrorResponse "Internal Server Error"
// @Router /user-service/users [post]
func (s *UserService) GetUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	_, err := s.isAuthorized(r, usertypes.PermissionUserView.Id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := usertypes.GetUsersRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	options := &usertypes.GetUsersOptions{
		Query: req.Query,
	}
	if options.Query == nil {
		options.Query = &datastore.Query{}
	}
	if options.Query.Limit == 0 {
		options.Query.Limit = 20
	}

	users, count, err := s.getUsers(options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := range users {
		users[i].PasswordHash = ""
	}

	bs, _ := json.Marshal(usertypes.GetUsersResponse{
		Users: users,
		Count: count,
	})
	w.Write(bs)
}
