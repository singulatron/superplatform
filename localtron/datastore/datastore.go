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
package datastore

import "errors"

type Row interface {
	GetId() string
}

var (
	ErrEntryAlreadyExists = errors.New("entry already exists")
)

type DataStore[T Row] interface {
	/*
	 * Create an object.
	 * Returns ErrEntryAlreadyExists if the object already exists.
	 */
	Create(obj T) error
	/* Create many objects
	* Returns ErrEntryAlreadyExists if any of the objects are already in set,
	* and no object will be created.
	 */
	CreateMany(objs []T) error
	/* Create or Update an object */
	Upsert(obj T) error
	/* Create or Update many objects */
	UpsertMany(objs []T) error

	Query(condition Condition, conditions ...Condition) QueryBuilder[T]

	BeginTransaction() (DataStore[T], error)
	Commit() error
	Rollback() error
	IsInTransaction() bool

	SetDebug(debug bool)
}

type QueryBuilder[T Row] interface {
	OrderBy(field string, desc bool) QueryBuilder[T]
	Limit(limit int) QueryBuilder[T]
	After(value ...any) QueryBuilder[T]

	Select(fields ...string) QueryBuilder[T]
	Find() ([]T, error)
	FindOne() (T, bool, error)
	Count() (int64, error)

	// Update by query. Errors if no update happens
	Update(obj T) error
	// Upsert tries to update by query, and if no update appened, calls create.
	Upsert(obj T) error
	UpdateFields(fields map[string]interface{}) error
	Delete() error
}

type Condition struct {
	Equal *EqualCondition `json:"equal,omitempty"`
	All   *AllCondition   `json:"all,omitempty"`
}

type EqualCondition struct {
	FieldName string `json:"fieldName,omitempty"`
	Value     any    `json:"value,omitempty"`
}

type AllCondition struct {
}

func Equal(fieldName string, value any) Condition {
	return Condition{
		Equal: &EqualCondition{
			FieldName: fieldName,
			Value:     value,
		},
	}
}

func All() Condition {
	return Condition{
		All: &AllCondition{},
	}
}

func Id(id string) Condition {
	return Condition{
		Equal: &EqualCondition{
			FieldName: "id",
			Value:     id,
		},
	}
}
