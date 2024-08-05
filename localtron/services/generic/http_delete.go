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

	generictypes "github.com/singulatron/singulatron/localtron/services/generic/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Delete removes a generic object based on the provided conditions
// @Summary     Delete a Generic Object
// @Description Removes a generic object from the system based on the provided conditions. Requires authorization and user authentication.
// @Tags        Generic Service
// @Accept      json
// @Produce     json
// @Param       objectId  path     string  true  "Object ID"
// @Param       body      body     generictypes.DeleteRequest true "Delete request payload"
// @Success     200       {object} generictypes.DeleteResponse "Successful deletion of object"
// @Failure     400       {object} generictypes.ErrorResponse "Invalid JSON"
// @Failure     401       {object} generictypes.ErrorResponse "Unauthorized"
// @Failure     500       {object} generictypes.ErrorResponse "Internal Server Error"
// @Security    BearerAuth
// @Router      /generic-svc/object/{objectId} [delete]
func (g *GenericService) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {

	rsp := &usertypes.IsAuthorizedResponse{}
	err := g.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", generictypes.PermissionGenericDelete.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := &generictypes.DeleteRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = g.delete(req.Table, rsp.User.Id, req.Conditions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{}`))
}
