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

func (q *SQLQueryBuilder[T]) buildConditions(start ...int) ([]string, []interface{}, error) {
	var params []interface{}
	paramCounter := 1
	if len(start) > 0 {
		paramCounter += start[0]
	}
	var conditions []string

	for _, cond := range q.conditions {

		var param any
		var err error

		if cond.Equal != nil {
			fieldName := q.store.fieldName(cond.Equal.FieldName)
			placeHolder := q.store.placeholder(paramCounter)

			if reflect.TypeOf(cond.Equal.Value).Kind() == reflect.Slice {
				conditions = append(conditions, fmt.Sprintf("%s = ANY(%s)", fieldName, placeHolder))
				param, err = q.store.convertParam(cond.Equal.Value)
			} else if typ, hasTyp := q.store.fieldTypes[fieldName]; hasTyp && typ.Kind() == reflect.Slice {
				// "reverse" IN clause
				conditions = append(conditions, fmt.Sprintf("%s = ANY(%s)", placeHolder, fieldName))
				param, err = q.store.convertParam(cond.Equal.Value)
			} else {
				conditions = append(conditions, fmt.Sprintf("%s = %s", fieldName, placeHolder))
				param, err = q.store.convertParam(cond.Equal.Value)
			}

			params = append(params, param)
			paramCounter++
		}

		if err != nil {
			return nil, nil, err
		}

	}

	if len(q.after) > 0 {
		for i, afterValue := range q.after {
			fieldName := q.store.fieldName(q.orderFields[i])
			placeHolder := q.store.placeholder(paramCounter)
			if q.orderDescs[i] {
				conditions = append(conditions, fmt.Sprintf("%s < %s", fieldName, placeHolder))
			} else {
				conditions = append(conditions, fmt.Sprintf("%s > %s", fieldName, placeHolder))
			}
			params = append(params, afterValue)
			paramCounter++
		}
	}

	return conditions, params, nil
}
