/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package usertypes

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/logger"
	"github.com/singulatron/singulatron/localtron/router"
)

type User struct {
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Name      string     `json:"name,omitempty"`

	// Email or username
	Email string `json:"email,omitempty"`

	PasswordHash string `json:"passwordHash,omitempty"`

	RoleIds      []string `json:"roleIds,omitempty"`
	AuthTokenIds []string `json:"authTokenIds,omitempty"`

	IsService bool `json:"isService,omitempty"`
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

// RegisterService registers a service or logs in with credentials loaded
// from the credentials store.
// Every service should have a user account in the user service and this method creates
// that user account.
func RegisterService(serviceEmail, serviceName string, router *router.Router, store datastore.DataStore) (string, error) {
	res, err := store.Query(datastore.All()).Find()
	if err != nil {
		return "", err
	}

	email := serviceName
	pw := ""

	if len(res) > 0 {
		cred := res[0].(*Credential)
		email = cred.Email
		pw = cred.Password
	} else {
		logger.Info(fmt.Sprintf("Registering the %v service", serviceEmail))

		pw = uuid.New().String()
		err = router.Post(context.Background(), "user", "/register", RegisterRequest{
			Email:    email,
			Name:     serviceName,
			Password: pw,
		}, nil)
		if err != nil {
			return "", err
		}
		err = store.Upsert(&Credential{
			Email:    email,
			Password: pw,
		})
		if err != nil {
			return "", err
		}
	}

	rsp := LoginResponse{}
	err = router.Post(context.Background(), "user", "/login", LoginRequest{
		Email:    email,
		Password: pw,
	}, &rsp)
	if err != nil {
		return "", err
	}

	return rsp.Token.Token, nil
}

func RegisterUser(router *router.Router, email, password, username string) (string, error) {
	err := router.Post(context.Background(), "user", "/register", &RegisterRequest{
		Email:    email,
		Password: password,
		Name:     username,
	}, nil)
	if err != nil {
		return "", err
	}

	loginRsp := LoginResponse{}
	err = router.Post(context.Background(), "user", "/login", &LoginRequest{
		Email:    email,
		Password: password,
	}, &loginRsp)
	if err != nil {
		return "", err
	}

	return loginRsp.Token.Token, nil
}
