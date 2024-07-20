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

	"github.com/singulatron/singulatron/localtron/datastore"

	generictypes "github.com/singulatron/singulatron/localtron/services/generic/types"
)

type FindOptions struct {
	Table      string
	Conditions []datastore.Condition
	UserId     string
	Public     bool
	OrderBys   []datastore.OrderBy
}

func (g *GenericService) Find(options FindOptions) ([]*generictypes.GenericObject, error) {
	if options.Table == "" {
		return nil, errors.New("no table name")
	}

	conditions := []datastore.Condition{}
	conditions = append(conditions, options.Conditions...)

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

	if len(options.OrderBys) > 1 {
		q.OrderBy(options.OrderBys[0], options.OrderBys...)
	} else if len(options.OrderBys) > 0 {
		q.OrderBy(options.OrderBys[0])
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
