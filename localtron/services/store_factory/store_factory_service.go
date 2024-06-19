package storefactoryservice

import (
	"path"

	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/datastore/localstore"
)

var LocalStorePath = ""

// GetStore is a standalone function that returns a new instance of a generic datastore.
// This function needs to be standalone due to Go's current limitation on methods with type parameters.
// In Go, methods defined on a concrete type (like StoreFactoryService) cannot have type parameters.
// By defining GetStore as a standalone function, we can utilize type parameters, allowing
// the function to return a datastore with the appropriate type. This approach ensures flexibility
// and type safety when creating instances of datastores.
//
// Unfortunately this means globals must be utilized to configure this package.
func GetStore[T datastore.Row](tableName string) (datastore.DataStore[T], error) {
	return localstore.NewLocalStore[T](path.Join(LocalStorePath, tableName)), nil
}
