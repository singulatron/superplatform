/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/singulatron/singulatron/localtron/internal/di"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/datastore/sqlstore"
	"github.com/singulatron/singulatron/sdk/go/logger"
	"github.com/singulatron/singulatron/sdk/go/router"

	_ "github.com/singulatron/singulatron/localtron/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

var port = router.GetDefaultPort()

// @title           Singulatron
// @version         0.2
// @description     AI management and development platform.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://superplatform.ai/
// @contact.email  sales@singulatron.com

// @license.name  AGPL v3.0
// @license.url   https://www.gnu.org/licenses/agpl-3.0.html

// @host      localhost:58231
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and token acquired from the User sService Login endpoint.

// @externalDocs.description  Singulatron API
// @externalDocs.url          https://superplatform.ai/docs/category/singulatron-api
func main() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Panic in main",
				slog.String("error", fmt.Sprintf("%v", r)),
				slog.String("trace", string(debug.Stack())),
			)
			os.Exit(1)
		}
	}()

	options := &di.Options{
		Test: false,
	}
	db := os.Getenv("SINGULATRON_DB")
	if db != "" {
		options.DatastoreFactory = func(tableName string, instance any) (datastore.DataStore, error) {
			return sqlstore.NewSQLStore(
				instance,
				os.Getenv("SINGULATRON_DB_DRIVER"),
				os.Getenv("SINGULATRON_DB_STRING"),
				tableName,
				false,
			)
		}
	}

	router, starter, err := di.BigBang(options)
	if err != nil {
		logger.Error("Cannot make universe", slog.Any("error", err))
		os.Exit(1)
	}

	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	srv := &http.Server{
		Handler: router,
	}

	logger.Info("Server started", slog.String("port", port))
	go func() {
		time.Sleep(5 * time.Millisecond)
		err = starter()
		if err != nil {
			logger.Error("Cannot start universe", slog.Any("error", err))
			os.Exit(1)
		}
	}()
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), srv.Handler)
	if err != nil {
		logger.Error("HTTP listen failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
