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

	user "github.com/singulatron/singulatron/localtron/services/user/types"
)

// @ID register
// @Summary Register
// @Description Register a new user with a name, email, and password.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.RegisterRequest true "Register Request"
// @Success 200 {object} user.RegisterResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Router /user-svc/register [post]
func (s *UserService) Register(w http.ResponseWriter, r *http.Request) {
	req := user.RegisterRequest{}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = s.createUser(&user.User{
		Name:     req.Name,
		Slug:     req.Slug,
		Contacts: []user.Contact{req.Contact},
	}, req.Password, []string{user.RoleUser.Id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.RegisterResponse{})
	w.Write(bs)
}
