/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package genericservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	generic "github.com/singulatron/singulatron/localtron/internal/services/generic/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// Update modifies existing generic objects based on given conditions
// @ID updateObjects
// @Summary Update Generic Objects
// @Description Updates objects in a specified table based on provided conditions. Requires authorization and user authentication.
// @Tags Generic Svc
// @Accept json
// @Produce json
// @Param body body generic.UpdateObjectRequest true "Update request payload"
// @Success 200 {object} generic.UpdateObjectResponse "Successful update of objects"
// @Failure 400 {object} generic.ErrorResponse "Invalid JSON"
// @Failure 401 {object} generic.ErrorResponse "Unauthorized"
// @Failure 500 {object} generic.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /generic-svc/objects/update [post]
func (g *GenericService) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := g.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", generic.PermissionGenericEdit.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := &generic.UpdateObjectRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = g.update(req.Table, rsp.User.Id, req.Conditions, req.Object)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{}`))
}
