/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package usertypes

import (
	"time"

	"github.com/singulatron/singulatron/localtron/datastore"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type GetUsersOptions struct {
	Query *datastore.Query `json:"query"`
}

type User struct {
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Name      string     `json:"name,omitempty"`

	// Email or username
	Email string `json:"email,omitempty"`

	PasswordHash string `json:"passwordHash,omitempty"`

	/* Many to many relationship between User and Role */
	RoleIds []string `json:"roleIds,omitempty"`

	IsService bool `json:"isService,omitempty"`
}

type KeyPair struct {
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	PrivateKey string `json:"privateKey,omitempty"`
	PublicKey  string `json:"publicKey,omitempty"`
}

func (k *KeyPair) GetId() string {
	return k.Id
}

func (c *User) GetId() string {
	return c.Id
}

func (c *User) GetUpdatedAt() string {
	return c.Id
}

type ReadUserByTokenRequest struct {
	Token string `json:"token,omitempty"`
}

type ReadUserByTokenResponse struct {
	User *User `json:"user,omitempty"`
}

type RegisterRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type RegisterResponse struct {
}

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token *AuthToken `json:"token,omitempty"`
}

type SaveProfileRequest struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type SaveProfileResponse struct {
}

type ChangePasswordRequest struct {
	Email           string `json:"email,omitempty"`
	CurrentPassword string `json:"currentPassword,omitempty"`
	NewPassword     string `json:"newPassword,omitempty"`
}

type ChangePasswordResponse struct{}

type ChangePasswordAdminRequest struct {
	Email       string `json:"email,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}

type ChangePasswordAdminResponse struct{}

type GetUsersRequest struct {
	Query *datastore.Query `json:"query"`
}

type GetUsersResponse struct {
	Users []*User   `json:"users,omitempty"`
	After time.Time `json:"after,omitempty"`
	Count int64     `json:"count"`
}

type CreateUserRequest struct {
	User     *User    `json:"user,omitempty"`
	Password string   `json:"password,omitempty"`
	RoleIds  []string `json:"roleIds,omitempty"`
}

type CreateUserResponse struct {
}

type DeleteUserRequest struct {
	UserId string `json:"userId,omitempty"`
}

type DeleteUserResponse struct{}

type Credential struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (c *Credential) GetId() string {
	return c.Email
}

type GetPublicKeyRequest struct{}
type GetPublicKeyResponse struct {
	PublicKey string `json:"publicKey,omitempty"`
}
