/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package datastore

import (
	"errors"
	"fmt"
	"time"
)

type Row interface {
	GetId() string
}

var (
	ErrEntryAlreadyExists = errors.New("entry already exists")
)

type DataStore interface {
	/*
	 * Create an object.
	 * Returns ErrEntryAlreadyExists if the object already exists.
	 */
	Create(obj Row) error
	/* Create many objects
	* Returns ErrEntryAlreadyExists if any of the objects are already in set,
	* and no object will be created.
	 */
	CreateMany(objs []Row) error
	/* Create or Update an object */
	Upsert(obj Row) error
	/* Create or Update many objects */
	UpsertMany(objs []Row) error

	Query(filters ...Filter) QueryBuilder

	BeginTransaction() (DataStore, error)
	Commit() error
	Rollback() error
	IsInTransaction() bool

	SetDebug(debug bool)
}

type QueryBuilder interface {
	OrderBy(options ...OrderBy) QueryBuilder
	Limit(limit int64) QueryBuilder
	After(value ...any) QueryBuilder

	Select(fields ...string) QueryBuilder
	Find() ([]Row, error)
	FindOne() (Row, bool, error)
	Count() (int64, error)

	// Update by query. Errors if no update happens
	Update(obj Row) error
	// Upsert tries to update by query, and if no update appened, calls create.
	Upsert(obj Row) error
	UpdateFields(fields map[string]interface{}) error
	Delete() error
}

type Op string

const (
	// OpEquals selects objects where value of a field equals (=) to the specified value in the query.
	OpEquals Op = "equals"

	// OpContains selects all objects where the field(s) values contain a particular string or slice element.
	OpContains Op = "contains"

	// OpStartsWith selects all objects where the field(s) values start with a particular string.
	OpStartsWith Op = "startsWith"

	// OpIntersects selects objects where the slice value of a field intersects with the slice value in the query.
	OpIntersects Op = "intersects"
)

type Filter struct {
	Fields []string `json:"fields,omitempty"`

	Values []any `json:"values,omitempty"`

	Op Op `json:"op"`
}

func (c Filter) FieldIs(fieldName string) bool {
	for _, field := range c.Fields {
		if fieldName == field {
			return true
		}
	}
	return false
}

// Query as a type is not used in the DataStore interface but mostly to accept
// a DataStore query through a HTTP API
type Query struct {
	// Filters are filtering options of a query. It is advised to use
	// It's advised to use helper functions in your respective client library such as filter constructors (`all`, `equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
	Filters []Filter `json:"filters,omitempty"`

	// After is used for paginations. Instead of offset-based pagination,
	// we support cursor-based pagination because it works better in a scalable,
	// distributed environment.
	After []any `json:"after,omitempty"`

	// Limit the number of records in the result set.
	Limit int64 `json:"limit,omitempty"`

	// OrderBys order the result set.
	OrderBys []OrderBy `json:"orderBys,omitempty"`

	// Count true means return the count of the dataset filtered by Filters
	// without after or limit.
	Count bool `json:"count,omitempty"`
}

func (q *Query) HasFieldFilter(fieldName string) bool {
	for _, v := range q.Filters {
		if v.FieldIs(fieldName) {
			return true
		}
	}

	return false
}

type OrderBy struct {
	// The field by which to order the results
	Field string `json:"field,omitempty"`

	// Desc indicates whether the sorting should be in descending order.
	Desc bool `json:"desc,omitempty"`

	// Randomize indicates that the results should be randomized instead of ordered by the `field` and `desc` criteria
	Randomize bool `json:"randomize,omitempty"`
}

// random order. not advised for large datasets due to its slow speed
// in a distributed setting
func OrderByRandom() OrderBy {
	return OrderBy{
		Randomize: true,
	}
}

func OrderByField(field string, desc bool) OrderBy {
	return OrderBy{
		Field: field,
		Desc:  desc,
	}
}

type AllMatch struct {
}

func Equals(fields []string, value any) Filter {
	return Filter{
		Fields: fields,
		Values: []any{value},
		Op:     OpEquals,
	}
}

func Intersects(fields []string, values []any) Filter {
	return Filter{
		Fields: fields,
		Values: values,
		Op:     OpIntersects,
	}
}

func StartsWith(fields []string, value any) Filter {
	return Filter{
		Fields: fields,
		Values: []any{value},
		Op:     OpContains,
	}
}

func Contains(fields []string, value any) Filter {
	return Filter{
		Fields: fields,
		Values: []any{value},
		Op:     OpContains,
	}
}

func Id(id any) Filter {
	return Filter{
		Fields: []string{"id"},
		Values: []any{id},
		Op:     OpEquals,
	}
}

func Field(fieldName string) []string {
	return []string{fieldName}
}

func Fields(fieldNames ...string) []string {
	return fieldNames
}

func AnyField() []string {
	return nil
}

var dateFormats = []string{
	time.RFC3339,
	time.RFC1123,
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02",
	"2006/01/02 15:04:05",
	"2006/01/02 15:04",
	"2006/01/02",
	"02-Jan-2006 15:04:05",
	"02-Jan-2006 15:04",
	"02-Jan-2006",
	"02/01/2006 15:04:05",
	"02/01/2006 15:04",
	"02/01/2006",
	"01/02/2006 15:04:05",
	"01/02/2006 15:04",
	"01/02/2006",
	"2006-1-2 15:04:05",
	"2006-1-2 15:04",
	"2006-1-2",
	"1/2/2006 15:04:05",
	"1/2/2006 15:04",
	"1/2/2006",
}

func ParseAnyDate(input string) (time.Time, error) {
	var t time.Time
	var err error
	for _, format := range dateFormats {
		t, err = time.Parse(format, input)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("could not parse date: %v", input)
}
