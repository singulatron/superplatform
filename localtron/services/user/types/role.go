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
package usertypes

import "time"

var RoleAdmin = &Role{
	Id:            "role.admin",
	Name:          "Admin Role",
	PermissionIds: []string{},
}

var RoleUser = &Role{
	Id:            "role.user",
	Name:          "User Role",
	PermissionIds: []string{},
}

type Role struct {
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name          string   `json:"name"`
	Description   string   `json:"description"`
	PermissionIds []string `json:"permissionIds"`
}

func (c *Role) GetId() string {
	return c.Id
}

func (c *Role) GetUpdatedAt() string {
	return c.Id
}

type CreateRoleRequest struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	PermissionIds []string `json:"permissionIds"`
}

type CreateRoleResponse struct {
	Role *Role `json:"role,omitempty"`
}

type DeleteRoleRequest struct {
	RoleId string `json:"roleId,omitempty"`
}

type DeleteRoleResponse struct{}

type RemoveRoleRequest struct {
	UserId string `json:"userId,omitempty"`
	RoleId string `json:"roleId,omitempty"`
}

type RemoveRoleResponse struct{}

type GetRolesRequest struct {
}

type GetRolesResponse struct {
	Roles []*Role `json:"roles,omitempty"`
}

type SetRolePermissionsRequest struct {
	RoleId        string   `json:"roleId,omitempty"`
	PermissionIds []string `json:"permissionIds"`
}

type SetRolePermissionsResponse struct{}
