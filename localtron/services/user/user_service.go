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

	"github.com/singulatron/singulatron/localtron/lib"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

type UserService struct {
	configService *configservice.ConfigService

	usersMem        *lib.MemoryStore[*usertypes.User]
	usersFile       *lib.StateManager[*usertypes.User]
	rolesMem        *lib.MemoryStore[*usertypes.Role]
	rolesFile       *lib.StateManager[*usertypes.Role]
	permissionsMem  *lib.MemoryStore[*usertypes.Permission]
	permissionsFile *lib.StateManager[*usertypes.Permission]
	authTokensMem   *lib.MemoryStore[*usertypes.AuthToken]
	authTokensFile  *lib.StateManager[*usertypes.AuthToken]
}

func NewUserService(cs *configservice.ConfigService) (*UserService, error) {
	usersPath := path.Join(cs.ConfigDirectory, "data", "users.json")
	rolesPath := path.Join(cs.ConfigDirectory, "data", "roles.json")
	permissionsPath := path.Join(cs.ConfigDirectory, "data", "permissions.json")
	authTokensPath := path.Join(cs.ConfigDirectory, "data", "authTokens.json")

	um := lib.NewMemoryStore[*usertypes.User]()
	rm := lib.NewMemoryStore[*usertypes.Role]()
	pm := lib.NewMemoryStore[*usertypes.Permission]()
	am := lib.NewMemoryStore[*usertypes.AuthToken]()

	service := &UserService{
		configService:   cs,
		usersMem:        um,
		usersFile:       lib.NewStateManager[*usertypes.User](um, usersPath),
		rolesMem:        rm,
		rolesFile:       lib.NewStateManager[*usertypes.Role](rm, rolesPath),
		permissionsMem:  pm,
		permissionsFile: lib.NewStateManager[*usertypes.Permission](pm, permissionsPath),
		authTokensMem:   am,
		authTokensFile:  lib.NewStateManager[*usertypes.AuthToken](am, authTokensPath),
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

	go service.usersFile.PeriodicSaveState(2 * time.Second)
	go service.rolesFile.PeriodicSaveState(2 * time.Second)
	go service.permissionsFile.PeriodicSaveState(2 * time.Second)
	go service.authTokensFile.PeriodicSaveState(2 * time.Second)

	return service, nil
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
