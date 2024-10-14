/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"encoding/json"

	"github.com/singulatron/superplatform/sdk/go/datastore"
	usertypes "github.com/singulatron/superplatform/server/internal/services/user/types"
)

func (s *UserService) getUsers(options *usertypes.GetUsersOptions) ([]*usertypes.User, int64, error) {
	q := s.usersStore.Query(
		options.Query.Filters...,
	).Limit(options.Query.Limit)

	if len(options.Query.OrderBys) > 0 {
		q = q.OrderBy(options.Query.OrderBys...)
	} else {
		q = q.OrderBy(datastore.OrderByField("createdAt", true))
	}

	if options.Query.JSONAfter != "" {
		v := []any{}
		err := json.Unmarshal([]byte(options.Query.JSONAfter), &v)
		if err != nil {
			return nil, 0, err
		}
		q = q.After(v...)
	}

	res, err := q.Find()
	if err != nil {
		return nil, 0, err
	}

	var count int64
	if options.Query.Count {
		var err error
		count, err = q.Count()
		if err != nil {
			return nil, 0, err
		}
	}

	users := []*usertypes.User{}
	for _, v := range res {
		users = append(users, v.(*usertypes.User))
	}

	return users, count, nil
}
