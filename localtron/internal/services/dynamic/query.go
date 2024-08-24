/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dynamicservice

import (
	"errors"

	"github.com/singulatron/singulatron/sdk/go/datastore"

	dynamictypes "github.com/singulatron/singulatron/localtron/internal/services/dynamic/types"
)

func (g *DynamicService) query(readers []string, options dynamictypes.QueryOptions) ([]*dynamictypes.Object, error) {
	if options.Table == "" {
		return nil, errors.New("no table name")
	}

	conditions := []datastore.Filter{}
	if options.Query != nil {
		conditions = append(conditions, options.Query.Filters...)
	}

	conditions = append(conditions,
		datastore.Equals(datastore.Field("table"), options.Table),
	)

	conditions = append(conditions,
		datastore.Equals(datastore.Field("readers"), readers),
	)

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

	objects := []*dynamictypes.Object{}
	for _, objectI := range objectIs {
		objects = append(objects, objectI.(*dynamictypes.Object))
	}

	return objects, nil
}
