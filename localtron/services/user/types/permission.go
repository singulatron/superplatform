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

type Permission struct {
	// eg. "user.viewer"
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	// eg. "User Viewer"
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *Permission) GetId() string {
	return c.Id
}

func (c *Permission) GetUpdatedAt() string {
	return c.Id
}

type CreatePermissionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreatePermissionResponse struct {
}

var PermissionUserCreate = Permission{
	Id:   "user.create",
	Name: "User Create",
}

var PermissionUserView = Permission{
	Id:   "user.view",
	Name: "User View",
}

var PermissionUserEdit = Permission{
	Id:   "user.edit",
	Name: "User Edit",
}

var PermissionUserDelete = Permission{
	Id:   "user.delete",
	Name: "User Delete",
}

var PermissionUserStream = Permission{
	Id:   "user.stream",
	Name: "User Stream",
}

var UserPermissions = []Permission{
	PermissionUserCreate,
	PermissionUserView,
	PermissionUserEdit,
	PermissionUserDelete,
	PermissionUserStream,
}
