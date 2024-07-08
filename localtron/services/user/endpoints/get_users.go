/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package userendpoints

import (
	"encoding/json"
	"net/http"

	"github.com/singulatron/singulatron/localtron/datastore"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func GetUsers(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
) {
	err := userService.IsAuthorized(usertypes.PermissionUserView.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := usertypes.GetUsersRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	options := &userservice.GetUsersOptions{
		Query: req.Query,
	}
	if options.Query == nil {
		options.Query = &datastore.Query{}
	}
	if options.Query.Limit == 0 {
		options.Query.Limit = 20
	}

	users, count, err := userService.GetUsers(options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := range users {
		users[i].PasswordHash = ""
		users[i].AuthTokenIds = nil
	}

	bs, _ := json.Marshal(usertypes.GetUsersResponse{
		Users: users,
		Count: count,
	})
	w.Write(bs)
}
