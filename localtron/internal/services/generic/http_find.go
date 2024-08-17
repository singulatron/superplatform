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

// Query retrieves objects based on provided criteria
// @ID query
// @Summary Query Generic Objects
// @Description Retrieves objects from a specified table based on search criteria.
// @Description Requires authorization and user authentication.
// @Description
// @Description
// @Description Use helper functions in your respective client library such as condition constructors (`equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
// @Tags Generic Svc
// @Accept json
// @Produce json
// @Param body body generic.QueryRequest false "Query Request"
// @Success 200 {object} generic.QueryResponse "Successful retrieval of objects"
// @Failure 400 {object} generic.ErrorResponse "Invalid JSON"
// @Failure 401 {object} generic.ErrorResponse "Unauthorized"
// @Failure 500 {object} generic.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /generic-svc/objects [post]
func (g *GenericService) Query(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	rsp := &usertypes.IsAuthorizedResponse{}
	err := g.router.AsRequestMaker(r).Post(r.Context(), "user-svc", fmt.Sprintf("/permission/%v/is-authorized", generic.PermissionGenericView.Id), &usertypes.IsAuthorizedRequest{}, rsp)
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

	req := &generic.QueryRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	objects, err := g.find(generic.QueryOptions{
		Table:  req.Table,
		UserId: rsp.User.Id,
		Public: req.Public,
		Query:  req.Query,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(generic.QueryResponse{
		Objects: objects,
	})
	w.Write(bs)
}
