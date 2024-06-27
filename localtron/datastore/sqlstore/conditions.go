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
package sqlstore

import (
	"fmt"
	"reflect"
)

func (s *SQLStore[T]) placeholder(counter int) string {
	switch s.placeholderStyle {
	case QuestionMarkPlaceholder:
		return "?"
	case DollarSignPlaceholder:
		return fmt.Sprintf("$%d", counter)
	default:
		return "?"
	}
}

func (q *SQLQueryBuilder[T]) buildConditions() ([]string, []interface{}, error) {
	var params []interface{}
	paramCounter := 1
	var conditions []string
	for _, cond := range q.conditions {
		if cond.Equal != nil {
			if reflect.TypeOf(cond.Equal.Value).Kind() == reflect.Slice {
				conditions = append(conditions, fmt.Sprintf("%s = ANY(%s)", q.store.fieldName(cond.Equal.FieldName), q.store.placeholder(paramCounter)))
			} else {
				conditions = append(conditions, fmt.Sprintf("%s = %s", q.store.fieldName(cond.Equal.FieldName), q.store.placeholder(paramCounter)))
			}
			param, err := q.store.convertParam(cond.Equal.Value)
			if err != nil {
				return nil, nil, err
			}
			params = append(params, param)
			paramCounter++
		}

	}

	return conditions, params, nil
}
