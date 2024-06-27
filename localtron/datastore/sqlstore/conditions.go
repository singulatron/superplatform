package sqlstore

import "fmt"

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
			conditions = append(conditions, fmt.Sprintf("%s = %s", cond.Equal.FieldName, q.store.placeholder(paramCounter)))
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
