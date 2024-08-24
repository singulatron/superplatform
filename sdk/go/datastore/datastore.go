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

	Query(condition Filter, conditions ...Filter) QueryBuilder

	BeginTransaction() (DataStore, error)
	Commit() error
	Rollback() error
	IsInTransaction() bool

	SetDebug(debug bool)
}

type QueryBuilder interface {
	OrderBy(option OrderBy, options ...OrderBy) QueryBuilder
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

type FieldSelector struct {
	// Field matchies a single field
	Field string `json:"field,omitempty"`

	// OneOf matches a number of fields
	OneOf []string `json:"oneOf,omitempty"`

	// Any matches any fields in the object
	Any bool `json:"any,omitempty"`
}

type Filter struct {
	// Equals condition returns objects where value of a field equals (=) to the specified value in the query.
	Equals *EqualsFilter `json:"equal,omitempty"`

	// Contains condition returns all objects where the field(s) values contain a particular string or slice element.
	Contains *ContainsFilter `json:"contains,omitempty"`

	// Intersects condition returns objects where the slice value of a field intersects with the slice value in the query.
	Intersects *IntersectsFilter `json:"intersects,omitempty"`

	// All condition returns all objects.
	All *AllFilter `json:"all,omitempty"`

	// StartsWith condition returns all objects where the field(s) values start with a particular string.
	StartsWith *StartsWithFilter `json:"startsWith,omitempty"`
}

func (c Filter) FieldIs(fieldName string) bool {
	if c.Equals != nil && c.Equals.Selector != nil && c.Equals.Selector.Field == fieldName {
		return true
	}
	if c.StartsWith != nil && c.StartsWith.Selector != nil && c.StartsWith.Selector.Field == fieldName {
		return true
	}
	if c.Contains != nil && c.Contains.Selector != nil && c.Contains.Selector.Field == fieldName {
		return true
	}
	if c.Intersects != nil && c.Intersects.Selector != nil && c.Intersects.Selector.Field == fieldName {
		return true
	}

	return false
}

type EqualsFilter struct {
	// Selector selects one, more or all fields
	Selector *FieldSelector `json:"selector,omitempty"`
	Value    any            `json:"value,omitempty"`
}

type StartsWithFilter struct {
	// Selector selects one, more or all fields
	Selector *FieldSelector `json:"selector,omitempty"`
	Value    any            `json:"value,omitempty"`
}

type ContainsFilter struct {
	// Selector selects one, more or all fields
	Selector *FieldSelector `json:"selector,omitempty"`
	Value    any            `json:"value,omitempty"`
}

type IntersectsFilter struct {
	Selector *FieldSelector `json:"selector,omitempty"`
	Values   []any          `json:"values,omitempty"`
}

// Query as a type is not used in the DataStore interface but mostly to accept
// a DataStore query through a HTTP API
type Query struct {
	// Filters are filtering options of a query. It is advised to use
	// It's advised to use helper functions in your respective client library such as condition constructors (`all`, `equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
	Filters []Filter `json:"conditions,omitempty"`

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

type AllFilter struct {
}

func Equals(selector *FieldSelector, value any) Filter {
	return Filter{
		Equals: &EqualsFilter{
			Selector: selector,
			Value:    value,
		},
	}
}

func Intersects(selector *FieldSelector, values []any) Filter {
	return Filter{
		Intersects: &IntersectsFilter{
			Selector: selector,
			Values:   values,
		},
	}
}

func StartsWith(selector *FieldSelector, value any) Filter {
	return Filter{
		StartsWith: &StartsWithFilter{
			Selector: selector,
			Value:    value,
		},
	}
}

func Contains(selector *FieldSelector, value any) Filter {
	return Filter{
		Contains: &ContainsFilter{
			Selector: selector,
			Value:    value,
		},
	}
}

func All() Filter {
	return Filter{
		All: &AllFilter{},
	}
}

func Id(id any) Filter {
	return Filter{
		Equals: &EqualsFilter{
			Selector: Field("id"),
			Value:    id,
		},
	}
}

func Field(fieldName string) *FieldSelector {
	return &FieldSelector{
		Field: fieldName,
	}
}

func Fields(fieldNames []string) *FieldSelector {
	return &FieldSelector{
		OneOf: fieldNames,
	}
}

func AnyField() *FieldSelector {
	return &FieldSelector{
		Any: true,
	}
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
