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
package userservice

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/singulatron/singulatron/localtron/logger"

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) IsAuthorized(permissionId string, request *http.Request) error {
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return errors.New("unauthorized")
	}
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	// @todo this is very inefficient
	var token *usertypes.AuthToken
	found := s.authTokensMem.ForeachStop(func(i int, tk *usertypes.AuthToken) bool {
		if tk.Token == authHeader {
			token = tk
			return true
		}
		return false
	})

	if !found {
		return errors.New("unauthorized")
	}

	user, found := s.usersMem.FindById(token.UserId)
	if !found {
		logger.Error("Token refers to nonexistent user",
			slog.String("userId", token.UserId),
			slog.String("tokenId", token.Id),
		)
		return errors.New("unauthorized")
	}

	for _, roleId := range user.RoleIds {
		role, found := s.rolesMem.FindById(roleId)
		if !found {
			logger.Error("User refers to a nonexistent role",
				slog.String("userId", token.UserId),
				slog.String("roleId", roleId),
			)
			return errors.New("unauthorized")
		}
		for _, permId := range role.PermissionIds {
			if permId == permissionId {
				return nil
			}
		}
	}

	return errors.New("unauthorized")
}

func (s *UserService) GetUserFromRequest(request *http.Request) (*usertypes.User, bool) {
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return nil, false
	}
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	// @todo this is very inefficient
	var token *usertypes.AuthToken
	found := s.authTokensMem.ForeachStop(func(i int, tk *usertypes.AuthToken) bool {
		if tk.Token == authHeader {
			token = tk
			return true
		}
		return false
	})

	if !found {
		return nil, false
	}

	user, found := s.usersMem.FindById(token.UserId)
	if !found {
		logger.Error("Token refers to nonexistent user",
			slog.String("userId", token.UserId),
			slog.String("tokenId", token.Id),
		)
		return nil, false
	}

	return user, true
}
