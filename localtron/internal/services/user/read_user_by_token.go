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

	"github.com/singulatron/superplatform/sdk/go/datastore"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

func (s *UserService) readUserByToken(token string) (*usertypes.User, error) {
	authTokenI, found, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("token"), token),
	).FindOne()
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("token not found")
	}
	authToken := authTokenI.(*usertypes.AuthToken)

	userI, found, err := s.usersStore.Query(
		datastore.Equals(datastore.Field("id"), authToken.UserId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	ret := &usertypes.User{
		Id:        user.Id,
		Name:      user.Name,
		Slug:      user.Slug,
		Contacts:  user.Contacts,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return ret, nil
}

func (s *UserService) getUserOrganizations(userId string) ([]*usertypes.Organization, string, error) {
	links, err := s.organizationUserLinksStore.Query(
		datastore.Equals(
			datastore.Field("userId"),
			userId,
		),
	).Find()
	if err != nil {
		return nil, "", err
	}

	organizationIds := []any{}
	activeOrganizationId := ""
	for _, linkI := range links {
		link := linkI.(*usertypes.OrganizationUserLink)
		if link.Active {
			activeOrganizationId = link.OrganizationId
		}
		organizationIds = append(organizationIds, link.OrganizationId)
	}

	orgIs, err := s.organizationsStore.Query(
		datastore.IsInList(
			datastore.Field("id"),
			organizationIds...,
		),
	).Find()
	if err != nil {
		return nil, "", err
	}

	orgs := []*usertypes.Organization{}
	for _, orgI := range orgIs {
		org := orgI.(*usertypes.Organization)
		orgs = append(orgs, org)
	}

	return orgs, activeOrganizationId, nil
}
