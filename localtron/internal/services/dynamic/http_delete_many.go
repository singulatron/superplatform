/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dynamicservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	generic "github.com/singulatron/singulatron/localtron/internal/services/dynamic/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// Delete removes a generic object based on the provided conditions
// @ID deleteObjects
// @Summary     Delete a Generic Object
// @Description Removes a generic object from the system based on the provided conditions. Requires authorization and user authentication.
// @Tags        Generic Svc
// @Accept      json
// @Produce     json
// @Param       objectId  path     string  true  "Object ID"
// @Param       body      body     generic.DeleteObjectRequest true "Delete request payload"
// @Success     200       {object} generic.DeleteObjectResponse "Successful deletion of object"
// @Failure     400       {object} generic.ErrorResponse "Invalid JSON"
// @Failure     401       {object} generic.ErrorResponse "Unauthorized"
// @Failure     500       {object} generic.ErrorResponse "Internal Server Error"
// @Security    BearerAuth
// @Router      /dynamic-svc/objects/delete [post]
func (g *DynamicService) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {

	rsp := &usertypes.IsAuthorizedResponse{}
	err := g.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", generic.PermissionGenericDelete.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req := &generic.DeleteObjectRequest{}
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
