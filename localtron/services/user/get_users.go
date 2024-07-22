/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"github.com/singulatron/singulatron/localtron/datastore"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) GetUsers(options *usertypes.GetUsersOptions) ([]*usertypes.User, int64, error) {
	if len(options.Query.Conditions) == 0 {
		options.Query.Conditions = append(options.Query.Conditions, datastore.All())
	}

	additional := []datastore.Condition{}
	if len(options.Query.Conditions) > 1 {
		additional = options.Query.Conditions[1:]
	}
	q := s.usersStore.Query(
		options.Query.Conditions[0], additional...,
	).Limit(options.Query.Limit)

	if len(options.Query.OrderBys) > 1 {
		q = q.OrderBy(options.Query.OrderBys[0], options.Query.OrderBys[1:]...)
	} else if len(options.Query.OrderBys) > 0 {
		q = q.OrderBy(options.Query.OrderBys[0])
	} else {
		q = q.OrderBy(datastore.OrderByField("createdAt", true))
	}

	if options.Query.After != nil {
		q = q.After(options.Query.After...)
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
