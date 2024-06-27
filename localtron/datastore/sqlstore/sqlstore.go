package sqlstore

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/singulatron/singulatron/localtron/datastore"
)

type PlaceholderStyle int

const (
	QuestionMarkPlaceholder PlaceholderStyle = iota
	DollarSignPlaceholder
)

type Driver string

const (
	DriverMySQL    = "mysql"
	DriverPostGRES = "postgres"
)

type SQLStore[T datastore.Row] struct {
	db               *sql.DB
	mu               sync.RWMutex
	inTransaction    bool
	tx               *sql.Tx
	placeholderStyle PlaceholderStyle
	driverName       string
}

func NewSQLStore[T datastore.Row](driverName, connStr string, tableName string) (*SQLStore[T], error) {
	db, err := sql.Open(driverName, connStr)
	if err != nil {
		return nil, err
	}

	placeholderStyle := DollarSignPlaceholder
	if driverName == "mysql" {
		placeholderStyle = QuestionMarkPlaceholder
	}

	sstore := &SQLStore[T]{
		driverName:       driverName,
		placeholderStyle: placeholderStyle,
		db:               db,
	}

	if err := sstore.createTable(db, tableName); err != nil {
		panic(err)
	}

	return sstore, nil
}

func (s *SQLStore[T]) Create(obj T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	query, values, err := s.buildInsertQuery(obj)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, values...)
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
		return err
	}
	for _, obj := range objs {
		query, values, err := s.buildInsertQuery(obj)
		if err != nil {
			return err
		}
		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			if strings.Contains(err.Error(), "duplicate key value") {
				return datastore.ErrEntryAlreadyExists
			}
			return err
		}
	}
	return tx.Commit()
}

func (s *SQLStore[T]) Upsert(obj T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	query, values, err := s.buildUpsertQuery(obj)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, values...)
	return err
}

func (s *SQLStore[T]) UpsertMany(objs []T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	for _, obj := range objs {
		query, values, err := s.buildUpsertQuery(obj)
		if err != nil {
			return err
		}
		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return err
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
		return nil, err
	}

	return &SQLStore[T]{
		db:            s.db,
		inTransaction: true,
		tx:            tx,
	}, nil
}

func (s *SQLStore[T]) Commit() error {
	if !s.inTransaction {
		return errors.New("not in a transaction")
	}

	err := s.tx.Commit()
	if err != nil {
		return err
	}

	s.inTransaction = false
	return nil
}

func (s *SQLStore[T]) Rollback() error {
	if !s.inTransaction {
		return errors.New("not in a transaction")
	}

	err := s.tx.Rollback()
	if err != nil {
		return err
	}

	s.inTransaction = false
	return nil
}

func (s *SQLStore[T]) IsInTransaction() bool {
	return s.inTransaction
}

