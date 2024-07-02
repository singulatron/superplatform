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
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log/slog"
	"reflect"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/logger"

	"github.com/pkg/errors"
)

type PlaceholderStyle int

const (
	InvalidPlaceholder PlaceholderStyle = iota
	QuestionMarkPlaceholder
	DollarSignPlaceholder
)

type Driver string

const (
	DriverMySQL    = "mysql"
	DriverPostGRES = "postgres"
)

type SQLStore[T datastore.Row] struct {
	db               *DebugDB
	mu               sync.RWMutex
	inTransaction    bool
	tx               *DebugTx
	placeholderStyle PlaceholderStyle
	driverName       string
	tableName        string
	fieldTypes       map[string]reflect.Type
	idFieldName      string
}

func NewSQLStore[T datastore.Row](driverName, connStr string, tableName string, debug bool) (*SQLStore[T], error) {
	db, err := sql.Open(driverName, connStr)
	if err != nil {
		return nil, errors.Wrap(err, "error opening sql db")
	}

	placeholderStyle := DollarSignPlaceholder
	if driverName == "mysql" {
		placeholderStyle = QuestionMarkPlaceholder
	}

	if tableName == "" {
		tableName = reflect.TypeOf(new(T)).Elem().Name()
	}

	sstore := &SQLStore[T]{
		driverName:       driverName,
		tableName:        tableName,
		placeholderStyle: placeholderStyle,
		db:               NewDebugDB(db, tableName),
		fieldTypes:       map[string]reflect.Type{},
	}
	sstore.db.Debug = debug

	if err := sstore.createTable(sstore.db, tableName); err != nil {
		return nil, errors.Wrap(err, "error creating table")
	}

	var v T

	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	fieldName := sstore.fieldName(typ.Field(0).Name)
	sstore.idFieldName = fieldName

	_, err = sstore.db.Exec(fmt.Sprintf("ALTER TABLE %v ADD CONSTRAINT %v_%v_unique UNIQUE (%v);",
		sstore.tableName,
		sstore.tableName,
		fieldName,
		fieldName,
	))
	if err != nil {
		logger.Debug("Error adding constraint", slog.Any("error", err))
	}

	return sstore, nil
}

func (s *SQLStore[T]) SetDebug(debug bool) {
	s.db.Debug = true
}

func (s *SQLStore[T]) Create(obj T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	query, values, err := s.buildInsertQuery(obj)
	if err != nil {
		return err
	}

	if s.inTransaction {
		_, err = s.tx.Exec(query, values...)
	} else {
		_, err = s.db.Exec(query, values...)
	}
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return datastore.ErrEntryAlreadyExists
		}
		return err
	}

	return nil
}

func (s *SQLStore[T]) CreateMany(objs []T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "error beginning transacion in create many")
	}
	for _, obj := range objs {
		query, values, err := s.buildInsertQuery(obj)
		if err != nil {
			return errors.Wrap(err, "error building insert query in create many")
		}
		if s.inTransaction {
			_, err = s.tx.Exec(query, values...)
		} else {
			_, err = s.db.Exec(query, values...)
		}
		if err != nil {
			tx.Rollback()
			if strings.Contains(err.Error(), "duplicate key value") {
				return datastore.ErrEntryAlreadyExists
			}
			return errors.Wrap(err, "error executing query in create many")
		}
	}

	return tx.Commit()
}

func (s *SQLStore[T]) Upsert(obj T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	query, values, err := s.buildUpsertQuery(obj)
	if err != nil {
		return errors.Wrap(err, "error building query in upsert")
	}

	if s.inTransaction {
		_, err = s.tx.Exec(query, values...)
	} else {
		_, err = s.db.Exec(query, values...)
	}

	return err
}

func (s *SQLStore[T]) UpsertMany(objs []T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "error beginning transaction in upsert many")
	}
	for _, obj := range objs {
		query, values, err := s.buildUpsertQuery(obj)
		if err != nil {
			return errors.Wrap(err, "error building query in upsert many")
		}
		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "error executing query in upsert many")
		}
	}
	return tx.Commit()
}

