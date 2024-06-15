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

type DataStore[T any] interface {
	Create(obj T) error
	CreateMany(objs []T) error

	Query(condition Condition, conditions ...Condition) QueryBuilder[T]

	BeginTransaction() (DataStore[T], error)
	Commit() error
	Rollback() error
	IsInTransaction() bool
}

type QueryBuilder[T any] interface {
	OrderBy(field string, desc bool) QueryBuilder[T]
	Limit(limit int) QueryBuilder[T]
	Offset(offset int) QueryBuilder[T]

	Select(fields ...string) QueryBuilder[T]
	Find() ([]T, error)
	FindOne() (T, bool, error)
	Count() (int64, error)

	Update(obj T) error
	UpdateFields(fields map[string]interface{}) error
	Delete() error
}

type Condition struct {
	Equal *EqualCondition
	All   *AllCondition
}

type EqualCondition struct {
	FieldName string
	Value     any
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
