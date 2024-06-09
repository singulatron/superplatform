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

	userservice "github.com/singulatron/singulatron/localtron/services/user"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func Login(w http.ResponseWriter, r *http.Request, userService *userservice.UserService) {
	req := usertypes.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	token, err := userService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.LoginResponse{
		Token: token,
	})
	w.Write(bs)
}
