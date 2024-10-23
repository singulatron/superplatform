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

// Create creates a new generic object
// @ID createObject
// @Summary Create a Generic Object
// @Description Creates a new object with the provided details. Requires authorization and user authentication.
// @Tags Dynamic Svc
// @Accept json
// @Produce json
// @Param body body dynamic.CreateObjectRequest true "Create request payload"
// @Success 200 {object} dynamic.CreateObjectResponse "Success"
// @Failure 400 {object} dynamic.ErrorResponse "Invalid JSON"
// @Failure 401 {object} dynamic.ErrorResponse "Unauthorized"
// @Failure 500 {object} dynamic.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /dynamic-svc/object [post]
func (g *DynamicService) Create(
	w http.ResponseWriter,
	r *http.Request,
) {

	rsp := &usertypes.IsAuthorizedResponse{}
	err := g.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", dynamic.PermissionGenericCreate.Id), &usertypes.IsAuthorizedRequest{}, rsp)
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

	req := &dynamic.CreateObjectRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	for i, v := range req.Object.Readers {
		if v == "_self" {
			req.Object.Readers[i] = rsp.User.Id
		}
	}
	for i, v := range req.Object.Writers {
		if v == "_self" {
			req.Object.Writers[i] = rsp.User.Id
		}
	}
	for i, v := range req.Object.Deleters {
		if v == "_self" {
			req.Object.Deleters[i] = rsp.User.Id
		}
	}

	err = g.create(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{}`))
}
