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

// @Summary Ge Public Key
// @Description Get the public key to descrypt the JWT.
// @Tags User Service
// @Accept json
// @Produce json
// @Success 200 {object} usertypes.GetPublicKeyResponse
// @Failure 400 {object} usertypes.ErrorResponse "Invalid JSON or missing permission id"
// @Failure 401 {object} usertypes.ErrorResponse "Unauthorized"
// @Router /user-service/public-key [get]
func (s *UserService) GetPublicKey(
	w http.ResponseWriter,
	r *http.Request) {

	bs, _ := json.Marshal(usertypes.GetPublicKeyResponse{
		PublicKey: s.publicKey.N.String(),
	})
	w.Write(bs)
}