func (s *SQLStore[T]) Query(condition datastore.Condition, conditions ...datastore.Condition) datastore.QueryBuilder[T] {
	return &SQLQueryBuilder[T]{
		store:      s,
		conditions: append([]datastore.Condition{condition}, conditions...),
	}
}

func (s *SQLStore[T]) BeginTransaction() (datastore.DataStore[T], error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.inTransaction {
		return nil, errors.New("already in a transaction")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "error beginning transaction")
	}

	return &SQLStore[T]{
		db:               s.db,
		tableName:        s.tableName,
		driverName:       s.driverName,
		inTransaction:    true,
		tx:               tx,
		fieldTypes:       s.fieldTypes,
		placeholderStyle: s.placeholderStyle,
	}, nil
}

func (s *SQLStore[T]) Commit() error {
	if !s.inTransaction {
		return errors.New("not in a transaction")
	}

	err := s.tx.Commit()
	if err != nil {
		return errors.Wrap(err, "error committing transaction")
	}

	s.inTransaction = false
	return nil
}

func (s *SQLStore[T]) Rollback() error {
	defer func() {
		s.inTransaction = false
	}()

	if !s.inTransaction {
		return errors.New("not in a transaction")
	}

	err := s.tx.Rollback()
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLStore[T]) IsInTransaction() bool {
	return s.inTransaction
}

func (s *SQLStore[T]) convertParam(param any) (any, error) {
	t := reflect.TypeOf(param)
	v := reflect.ValueOf(param)

	switch t.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return nil, nil
		}
		return s.convertParam(reflect.ValueOf(param).Elem().Interface())
	case reflect.Struct:
		if reflect.TypeOf(param) == reflect.TypeOf(time.Time{}) {
			return param, nil
		}
		bs, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		return string(bs), nil
	case reflect.Map:
		bs, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		return string(bs), nil
	case reflect.Slice:
		if v.Len() == 0 {
			switch s.driverName {
			case DriverMySQL:
				return "[]", nil
			case DriverPostGRES:
				return pq.Array(nil), nil
			default:
				return nil, fmt.Errorf("unrecognized driver: '%v'", s.driverName)
			}
		}

		switch s.driverName {
		case DriverMySQL:
			bs, err := json.Marshal(param)
			if err != nil {
				return nil, err
			}
			return string(bs), nil
		case DriverPostGRES:
			return pq.Array(param), nil
		default:
			return nil, fmt.Errorf("unrecognized driver: '%v'", s.driverName)
		}
	}

	return param, nil
}

func (s *SQLStore[T]) buildInsertQuery(obj T) (string, []interface{}, error) {
	val := reflect.ValueOf(obj)
	typ := val.Type()

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	var fields []string
	var placeholders []string
	var params []interface{}
	paramCounter := 1

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if !field.IsExported() {
			continue
		}
		fields = append(fields, s.fieldName(field.Name))
		placeholders = append(placeholders, s.placeholder(paramCounter))
		param := val.Field(i).Interface()
		param, err := s.convertParam(param)
		if err != nil {
			return "", nil, err
		}
		params = append(params, param)
		paramCounter++
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		strings.ToLower(s.tableName),
		strings.Join(fields, ", "),
		strings.Join(placeholders, ", "))

	return query, params, nil
}

