/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptservice

import (
	"github.com/singulatron/singulatron/localtron/datastore"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

type ListPromptOptions struct {
	Query *datastore.Query `json:"query"`
}

func (p *PromptService) ListPrompts(options *ListPromptOptions) ([]*prompttypes.Prompt, int64, error) {
	q := p.promptsStore.Query(
		options.Query.Conditions[0], options.Query.Conditions[1:]...,
	).Limit(int(options.Query.Limit))

	if len(options.Query.OrderBys) > 0 {
		for _, orderBy := range options.Query.OrderBys {
			q = q.OrderBy(orderBy.Field, orderBy.Desc)
		}
	} else {
		q = q.OrderBy("createdAt", false)
	}

	if options.Query.After != nil {
		q = q.After(options.Query.After...)
	}

	resI, err := q.Find()
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

	res := []*prompttypes.Prompt{}
	for _, v := range resI {
		res = append(res, v.(*prompttypes.Prompt))
	}

	return res, count, nil
}
