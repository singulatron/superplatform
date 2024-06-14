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
	"path"
	"time"

	"github.com/singulatron/singulatron/localtron/logger"
	"github.com/singulatron/singulatron/localtron/memorystore"
	"github.com/singulatron/singulatron/localtron/statemanager"

	configservice "github.com/singulatron/singulatron/localtron/services/config"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

type UserService struct {
	configService *configservice.ConfigService

	usersMem        *memorystore.MemoryStore[*usertypes.User]
	usersFile       *statemanager.StateManager[*usertypes.User]
	rolesMem        *memorystore.MemoryStore[*usertypes.Role]
	rolesFile       *statemanager.StateManager[*usertypes.Role]
	permissionsMem  *memorystore.MemoryStore[*usertypes.Permission]
	permissionsFile *statemanager.StateManager[*usertypes.Permission]
	authTokensMem   *memorystore.MemoryStore[*usertypes.AuthToken]
	authTokensFile  *statemanager.StateManager[*usertypes.AuthToken]
}

func NewUserService(cs *configservice.ConfigService) (*UserService, error) {
	usersPath := path.Join(cs.ConfigDirectory, "data", "users")
	rolesPath := path.Join(cs.ConfigDirectory, "data", "roles")
	permissionsPath := path.Join(cs.ConfigDirectory, "data", "permissions")
	authTokensPath := path.Join(cs.ConfigDirectory, "data", "authTokens")

	um := memorystore.New[*usertypes.User]()
	rm := memorystore.New[*usertypes.Role]()
	pm := memorystore.New[*usertypes.Permission]()
	am := memorystore.New[*usertypes.AuthToken]()

	service := &UserService{
		configService:   cs,
		usersMem:        um,
		usersFile:       statemanager.New[*usertypes.User](um, usersPath),
		rolesMem:        rm,
		rolesFile:       statemanager.New[*usertypes.Role](rm, rolesPath),
		permissionsMem:  pm,
		permissionsFile: statemanager.New[*usertypes.Permission](pm, permissionsPath),
		authTokensMem:   am,
		authTokensFile:  statemanager.New[*usertypes.AuthToken](am, authTokensPath),
	}

	err := service.usersFile.LoadState()
	if err != nil {
		return nil, err
	}

	err = service.rolesFile.LoadState()
	if err != nil {
		return nil, err
	}

	err = service.permissionsFile.LoadState()
	if err != nil {
		return nil, err
	}

	err = service.authTokensFile.LoadState()
	if err != nil {
		return nil, err
	}

	err = service.registerRoles()
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

	go service.usersFile.PeriodicSaveState(2 * time.Second)
	go service.rolesFile.PeriodicSaveState(2 * time.Second)
	go service.permissionsFile.PeriodicSaveState(2 * time.Second)
	go service.authTokensFile.PeriodicSaveState(2 * time.Second)

	return service, nil
}

func (s *UserService) bootstrap() error {
	if s.usersMem.Count() == 0 {
		logger.Info("Bootstrapping users")
		_, err := s.Register("singulatron", "changeme", "Admin", []*usertypes.Role{
			usertypes.RoleAdmin,
		})
		return err
	}
	return nil
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