func (s *SQLStore[T]) buildUpsertQuery(obj T) (string, []interface{}, error) {
	val := reflect.ValueOf(obj)
	typ := val.Type()

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	var fields []string
	var placeholders []string
	var updateFields []string
	var params []interface{}
	paramCounter := 1

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if !field.IsExported() {
			continue
		}
		fieldName := s.fieldName(field.Name)

		fields = append(fields, fieldName)
		placeholders = append(placeholders, s.placeholder(paramCounter))
		param := val.Field(i).Interface()
		param, err := s.convertParam(param)
		if err != nil {
			return "", nil, err
		}
		params = append(params, param)
		updateFields = append(updateFields, fmt.Sprintf("%s=EXCLUDED.%s", fieldName, fieldName))
		paramCounter++
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (%s) DO UPDATE SET %s;",
		strings.ToLower(s.tableName),
		strings.Join(fields, ", "),
		strings.Join(placeholders, ", "),
		strings.ToLower(s.fieldName(typ.Field(0).Name)),
		strings.Join(updateFields, ", "))

	return query, params, nil
}

type SQLQueryBuilder[T datastore.Row] struct {
	store        *SQLStore[T]
	conditions   []datastore.Condition
	orderFields  []string
	orderDescs   []bool
	limit        int
	after        []any
	selectFields []string
}

func (q *SQLQueryBuilder[T]) OrderBy(field string, desc bool) datastore.QueryBuilder[T] {
	q.orderFields = append(q.orderFields, field)
	q.orderDescs = append(q.orderDescs, desc)
	return q
}

func (q *SQLQueryBuilder[T]) Limit(limit int) datastore.QueryBuilder[T] {
	q.limit = limit
	return q
}

func (q *SQLQueryBuilder[T]) After(value ...any) datastore.QueryBuilder[T] {
	q.after = value
	return q
}

func (q *SQLQueryBuilder[T]) Select(fields ...string) datastore.QueryBuilder[T] {
	q.selectFields = fields
	return q
}

type GenericArray struct {
	Array interface{}
}

func (a *GenericArray) Scan(src interface{}) error {
	return pq.Array(a.Array).Scan(src)
}

func (a *GenericArray) Value() (driver.Value, error) {
	return pq.Array(a.Array).Value()
}

