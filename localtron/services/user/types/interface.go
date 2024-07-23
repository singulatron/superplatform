/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package usertypes

import (
	"net/http"

	"github.com/singulatron/singulatron/localtron/datastore"
)

type GetUsersOptions struct {
	Query *datastore.Query `json:"query"`
}

type UserServiceI interface {
	UpsertRole(id, name, description string, permissionIds []string) (*Role, error)
	UpsertPermission(id, name, description string) (*Permission, error)
	SetRolePermissions(roleId string, permissionIds []string) error
	SaveProfile(email, newName string) error
	RemoveRole(userId string, roleId string) error
	Register(email, password, name string, roleIds []string) (*AuthToken, error)
	ReadUserByToken(token string) (*User, error)
	Login(email, password string) (*AuthToken, error)
	IsAuthorized(permissionId string, request *http.Request) error
	GetUsers(options *GetUsersOptions) ([]*User, int64, error)
	GetRoles() ([]*Role, error)
	GetPermissions() ([]*Permission, error)
	DeleteUser(userId string) error
	DeleteRole(roleId string) error
	DeletePermission(permissionId string) error
	CreateUser(user *User, password string, roleIds []string) error
	CreateRole(name, description string, permissionIds []string) (*Role, error)
	CreatePermission(id, name, description string) (*Permission, error)
	ChangePassword(email, currentPassword, newPassword string) error
	ChangePasswordAdmin(email, newPassword string) error
	AddRole(userId string, role *Role) error
	AddPermissionToRole(roleId, permissionId string) error
	GetUserFromRequest(request *http.Request) (*User, bool, error)
}
