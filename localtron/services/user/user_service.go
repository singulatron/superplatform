/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package userservice

import (
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/logger"

	configservice "github.com/singulatron/singulatron/localtron/services/config"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

type UserService struct {
	configService *configservice.ConfigService

	usersStore       datastore.DataStore[*usertypes.User]
	rolesStore       datastore.DataStore[*usertypes.Role]
	permissionsStore datastore.DataStore[*usertypes.Permission]
	authTokensStore  datastore.DataStore[*usertypes.AuthToken]
}

func NewUserService(
	cs *configservice.ConfigService,
	usersStore datastore.DataStore[*usertypes.User],
	rolesStore datastore.DataStore[*usertypes.Role],
	authTokensStore datastore.DataStore[*usertypes.AuthToken],
	permissionsStore datastore.DataStore[*usertypes.Permission]) (*UserService, error) {

	service := &UserService{
		configService: cs,

		usersStore:       usersStore,
		rolesStore:       rolesStore,
		authTokensStore:  authTokensStore,
		permissionsStore: permissionsStore,
	}

	err := service.registerRoles()
	if err != nil {
		return nil, err
	}

	err = service.registerPermissions()
	if err != nil {
		return nil, err
	}

	err = service.bootstrap()
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (s *UserService) bootstrap() error {
	count, err := s.usersStore.Query(
		datastore.All(),
	).Count()

	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	logger.Info("Bootstrapping users")

	_, err = s.Register("singulatron", "changeme", "Admin", []string{
		usertypes.RoleAdmin.Id,
	})
	return err
}

func (s *UserService) registerRoles() error {
	_, err := s.UpsertRole(
		usertypes.RoleAdmin.Id,
		usertypes.RoleAdmin.Name,
		"",
		usertypes.RoleAdmin.PermissionIds,
	)
	if err != nil {
		return err
	}

	_, err = s.UpsertRole(
		usertypes.RoleUser.Id,
		usertypes.RoleUser.Name,
		"",
		usertypes.RoleUser.PermissionIds,
	)
	if err != nil {
		return err
	}

	return nil
}
