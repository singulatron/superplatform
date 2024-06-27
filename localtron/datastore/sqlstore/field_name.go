package sqlstore

import (
	"strings"
)

func (s *SQLStore[T]) fieldName(fieldName string) string {
	if len(fieldName) == 0 {
		return ""
	}

	return strings.ToLower(fieldName[0:1]) + fieldName[1:]
}
