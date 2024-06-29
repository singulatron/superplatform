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
package storefactoryservice

import (
	"os"
	"path"

	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/datastore/localstore"
	"github.com/singulatron/singulatron/localtron/datastore/sqlstore"
)

var LocalStorePath = ""

func init() {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	LocalStorePath = path.Join(dir, "data", "test")
}

// GetStore is a standalone function that returns a new instance of a generic datastore.
// This function needs to be standalone due to Go's current limitation on methods with type parameters.
// In Go, methods defined on a concrete type (like StoreFactoryService) cannot have type parameters.
// By defining GetStore as a standalone function, we can utilize type parameters, allowing
// the function to return a datastore with the appropriate type. This approach ensures flexibility
// and type safety when creating instances of datastores.
//
// Unfortunately this means globals must be utilized to configure this package.
func GetStore[T datastore.Row](tableName string) (datastore.DataStore[T], error) {
	db := os.Getenv("SINGULATRON_DB")
	connStr := os.Getenv("SINGULATRON_DB_SQL_CONNECTION_STRING")
	if db == "postgres" {
		debug := os.Getenv("SINGULATRON_DB_DEBUG") == "true"
		return sqlstore.NewSQLStore[T]("postgres", connStr, tableName, debug)
	}
	return localstore.NewLocalStore[T](path.Join(LocalStorePath, tableName)), nil
}
