package sqlstore

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func (s *SQLStore[T]) createTable(db *sql.DB, tableName string) error {
	var v T
	obj := &v
	val := reflect.ValueOf(obj).Elem()
	typ := val.Type()

	var fields []string
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldName := strings.ToLower(field.Name)
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

func (s *SQLStore[T]) sqlType(t reflect.Type) string {
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
