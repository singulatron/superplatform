/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"errors"
	"fmt"
	"time"

	"github.com/singulatron/singulatron/localtron/datastore"

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) createOrganization(userId, name, slug string) error {
	_, exists, err := s.contactsStore.Query(
		datastore.Equal(datastore.Field("slug"), slug),
	).FindOne()
	if err != nil {
		return err
	}

	if exists {
		return errors.New("organization already exists")
	}

	org := &usertypes.Organization{
		Name: name,
		Slug: slug,
	}

	err = s.organizationsStore.Create(org)
	if err != nil {
		return err
	}

	return s.addDynamicRoleToUser(userId, fmt.Sprintf("user-svc:org:%v:admin", slug))
}

func (s *UserService) addStaticRoleToUser(userId, roleId string) error {
	roleQ := s.rolesStore.Query(
		datastore.Id(roleId),
	)
	roleI, found, err := roleQ.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("cannot find role %v", roleId)
	}
	role := roleI.(*usertypes.Role)

	userQ := s.usersStore.Query(
		datastore.Id(userId),
	)
	userI, found, err := userQ.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("cannot find user %v", userId)
	}
	user := userI.(*usertypes.User)

	return s.userRoleLinksStore.Upsert(&usertypes.UserRoleLink{
		Id:        fmt.Sprintf("%v:%v", user.Id, role.Id),
		CreatedAt: time.Now(),

		RoleId: roleId,
		UserId: userId,
	})
}

func (s *UserService) addDynamicRoleToUser(userId, roleId string) error {
	userQ := s.usersStore.Query(
		datastore.Id(userId),
	)
	userI, found, err := userQ.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("cannot find user %v", userId)
	}
	user := userI.(*usertypes.User)

	return s.userRoleLinksStore.Upsert(&usertypes.UserRoleLink{
		Id:        fmt.Sprintf("%v:%v", user.Id, roleId),
		CreatedAt: time.Now(),

		RoleId: roleId,
		UserId: userId,
	})
}