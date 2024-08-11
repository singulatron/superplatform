/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

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

	// Full name of the organization
	Name string `json:"name,omitempty"`

	// URL-friendly unique (inside the Singularon platform) identifier for the `user`.
	Slug string `json:"slug,omitempty"`

	// Contacts are used for login and identification purposes.
	Contacts []Contact `json:"contact,omitempty"`

	PasswordHash string `json:"passwordHash,omitempty"`

	/* Many to many relationship between User and Role */
	RoleIds []string `json:"roleIds,omitempty"`

	/* Many to many relationship between User and Organization */
	OrganizationIds []string `json:"organizationIds,omitempty"`
}

type Contact struct {
	// The unique identifier, which can be a URL.
	//
	// Example values: "joe12" (singulatron username), "twitter.com/thejoe" (twitter url), "joe@joesdomain.com" (email)
	Id string `json:"id,omitempty" example:"twitter.com/thejoe"`

	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	UserId   string `json:"userId,omitempty"`
	Platform string `json:"platform,omitempty" example:"twitter"` // Platform of the contact (e.g., "email", "phone", "twitter")

	// Value is the platform local unique identifier.
	// Ie. while the `id` of a Twitter contact is `twitter.com/thejoe`, the value will be only `thejoe`.
	// For email and phones the `id` and the `value` will be the same.
	// This field mostly exists for display purposes.
	//
	// Example values: "joe12" (singulatron username), "thejoe" (twitter username), "joe@joesdomain.com" (email)
	Value string `json:"value,omitempty" example:"thejoe"`

	Verified bool `json:"verified,omitempty"` // Whether the contact is verified
	Primary  bool `json:"primary,omitempty"`  // If this is the primary contact method
}

type Organization struct {
	Id        string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Full name of the organization
	Name string `json:"name,omitempty"`

	// URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
	Slug string `json:"slug,omitempty"`
}

type ContactPlatform string

const (
	ContactPlatformEmail    ContactPlatform = "email"
	ContactPlatformPhone    ContactPlatform = "phone"
	ContactPlatformTwitter  ContactPlatform = "twitter"
	ContactPlatformLinkedIn ContactPlatform = "linkedin"
)

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
	Name     string  `json:"name,omitempty"`
	Slug     string  `json:"slug,omitempty"`
	Contact  Contact `json:"contact,omitempty"`
	Password string  `json:"password,omitempty"`
}

type RegisterResponse struct {
}

type LoginRequest struct {
	Slug     string `json:"slug,omitempty"`
	Contact  string `json:"contact,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token *AuthToken `json:"token,omitempty"`
}

type SaveProfileRequest struct {
	Slug string `json:"slug,omitempty"`
	Name string `json:"name,omitempty"`
}

type SaveProfileResponse struct {
}

type ChangePasswordRequest struct {
	Slug            string `json:"slug,omitempty"`
	CurrentPassword string `json:"currentPassword,omitempty"`
	NewPassword     string `json:"newPassword,omitempty"`
}

type ChangePasswordResponse struct{}

type ChangePasswordAdminRequest struct {
	Slug        string `json:"slug,omitempty"`
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
	Slug     string `json:"slug,omitempty"`
	Contact  string `json:"contact,omitempty"`
	Password string `json:"password,omitempty"`
}

func (c *Credential) GetId() string {
	return c.Contact
}

type GetPublicKeyRequest struct{}
type GetPublicKeyResponse struct {
	PublicKey string `json:"publicKey,omitempty"`
}
