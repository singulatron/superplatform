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

type User struct {
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
	Name         string     `json:"name,omitempty"`
	Email        string     `json:"email,omitempty"`
	PasswordHash string     `json:"passwordHash,omitempty"`

	RoleIds      []string `json:"roleIds,omitempty"`
	AuthTokenIds []string `json:"authTokenIds,omitempty"`
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
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
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

type GetUsersRequest struct{}

type GetUsersResponse struct {
	Users []*User `json:"users,omitempty"`
}
