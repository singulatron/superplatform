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
	"strings"
	"time"
)

func (s *SQLStore) createTable(instance any, db *DebugDB, tableName string) error {
	typ := reflect.TypeOf(instance)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	var fields []string

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !field.IsExported() {
			continue
		}
		fieldName := s.fieldName(field.Name)
		s.fieldTypes[fieldName] = field.Type
		fieldType := s.sqlType(field.Type)

		fields = append(fields, fmt.Sprintf("%s %s", fieldName, fieldType))
	}

	if tableName == "" {
		tableName = strings.ToLower(typ.Name())
	}
	createQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, strings.Join(fields, ", "))

	_, err := db.Exec(createQuery)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLStore) sqlType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.String:
		return "TEXT"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "INTEGER"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "INTEGER"
	case reflect.Float32, reflect.Float64:
		return "REAL"
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.Map:
		switch s.driverName {
		case DriverPostGRES:
			return "JSONB"
		case DriverMySQL:
			return "JSON"
		}
	case reflect.Struct:
		if t == reflect.TypeOf(time.Time{}) {
			return "TIMESTAMP"
		}
		switch s.driverName {
		case DriverPostGRES:
			return "JSONB"
		case DriverMySQL:
			return "JSON"
		}
	case reflect.Ptr:
		return s.sqlType(t.Elem())
	case reflect.Slice:
		switch s.driverName {
		case DriverMySQL:
			return "JSON"
		case DriverPostGRES:
			elemType := s.sqlType(t.Elem())
			if elemType != "" {
				return fmt.Sprintf("%s[]", elemType)
			}
			// ??
		}

	default:
		return "TEXT" // Default to TEXT for unknown types
	}
	return "TEXT"
}
