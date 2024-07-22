/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package genericendpoints

import (
	"encoding/json"
	"net/http"

	genericservice "github.com/singulatron/singulatron/localtron/services/generic"
	generictypes "github.com/singulatron/singulatron/localtron/services/generic/types"

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Create creates a new generic object
// @Summary Create a new generic object
// @Description Creates a new object with the provided details. Requires authorization and user authentication.
// @Tags generic
// @Accept json
// @Produce json
// @Param body body generictypes.CreateRequest true "Create request payload"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 400 {object} generictypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} generictypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} generictypes.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /generic/create [post]
func Create(
	w http.ResponseWriter,
	r *http.Request,
	userService usertypes.UserServiceI,
	genericService *genericservice.GenericService,
) {
	err := userService.IsAuthorized(generictypes.PermissionGenericCreate.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user, found, err := userService.GetUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !found {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	req := &generictypes.CreateRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	req.Object.UserId = user.Id

	err = genericService.Create(req.Object)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{}`))
}