func (q *SQLQueryBuilder[T]) Find() ([]T, error) {
	query, params, err := q.buildSelectQuery()
	if err != nil {
		return nil, errors.Wrap(err, "error building select query")
	}

	rows, err := q.store.db.Query(query, params...)
	if err != nil {
		return nil, errors.Wrap(err, "error querying")
	}
	defer rows.Close()

	var result []T

	var v T
	tType := reflect.TypeOf(v)
	tIsPointer := tType.Kind() == reflect.Pointer

	var safeNumFieldsType reflect.Type
	if tIsPointer {
		safeNumFieldsType = tType.Elem()
	} else {
		safeNumFieldsType = tType
	}

	for rows.Next() {
		var obj reflect.Value
		if tIsPointer {
			obj = reflect.New(tType.Elem()).Elem()
		} else {
			obj = reflect.New(tType).Elem()
		}

		fields := make([]interface{}, safeNumFieldsType.NumField())

		for i := 0; i < safeNumFieldsType.NumField(); i++ {
			field := obj.Field(i)
			if !field.IsExported() {
				continue
			}
			fieldType := field.Type()

			switch {
			case fieldType.Kind() == reflect.Slice:
				// Create a GenericArray with the appropriate type
				elemType := fieldType.Elem()
				slicePtr := reflect.New(reflect.SliceOf(elemType)).Interface()
				fields[i] = &GenericArray{Array: slicePtr}
			case fieldType.Kind() == reflect.Pointer:
				var str sql.NullString
				fields[i] = &str
			case fieldType.Kind() == reflect.Struct && fieldType != reflect.TypeOf(time.Time{}):
				var str sql.NullString
				fields[i] = &str
			case fieldType == reflect.TypeOf(time.Time{}):
				fields[i] = new(sql.NullTime)
			default:
				fields[i] = field.Addr().Interface()
			}
		}

		err := rows.Scan(fields...)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning")
		}

		for i := 0; i < safeNumFieldsType.NumField(); i++ {
			field := obj.Field(i)
			if !field.IsExported() {
				continue
			}

			fieldType := field.Type()

			switch {
			case fieldType.Kind() == reflect.Slice:
				// Set the scanned slice to the appropriate field
				genericArray, ok := fields[i].(*GenericArray)
				if ok {
					field.Set(reflect.ValueOf(genericArray.Array).Elem())
				} else {
					field.Set(reflect.Zero(fieldType))
				}
			case fieldType.Kind() == reflect.Pointer:
				str, ok := fields[i].(*sql.NullString)
				if ok && str.Valid {
					newField := reflect.New(fieldType.Elem()).Interface()
					err := json.Unmarshal([]byte(str.String), newField)
					if err != nil {
						return nil, errors.Wrap(err, "error unmarshaling struct")
					}
					field.Set(reflect.ValueOf(newField))
				} else {
					bin, ok := fields[i].(*[]uint8)
					if ok && bin != nil {
						newField := reflect.New(fieldType.Elem()).Interface()
						err := json.Unmarshal(*bin, newField)
						if err != nil {
							return nil, errors.Wrap(err, "error unmarshaling JSONB binary data")
						}
						field.Set(reflect.ValueOf(newField))
					} else {
						field.Set(reflect.Zero(fieldType))
					}
				}
			case fieldType.Kind() == reflect.Struct && fieldType != reflect.TypeOf(time.Time{}):
				str, ok := fields[i].(*sql.NullString)
				if ok && str.Valid {
					newField := reflect.New(fieldType).Interface()
					err := json.Unmarshal([]byte(str.String), newField)
					if err != nil {
						return nil, errors.Wrap(err, "error unmarshaling struct")
					}
					field.Set(reflect.ValueOf(newField).Elem())
				} else {
					// Handle JSONB binary data
					bin, ok := fields[i].(*[]uint8)
					if ok && bin != nil {
						newField := reflect.New(fieldType).Interface()
						err := json.Unmarshal(*bin, newField)
						if err != nil {
							return nil, errors.Wrap(err, "error unmarshaling JSONB binary data")
						}
						field.Set(reflect.ValueOf(newField).Elem())
					} else {
						field.Set(reflect.Zero(fieldType))
					}
				}
			case fieldType == reflect.TypeOf(time.Time{}):
				nullTime, ok := fields[i].(*sql.NullTime)
				if ok && nullTime.Valid {
					field.Set(reflect.ValueOf(nullTime.Time.UTC()))
				} else {
					field.Set(reflect.Zero(fieldType))
				}
			}
		}

		if tIsPointer {
			val := obj.Interface()
			valPtr := reflect.New(reflect.TypeOf(val))
			valPtr.Elem().Set(reflect.ValueOf(val))
			result = append(result, valPtr.Interface().(T))
		} else {
			result = append(result, obj.Interface().(T))
		}
	}

	return result, nil
}

func (q *SQLQueryBuilder[T]) FindOne() (T, bool, error) {
	var def T
	res, err := q.Find()
	if err != nil {
		return def, false, err
	}

	if len(res) == 0 {
		return def, false, nil
	}

	return res[0], true, nil
}

