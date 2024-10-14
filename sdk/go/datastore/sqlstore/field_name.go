/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package sqlstore

import (
	"fmt"
	"strings"
)

func (s *SQLStore) fieldName(fieldName string) string {
	if len(fieldName) == 0 {
		return ""
	}

	fieldParts := strings.Split(fieldName, ".")

	for i, v := range fieldParts {
		f := escape(strings.ToLower(v[0:1]) + v[1:])

		if i == 0 {
			fieldParts[i] = f
			continue
		}

		if i == len(fieldParts)-1 {
			fieldParts[i] = "->>" + fmt.Sprintf("'%v'", f)
		} else {
			fieldParts[i] = "->" + f
		}
	}

	return strings.Join(fieldParts, "")
}

func escape(columnName string) string {
	toEscape := false
	switch columnName {
	case "table":
		toEscape = true
	}

	if toEscape {
		return fmt.Sprintf(`"%v"`, columnName)
	}

	return columnName
}
