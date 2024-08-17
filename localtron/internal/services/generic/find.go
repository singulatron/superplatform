/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package genericservice

import (
	"errors"

	"github.com/singulatron/singulatron/sdk/go/datastore"

	generictypes "github.com/singulatron/singulatron/localtron/internal/services/generic/types"
)

func (g *GenericService) find(options generictypes.QueryOptions) ([]*generictypes.GenericObject, error) {
	if options.Table == "" {
		return nil, errors.New("no table name")
	}

	conditions := []datastore.Condition{}
	if options.Query != nil {
		conditions = append(conditions, options.Query.Conditions...)
	}

	conditions = append(conditions,
		datastore.Equal(datastore.Field("table"), options.Table),
	)

	if options.Public {
		conditions = append(conditions,
			datastore.Equal(datastore.Field("public"), true),
		)
	} else {
		if options.UserId == "" {
			return nil, errors.New("no user id when querying non public records")
		}
		conditions = append(conditions,
			datastore.Equal(datastore.Field("userId"), options.UserId),
		)
	}

	q := g.store.Query(
		conditions[0], conditions[1:]...,
	)

	if options.Query != nil {
		if len(options.Query.OrderBys) > 1 {
			q.OrderBy(options.Query.OrderBys[0], options.Query.OrderBys...)
		} else if len(options.Query.OrderBys) > 0 {
			q.OrderBy(options.Query.OrderBys[0])
		}

		if options.Query.Limit != 0 {
			q.Limit(options.Query.Limit)
		}

		if options.Query.After != nil {
			q.After(options.Query.After...)
		}
	}

	objectIs, err := q.Find()
	if err != nil {
		return nil, err
	}

	objects := []*generictypes.GenericObject{}
	for _, objectI := range objectIs {
		objects = append(objects, objectI.(*generictypes.GenericObject))
	}

	return objects, nil
}
