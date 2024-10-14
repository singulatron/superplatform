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

	// Recursive function to process struct fields (including embedded fields)
	var processFields func(reflect.Type)
	processFields = func(typ reflect.Type) {
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if !field.IsExported() {
				continue
			}

			// Handle embedded structs by recursion
			if field.Anonymous && field.Type.Kind() == reflect.Struct {
				processFields(field.Type)
				continue
			}

			// Use the json tag for the field name, if available, and strip out ",omitempty"
			fieldName := field.Tag.Get("json")
			if fieldName == "" {
				fieldName = s.fieldName(field.Name)
			} else if idx := strings.Index(fieldName, ","); idx != -1 {
				fieldName = fieldName[:idx]
			}
			fieldName = escape(fieldName)

			// Map field type to SQL type
			fieldType := s.sqlType(field.Type)

			fields = append(fields, fmt.Sprintf("%s %s", fieldName, fieldType))
		}
	}

	// Process all fields of the instance's struct
	processFields(typ)

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
			if elemType != "JSONB" {
				// You can store as JSONB, or as native array if preferred
				return fmt.Sprintf("%s[]", elemType) // Use PostgreSQL array type
			}

			// Default to JSONB for complex slices or fallback
			return "JSONB"
		}

	default:
		return "TEXT" // Default to TEXT for unknown types
	}
	return "TEXT"
}
