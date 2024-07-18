/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package modelservice

import (
	"github.com/singulatron/singulatron/localtron/datastore"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (p *ModelService) registerPermissions() error {
	for _, permission := range modeltypes.ModelPermissions {
		_, err := p.userService.UpsertPermission(
			permission.Id,
			permission.Name,
			permission.Description,
		)
		if err != nil {
			return err
		}
	}

	for _, role := range []*usertypes.Role{
		usertypes.RoleAdmin,
		usertypes.RoleUser,
	} {
		for _, permission := range modeltypes.ModelPermissions {
			p.userService.AddPermissionToRole(role.Id, permission.Id)
		}
	}

	return nil
}

func (p *ModelService) bootstrapModels() error {
	platformRows := []datastore.Row{}
	for _, v := range modeltypes.Platforms {
		platformRows = append(platformRows, v)
	}
	err := p.platformsStore.UpsertMany(platformRows)
	if err != nil {
		return err
	}

	modelRows := []datastore.Row{}
	for _, v := range modeltypes.Models {
		modelRows = append(modelRows, v)
	}

	return p.modelsStore.UpsertMany(modelRows)

}
