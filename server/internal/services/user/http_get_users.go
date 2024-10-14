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

	"github.com/singulatron/superplatform/sdk/go/datastore"
	user "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// GetUsers retrieves a list of users based on query parameters
// @ID getUsers
// @Summary List Users
// @Description Fetches a list of users with optional query filters and pagination.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param request body user.GetUsersRequest false "Get Users Request"
// @Success 200 {object} user.GetUsersResponse "List of users retrieved successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/users [post]
func (s *UserService) GetUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	_, err := s.isAuthorized(r, user.PermissionUserView.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := user.GetUsersRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	options := &user.GetUsersOptions{
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for i := range users {
		users[i].PasswordHash = ""
	}

	bs, _ := json.Marshal(user.GetUsersResponse{
		Users: users,
		Count: count,
	})
	w.Write(bs)
}
