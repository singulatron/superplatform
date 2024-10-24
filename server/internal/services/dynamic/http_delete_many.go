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

	dynamic "github.com/singulatron/superplatform/server/internal/services/dynamic/types"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

// Delete removes a dynamic object based on the provided conditions
// @ID deleteObjects
// @Summary     Delete a Generic Object
// @Description Removes a dynamic object from the system based on the provided conditions. Requires authorization and user authentication.
// @Tags        Dynamic Svc
// @Accept      json
// @Produce     json
// @Param       body      body     dynamic.DeleteObjectRequest true "Delete request payload"
// @Success     200       {object} dynamic.DeleteObjectResponse "Successful deletion of object"
// @Failure     400       {object} dynamic.ErrorResponse "Invalid JSON"
// @Failure     401       {object} dynamic.ErrorResponse "Unauthorized"
// @Failure     500       {object} dynamic.ErrorResponse "Internal Server Error"
// @Security    BearerAuth
// @Router      /dynamic-svc/objects/delete [post]
func (g *DynamicService) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {

	rsp := &usertypes.IsAuthorizedResponse{}
	err := g.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", dynamic.PermissionGenericDelete.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	if !rsp.Authorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	req := &dynamic.DeleteObjectRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = g.delete(req.Table, rsp.User.Id, req.Filters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{}`))
}