func (s *SQLStore[T]) convertParam(param any) (any, error) {
	switch reflect.TypeOf(param).Kind() {
	case reflect.Struct:
		if reflect.TypeOf(param) == reflect.TypeOf(time.Time{}) {
			return param, nil
		}
		bs, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		return string(bs), nil
	case reflect.Slice:
		switch s.driverName {
		case DriverMySQL:
			bs, err := json.Marshal(param)
			if err != nil {
				return nil, err
			}
			return string(bs), nil
		case DriverPostGRES:
			return pq.Array(param), nil
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

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fields = append(fields, strings.ToLower(field.Name))
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
		param := val.Field(i).Interface()
		param, err := s.convertParam(param)
		if err != nil {
			return "", nil, err
		}
		params = append(params, param)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		strings.ToLower(typ.Name()),
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

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldName := strings.ToLower(field.Name)
		fields = append(fields, fieldName)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
		param := val.Field(i).Interface()
		param, err := s.convertParam(param)
		if err != nil {
			return "", nil, err
		}
		params = append(params, param)
		updateFields = append(updateFields, fmt.Sprintf("%s=EXCLUDED.%s", fieldName, fieldName))
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (%s) DO UPDATE SET %s;",
		strings.ToLower(typ.Name()),
		strings.Join(fields, ", "),
		strings.Join(placeholders, ", "),
		strings.ToLower(typ.Field(0).Name),
		strings.Join(updateFields, ", "))

	return query, params, nil
}

type SQLQueryBuilder[T datastore.Row] struct {
	store        *SQLStore[T]
	conditions   []datastore.Condition
	orderField   string
	orderDesc    bool
	limit        int
	after        []any
	selectFields []string
}

func (q *SQLQueryBuilder[T]) OrderBy(field string, desc bool) datastore.QueryBuilder[T] {
	q.orderField = field
	q.orderDesc = desc
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
func (q *SQLQueryBuilder[T]) Find() ([]T, error) {
	query, params := q.buildSelectQuery()

	rows, err := q.store.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []T
	tType := reflect.TypeOf((*T)(nil)).Elem()

	for rows.Next() {
		obj := reflect.New(tType).Elem()
		fields := make([]interface{}, tType.NumField())

		for i := 0; i < tType.NumField(); i++ {
			field := obj.Field(i)
			fieldType := field.Type()

			switch {
			case fieldType.Kind() == reflect.Slice && fieldType.Elem().Kind() == reflect.String:
				var str sql.NullString
				fields[i] = &str
			case fieldType.Kind() == reflect.Struct && fieldType != reflect.TypeOf(time.Time{}):
				var str sql.NullString
				fields[i] = &str
			default:
				fields[i] = field.Addr().Interface()
			}
		}

		err := rows.Scan(fields...)
		if err != nil {
			return nil, err
		}

		for i := 0; i < tType.NumField(); i++ {
			field := obj.Field(i)
			fieldType := field.Type()

			switch {
			case fieldType.Kind() == reflect.Slice && fieldType.Elem().Kind() == reflect.String:
				str, ok := fields[i].(*sql.NullString)
				if ok && str.Valid {
					field.Set(reflect.ValueOf(strings.Split(str.String, ",")))
				} else {
					field.Set(reflect.Zero(fieldType))
				}
			case fieldType.Kind() == reflect.Struct && fieldType != reflect.TypeOf(time.Time{}):
				str, ok := fields[i].(*sql.NullString)
				if ok && str.Valid {
					newField := reflect.New(fieldType).Interface()
					err := json.Unmarshal([]byte(str.String), newField)
					if err != nil {
						return nil, err
					}
					field.Set(reflect.ValueOf(newField).Elem())
				}
			}
		}

		result = append(result, obj.Interface().(T))
	}

	return result, nil
}

func (q *SQLQueryBuilder[T]) FindOne() (T, bool, error) {
	query, params := q.buildSelectQuery()
	query += " LIMIT 1"

	row := q.store.db.QueryRow(query, params...)
	var obj T
	err := row.Scan(&obj)
	if err != nil {
		if err == sql.ErrNoRows {
			var empty T
			return empty, false, nil
		}
		return obj, false, err
	}

	return obj, true, nil
}

func (q *SQLQueryBuilder[T]) Count() (int64, error) {
	query, params := q.buildSelectQuery()
	query = fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS subquery", query)

	var count int64
	err := q.store.db.QueryRow(query, params...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q *SQLQueryBuilder[T]) Update(obj T) error {
	query := q.buildUpdateQuery(obj)
	_, err := q.store.db.Exec(query)
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
	query := q.buildUpdateFieldsQuery(fields)
	_, err := q.store.db.Exec(query)
	return err
}

func (q *SQLQueryBuilder[T]) Delete() error {
	query, values := q.buildDeleteQuery()
	_, err := q.store.db.Exec(query, values...)
	return err
}

func (q *SQLQueryBuilder[T]) buildSelectQuery() (string, []interface{}) {
	var conditions []string
	var params []interface{}
	paramCounter := 1
	placeholder := func(counter int) string {
		switch q.store.placeholderStyle {
		case QuestionMarkPlaceholder:
			return "?"
		case DollarSignPlaceholder:
			return fmt.Sprintf("$%d", counter)
		default:
			return "?"
		}
	}

	for _, cond := range q.conditions {
		if cond.Equal != nil {
			conditions = append(conditions, fmt.Sprintf("%s = %s", cond.Equal.FieldName, placeholder(paramCounter)))
			params = append(params, cond.Equal.Value)
			paramCounter++
		}
	}

	var query string
	if len(q.selectFields) > 0 {
		query = fmt.Sprintf("SELECT %s FROM %s", strings.Join(q.selectFields, ", "), strings.ToLower(reflect.TypeOf(new(T)).Elem().Name()))
	} else {
		query = fmt.Sprintf("SELECT * FROM %s", strings.ToLower(reflect.TypeOf(new(T)).Elem().Name()))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if q.orderField != "" {
		query += fmt.Sprintf(" ORDER BY %s", q.orderField)
		if q.orderDesc {
			query += " DESC"
		}
	}

	if len(q.after) > 0 {
		query += fmt.Sprintf(" OFFSET %s", placeholder(paramCounter))
		params = append(params, q.after[0])
		paramCounter++
	}

	if q.limit > 0 {
		query += fmt.Sprintf(" LIMIT %s", placeholder(paramCounter))
		params = append(params, q.limit)
	}

	return query, params
}

func (q *SQLQueryBuilder[T]) buildUpdateQuery(obj T) string {
	val := reflect.ValueOf(obj).Elem()
	typ := val.Type()

	var sets []string
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		sets = append(sets, fmt.Sprintf("%s = '%v'", strings.ToLower(field.Name), val.Field(i).Interface()))
	}

	var conditions []string
	for _, cond := range q.conditions {
		if cond.Equal != nil {
			conditions = append(conditions, fmt.Sprintf("%s = '%v'", cond.Equal.FieldName, cond.Equal.Value))
		}
	}

	query := fmt.Sprintf("UPDATE %s SET %s", strings.ToLower(typ.Name()), strings.Join(sets, ", "))
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	return query
}

func (q *SQLQueryBuilder[T]) buildUpdateFieldsQuery(fields map[string]interface{}) string {
	var sets []string
	for key, value := range fields {
		sets = append(sets, fmt.Sprintf("%s = '%v'", key, value))
	}

	var conditions []string
	for _, cond := range q.conditions {
		if cond.Equal != nil {
			conditions = append(conditions, fmt.Sprintf("%s = '%v'", cond.Equal.FieldName, cond.Equal.Value))
		}
	}

	query := fmt.Sprintf("UPDATE %s SET %s", strings.ToLower(reflect.TypeOf(new(T)).Elem().Name()), strings.Join(sets, ", "))
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	return query
}

func (q *SQLQueryBuilder[T]) buildDeleteQuery() (string, []interface{}) {
	var conditions []string
	var params []interface{}
	paramCounter := 1

	// Define a function for generating placeholders
	placeholder := func(counter int) string {
		switch q.store.placeholderStyle {
		case QuestionMarkPlaceholder:
			return "?"
		case DollarSignPlaceholder:
			return fmt.Sprintf("$%d", counter)
		default:
			return "?"
		}
	}

	for _, cond := range q.conditions {
		if cond.Equal != nil {
			conditions = append(conditions, fmt.Sprintf("%s = %s", cond.Equal.FieldName, placeholder(paramCounter)))
			params = append(params, cond.Equal.Value)
			paramCounter++
		}
	}

	query := fmt.Sprintf("DELETE FROM %s", strings.ToLower(reflect.TypeOf(new(T)).Elem().Name()))
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	return query, params
}
