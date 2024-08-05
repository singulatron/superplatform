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
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/logger"

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// @Summary Is Authorized
// @Description Check if a user is authorized for a specific permission.
// @Tags User Service
// @Accept json
// @Produce json
// @Param permissionId path string true "Permission ID"
// @Param body body usertypes.IsAuthorizedRequest true "Is Authorized Request"
// @Success 200 {object} usertypes.IsAuthorizedResponse
// @Failure 400 {object} usertypes.ErrorResponse "Invalid JSON or missing permission id"
// @Failure 401 {object} usertypes.ErrorResponse "Unauthorized"
// @Security BearerAuth
// @Router /user-service/permission/{permissionId}/is-authorized [post]
func (s *UserService) IsAuthorized(
	w http.ResponseWriter,
	r *http.Request,
) {
	req := &usertypes.IsAuthorizedRequest{}
	//m := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `Invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	permissionId := vars["permissionId"]

	if permissionId == "" {
		http.Error(w, `missing permission id`, http.StatusBadRequest)
		return
	}

	user, err := s.isAuthorized(r, permissionId, req.EmailsGranted)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	bs, _ := json.Marshal(&usertypes.IsAuthorizedResponse{
		Authorized: true,
		User:       user,
	})

	w.Write(bs)
}

func (s *UserService) isAuthorized(r *http.Request, permissionId string,
	emailsGranted []string) (*usertypes.User, error) {
	user, err := s.getUserFromRequest(r)
	if err != nil {
		return nil, err
	}

	emailGrant := false
	for _, v := range emailsGranted {
		if user.Email == v {
			emailGrant = true
		}
	}
	if emailGrant {
		return user, nil
	}

	roles, err := s.rolesStore.Query(
		datastore.Equal(datastore.Field("id"), user.RoleIds),
	).Find()
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		for _, permId := range role.(*usertypes.Role).PermissionIds {
			if permId == permissionId {
				return user, nil
			}
		}
	}

	return nil, errors.New("unauthorized")
}

func (s *UserService) getUserFromRequest(r *http.Request) (*usertypes.User, error) {
	authHeader := r.Header.Get("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	if authHeader == "" {
		return nil, fmt.Errorf("Unauthorized")
	}

	tokenI, found, err := s.authTokensStore.Query(
		datastore.Equal(datastore.Field("token"), authHeader),
	).FindOne()
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("unauthorized")
	}
	token := tokenI.(*usertypes.AuthToken)

	userI, found, err := s.usersStore.Query(
		datastore.Id(token.UserId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		logger.Error("Token refers to nonexistent user",
			slog.String("userId", token.UserId),
			slog.String("tokenId", token.Id),
		)
		return nil, errors.New("unauthorized")
	}
	user := userI.(*usertypes.User)
	return user, nil
}

func (s *UserService) GetUserFromRequest(request *http.Request) (*usertypes.User, bool, error) {
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return nil, false, nil
	}
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	tokenI, found, err := s.authTokensStore.Query(
		datastore.Equal(datastore.Field("token"), authHeader),
	).FindOne()
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, errors.New("unauthorized")
	}
	token := tokenI.(*usertypes.AuthToken)

	userI, found, err := s.usersStore.Query(
		datastore.Id(token.UserId),
	).FindOne()
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	user := userI.(*usertypes.User)
	return user, found, nil
}
