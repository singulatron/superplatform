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

	"github.com/gorilla/mux"
	generic "github.com/singulatron/singulatron/localtron/internal/services/dynamic/types"
	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// Upsert creates or updates a generic object based on the provided data
// @ID upsertObject
// @Summary Upsert a Generic Object
// @Description Creates a new generic object or updates an existing one based on the provided data. Requires authorization and user authentication.
// @Tags Generic Svc
// @Accept json
// @Produce json
// @Param objectId path string true  "Object ID"
// @Param body body generic.UpsertObjectRequest true "Upsert request payload"
// @Success 200 {object} generic.UpsertObjectResponse "Successful creation or update of object"
// @Failure 400 {object} generic.ErrorResponse "Invalid JSON"
// @Failure 401 {object} generic.ErrorResponse "Unauthorized"
// @Failure 500 {object} generic.ErrorResponse "Internal Server Error"
// @Security    BearerAuth
// @Router /dynamic-svc/object/{objectId} [put]
func (g *DynamicService) Upsert(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := g.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", generic.PermissionGenericCreate.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !rsp.Authorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	req := &generic.UpsertObjectRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()
	req.Object.UserId = rsp.User.Id

	objectId := mux.Vars(r)
	req.Object.Id = objectId["objectId"]

	err = g.upsert(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{}`))
}
