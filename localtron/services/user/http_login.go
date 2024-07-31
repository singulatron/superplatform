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

func (s *UserService) Login(w http.ResponseWriter, r *http.Request) {
	req := usertypes.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	token, err := s.login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(usertypes.LoginResponse{
		Token: token,
	})
	w.Write(bs)
}