func (q *SQLQueryBuilder[T]) Count() (int64, error) {
	query, params, err := q.buildSelectQuery()
	if err != nil {
		return 0, err
	}
	query = fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS subquery", query)

	var count int64
	err = q.store.db.QueryRow(query, params...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q *SQLQueryBuilder[T]) Update(obj T) error {
	query, params, err := q.buildUpdateQuery(obj)
	if err != nil {
		return err
	}
	_, err = q.store.db.Exec(query, params...)
	return err
}

func (q *SQLQueryBuilder[T]) Upsert(obj T) error {
	query, values, err := q.store.buildUpsertQuery(obj)
	if err != nil {
		return err
	}
	_, err = q.store.db.Exec(query, values...)
	return err
}

func (q *SQLQueryBuilder[T]) UpdateFields(fields map[string]interface{}) error {
	query, params, err := q.buildUpdateFieldsQuery(fields)
	if err != nil {
		return err
	}
	_, err = q.store.db.Exec(query, params...)
	return err
}

func (q *SQLQueryBuilder[T]) Delete() error {
	query, values, err := q.buildDeleteQuery()
	if err != nil {
		return err
	}
	_, err = q.store.db.Exec(query, values...)
	return err
}

func (q *SQLQueryBuilder[T]) buildSelectQuery() (string, []interface{}, error) {
	conditions, params, err := q.buildConditions()
	if err != nil {
		return "", nil, err
	}
	paramCounter := len(params) + 1

	var query string
	if len(q.selectFields) > 0 {
		selectFields := []string{}
		for _, selectField := range q.selectFields {
			selectFields = append(selectFields, q.store.fieldName(selectField))
		}
		query = fmt.Sprintf("SELECT %s FROM %s", strings.Join(selectFields, ", "), q.store.db.tableName)
	} else {
		query = fmt.Sprintf("SELECT * FROM %s", strings.ToLower(q.store.db.tableName))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if q.orderFields != nil {
		for i, orderField := range q.orderFields {
			orderDesc := q.orderDescs[i]
			query += fmt.Sprintf(" ORDER BY %s", q.store.fieldName(orderField))
			if orderDesc {
				query += " DESC"
			}
		}
	}

	if q.limit > 0 {
		query += fmt.Sprintf(" LIMIT %s", q.store.placeholder(paramCounter))
		params = append(params, q.limit)
	}

	return query, params, nil
}

func (q *SQLQueryBuilder[T]) buildUpdateQuery(obj T) (string, []any, error) {
	val := reflect.ValueOf(obj)
	typ := val.Type()

	if typ.Kind() == reflect.Pointer {
		val = reflect.ValueOf(obj).Elem()
		typ = val.Type()
	} else {
		val = reflect.ValueOf(obj)
		typ = val.Type()
	}

	var sets []string
	var params []interface{}
	paramCounter := 1

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if !field.IsExported() {
			continue
		}
		fieldName := q.store.fieldName(field.Name)

		param, err := q.store.convertParam(val.Field(i).Interface())
		if err != nil {
			return "", nil, err
		}
		placeHolder := q.store.placeholder(paramCounter)

		sets = append(sets, fmt.Sprintf("%s = %v", fieldName, placeHolder))
		params = append(params, param)
		paramCounter++
	}

	conditions, conditionParams, err := q.buildConditions(paramCounter - 1)
	if err != nil {
		return "", nil, err
	}
	params = append(params, conditionParams...)

	query := fmt.Sprintf("UPDATE %s SET %s", q.store.tableName, strings.Join(sets, ", "))
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	return query, params, nil
}

func (q *SQLQueryBuilder[T]) buildUpdateFieldsQuery(fields map[string]interface{}) (string, []any, error) {

	var sets []string
	var params []interface{}
	paramCounter := 1

	for key, value := range fields {
		fieldName := q.store.fieldName(key)
		param, err := q.store.convertParam(value)
		if err != nil {
			return "", nil, err
		}
		placeHolder := q.store.placeholder(paramCounter)
		sets = append(sets, fmt.Sprintf("%s = %v", fieldName, placeHolder))
		params = append(params, param)
		paramCounter++
	}

	conditions, conditionParams, err := q.buildConditions(paramCounter - 1)
	if err != nil {
		return "", nil, err
	}
	params = append(params, conditionParams...)

	query := fmt.Sprintf("UPDATE %s SET %s", q.store.tableName, strings.Join(sets, ", "))
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	return query, params, nil
}

func (q *SQLQueryBuilder[T]) buildDeleteQuery() (string, []interface{}, error) {
	conditions, params, err := q.buildConditions()
	if err != nil {
		return "", nil, err
	}

	query := fmt.Sprintf("DELETE FROM %s", q.store.tableName)
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	return query, params, nil
}
