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

	Query(condition Condition, conditions ...Condition) QueryBuilder

	BeginTransaction() (DataStore, error)
	Commit() error
	Rollback() error
	IsInTransaction() bool

	SetDebug(debug bool)
}

type QueryBuilder interface {
	OrderBy(option OrderBy, options ...OrderBy) QueryBuilder
	Limit(limit int) QueryBuilder
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
	Field string   `json:"field,omitempty"`
	OneOf []string `json:"oneOf,omitempty"`
	Any   bool     `json:"any,omitempty"`
}

type Condition struct {
	Equal      *EqualCondition      `json:"equal,omitempty"`
	All        *AllCondition        `json:"all,omitempty"`
	StartsWith *StartsWithCondition `json:"startsWith,omitempty"`
	Contains   *ContainsCondition   `json:"contains,omitempty"`
}

func (c Condition) FieldIs(fieldName string) bool {
	if c.Equal != nil && c.Equal.Selector != nil && c.Equal.Selector.Field == fieldName {
		return true
	}
	if c.StartsWith != nil && c.StartsWith.Selector != nil && c.StartsWith.Selector.Field == fieldName {
		return true
	}
	if c.Contains != nil && c.Contains.Selector != nil && c.Contains.Selector.Field == fieldName {
		return true
	}

	return false
}

type EqualCondition struct {
	Selector *FieldSelector `json:"selector,omitempty"`
	Value    any            `json:"value,omitempty"`
}

type StartsWithCondition struct {
	Selector *FieldSelector `json:"selector,omitempty"`
	Value    any            `json:"value,omitempty"`
}

type ContainsCondition struct {
	Selector *FieldSelector `json:"selector,omitempty"`
	Value    any            `json:"value,omitempty"`
}

// Query as a type is not used in the DataStore interface but mostly to accept
// a DataStore query through a HTTP API
type Query struct {
	Conditions []Condition `json:"conditions,omitempty"`
	After      []any       `json:"after,omitempty"`
	Limit      int64       `json:"limit,omitempty"`
	OrderBys   []OrderBy   `json:"orderBys,omitempty"`
	// Count true means return the count of the dataset filtered by Conditions
	// without after or limit
	Count bool `json:"count,omitempty"`
}

func (q *Query) HasFieldCondition(fieldName string) bool {
	for _, v := range q.Conditions {
		if v.FieldIs(fieldName) {
			return true
		}
	}

	return false
}

type OrderBy struct {
	// The field by which to order the results
	Field string `json:"field,omitempty"`

	// Indicates whether the sorting should be in descending order.
	Desc bool `json:"desc,omitempty"`

	// When set to true, indicates that the results should be randomized instead of ordered by the Field and Desc criteria
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

type AllCondition struct {
}

func Equal(selector *FieldSelector, value any) Condition {
	return Condition{
		Equal: &EqualCondition{
			Selector: selector,
			Value:    value,
		},
	}
}

func StartsWith(selector *FieldSelector, value any) Condition {
	return Condition{
		StartsWith: &StartsWithCondition{
			Selector: selector,
			Value:    value,
		},
	}
}

func Contains(selector *FieldSelector, value any) Condition {
	return Condition{
		Contains: &ContainsCondition{
			Selector: selector,
			Value:    value,
		},
	}
}

func All() Condition {
	return Condition{
		All: &AllCondition{},
	}
}

func Id(id any) Condition {
	return Condition{
		Equal: &EqualCondition{
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
