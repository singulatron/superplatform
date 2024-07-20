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
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

// Find retrieves objects based on provided criteria
// @Summary Retrieve generic objects based on criteria
// @Description Retrieves objects from a specified table based on search criteria. Requires authorization and user authentication.
// @Tags generic
// @Accept json
// @Produce json
// @Param body body generictypes.FindRequest true "Find request payload"
// @Success 200 {object} generictypes.FindResponse "Successful retrieval of objects"
// @Failure 400 {object} generictypes.ErrorResponse "Invalid JSON"
// @Failure 401 {object} generictypes.ErrorResponse "Unauthorized"
// @Failure 500 {object} generictypes.ErrorResponse "Internal Server Error"
// @Router /generic/find [post]
func Find(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	genericService *genericservice.GenericService,
) {
	err := userService.IsAuthorized(generictypes.PermissionGenericView.Id, r)
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

	req := &generictypes.FindRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	objects, err := genericService.Find(genericservice.FindOptions{
		Table:      req.Table,
		UserId:     user.Id,
		Public:     req.Public,
		Conditions: req.Conditions,
		OrderBys:   req.OrderBys,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(generictypes.FindResponse{
		Objects: objects,
	})
	w.Write(bs)
}
