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

type GetUsersOptions struct {
	Query *datastore.Query `json:"query"`
}

func (s *UserService) GetUsers(options *GetUsersOptions) ([]*usertypes.User, int64, error) {
	if len(options.Query.Conditions) == 0 {
		options.Query.Conditions = append(options.Query.Conditions, datastore.All())
	}

	additional := []datastore.Condition{}
	if len(options.Query.Conditions) > 1 {
		additional = options.Query.Conditions[1:]
	}
	q := s.usersStore.Query(
		options.Query.Conditions[0], additional...,
	).Limit(int(options.Query.Limit))

	if len(options.Query.OrderBys) > 0 {
		for _, orderBy := range options.Query.OrderBys {
			q = q.OrderBy(orderBy.Field, orderBy.Desc)
		}
	} else {
		q = q.OrderBy("createdAt", true)
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
