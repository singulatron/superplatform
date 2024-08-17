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
	"strings"

	"github.com/gorilla/mux"
	user "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

// UpsertPermission handles the creation or update of a permission
// @ID upsertPermission
// @Summary Upsert a Permission
// @Description Creates or updates a permission.
// @Description <b>The permission ID must be prefixed by the callers username (email).</b>
// @Description Eg. if the owner's email/username is `petstore-svc` the permission should look like `petstore-svc:pet:edit`.
// @Descripion The user account who creates the permission will become the owner of that permission, and only the owner will be able to edit the permission.
// @Description
// @Description Requires the `user-svc:permission:create` permission.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param permissionId path string true "Permission ID"
// @Param requestBody body user.UpserPermissionRequest true "Permission Details"
// @Success 200 {object} user.CreateUserResponse
// @Failure 400 {string} string "Bad Request: Invalid JSON or Bad Namespace"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/permission/{permissionId} [put]
func (s *UserService) UpsertPermission(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	// @todo add proper permission here
	usr, err := s.isAuthorized(r, user.PermissionPermissionCreate.Id, nil, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := user.UpserPermissionRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)

	if !strings.HasPrefix(vars["permissionId"], usr.Slug) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Bad Namespace`))
		return
	}

	_, err = s.upsertPermission(usr.Id, vars["permissionId"], req.Permission.Name, req.Permission.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.CreateUserResponse{})
	w.Write(bs)
}
