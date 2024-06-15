package datastore

type DataStore[T any] interface {
	Create(obj T) error
	Read(id string) (T, bool, error)
	Update(id string, obj T) error
	Delete(id string) error
	Query() QueryBuilder[T]
	BatchCreate(objs []T) error
	BatchUpdate(ids []string, objs []T) error
	BatchDelete(ids []string) error

	BeginTransaction() (DataStore[T], error)
	Commit() error
	Rollback() error
	IsInTransaction() bool
}

type QueryBuilder[T any] interface {
	Where(field string, value interface{}) QueryBuilder[T]
	AndWhere(field string, value interface{}) QueryBuilder[T]
	OrWhere(field string, value interface{}) QueryBuilder[T]
	OrderBy(field string, desc bool) QueryBuilder[T]
	Limit(limit int) QueryBuilder[T]
	Offset(offset int) QueryBuilder[T]
	Select(fields ...string) QueryBuilder[T]
	Find() ([]T, error)
	Count() (int64, error)
	UpdateFields(fields map[string]interface{}) error
	Delete() error
}
