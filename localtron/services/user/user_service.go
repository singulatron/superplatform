/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"crypto/rsa"
	"time"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/logger"
	"github.com/singulatron/singulatron/localtron/router"

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

type UserService struct {
	router *router.Router

	usersStore       datastore.DataStore
	rolesStore       datastore.DataStore
	permissionsStore datastore.DataStore
	authTokensStore  datastore.DataStore
	keyPairsStore    datastore.DataStore

	privateKey   *rsa.PrivateKey
	publicKeyPem string
}

func NewUserService(
	router *router.Router,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*UserService, error) {
	usersStore, err := datastoreFactory("users", &usertypes.User{})
	if err != nil {
		return nil, err
	}
	rolesStore, err := datastoreFactory("roles", &usertypes.Role{})
	if err != nil {
		return nil, err
	}
	authTokensStore, err := datastoreFactory("authTokens", &usertypes.AuthToken{})
	if err != nil {
		return nil, err
	}
	permissionsStore, err := datastoreFactory("permissions", &usertypes.Permission{})
	if err != nil {
		return nil, err
	}
	keyPairsStore, err := datastoreFactory("keyPairs", &usertypes.KeyPair{})
	if err != nil {
		return nil, err
	}

	service := &UserService{
		router:           router,
		usersStore:       usersStore,
		rolesStore:       rolesStore,
		authTokensStore:  authTokensStore,
		permissionsStore: permissionsStore,
		keyPairsStore:    keyPairsStore,
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

	return service, nil
}

func (s *UserService) bootstrap() error {
	keyPairs, err := s.keyPairsStore.Query(
		datastore.All(),
	).Find()
	if err != nil {
		return err
	}

	if len(keyPairs) > 0 {
		kp := keyPairs[0].(*usertypes.KeyPair)
		privKey, err := privateKeyFromString(kp.PrivateKey)
		if err != nil {
			return err
		}
		s.privateKey = privKey
		s.publicKeyPem = kp.PublicKey
	} else {
		privKey, pubKey, err := generateRSAKeys(4096)
		if err != nil {
			return err
		}
		now := time.Now()
		kp := &usertypes.KeyPair{
			Id:         uuid.New().String(),
			CreatedAt:  now,
			UpdatedAt:  now,
			PublicKey:  pubKey,
			PrivateKey: privKey,
		}
		err = s.keyPairsStore.Upsert(kp)
		if err != nil {
			return err
		}

		privKeyTyped, err := privateKeyFromString(kp.PrivateKey)
		if err != nil {
			return err
		}
		s.privateKey = privKeyTyped
		s.publicKeyPem = kp.PublicKey

	}

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

	_, err = s.register("singulatron", "changeme", "Admin", []string{
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
