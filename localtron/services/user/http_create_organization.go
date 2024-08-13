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

	user "github.com/singulatron/singulatron/localtron/services/user/types"
)

// CreateOrganization allows a user to create a new organization
// @ID createOrganization
// @Summary Create an Organization
// @Description Allows a logged-in user to create a new organization. The user initiating the request will be assigned the role of admin for that organization.
// @Description The initiating user will receive a dynamic role in the format `user-svc:org:$organization-slug:admin`, where `$organization-slug` is a unique identifier for the created organization.
// @Description Dynamic roles are generated based on specific user-resource associations, offering more flexible permission management compared to static roles.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param request body user.CreateOrganizationRequest true "Create User Request"
// @Success 200 {object} user.CreateOrganizationResponse "User created successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organization [post]
func (s *UserService) CreateOrganization(
	w http.ResponseWriter,
	r *http.Request) {
	usr, err := s.isAuthorized(r, user.PermissionUserCreate.Id, nil, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req := user.CreateOrganizationRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = s.createOrganization(usr.Id, req.Name, req.Slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, _ := json.Marshal(user.CreateOrganizationResponse{})
	w.Write(bs)
}
