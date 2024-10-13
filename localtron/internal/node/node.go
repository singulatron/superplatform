/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package node

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/gorilla/mux"
	"github.com/singulatron/singulatron/localtron/internal/di"
	node_types "github.com/singulatron/singulatron/localtron/internal/node/types"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/datastore/sqlstore"
	"github.com/singulatron/singulatron/sdk/go/logger"

	_ "github.com/singulatron/singulatron/localtron/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Start wraps the dependency injection universe creation
// so getting envars happens outside of that.
// The two could probably be merged.
func Start(options node_types.Options) (*mux.Router, func() error, error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Panic in node.Start()",
				slog.String("error", fmt.Sprintf("%v", r)),
				slog.String("trace", string(debug.Stack())),
			)
			os.Exit(1)
		}
	}()

	if options.GpuPlatform == "" {
		options.GpuPlatform = os.Getenv("SINGULATRON_GPU_PLATFORM")
	}
	if options.Address == "" {
		options.Address = os.Getenv("SINGULATRON_ADDRESS")
	}
	if options.Az == "" {
		options.Az = os.Getenv("SINGULATRON_AZ")
	}
	if options.Region == "" {
		options.Region = os.Getenv("SINGULATRON_AZ")
	}
	if options.LLMHost == "" {
		options.LLMHost = os.Getenv("SINGULATRON_LLM_HOST")
	}
	if options.VolumeName == "" {
		options.VolumeName = os.Getenv("SINGULATRON_VOLUME_NAME")
	}
	if options.ConfigPath == "" {
		options.ConfigPath = os.Getenv("SINGULATRON_CONFIG_PATH")
	}
	if options.Db == "" {
		options.Db = os.Getenv("SINGULATRON_DB")
	}
	if options.DbDriver == "" {
		options.DbDriver = os.Getenv("SINGULATRON_DB_DRIVER")
	}
	if options.DbString == "" {
		options.DbString = os.Getenv("SINGULATRON_DB_STRING")
	}

	diopt := &di.Options{
		NodeOptions: options,
		Test:        false,
	}

	var tablePrefix string
	if options.DbPrefix != "" {
		tablePrefix = options.DbPrefix
	}

	if options.Db != "" {
		if options.DbDriver == "" {
			options.DbDriver = "postgres"
		}
		if options.DbString == "" {
			options.DbString = "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
		}

		diopt.DatastoreFactory = func(tableName string, instance any) (datastore.DataStore, error) {
			return sqlstore.NewSQLStore(
				instance,
				options.DbDriver,
				options.DbString,
				tablePrefix+"_"+tableName,
				false,
			)
		}
	}

	router, starter, err := di.BigBang(diopt)
	if err != nil {
		logger.Error("Cannot make universe", slog.Any("error", err))
		os.Exit(1)
	}

	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	return router, starter, err
}
