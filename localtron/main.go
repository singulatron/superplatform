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

	"github.com/singulatron/singulatron/localtron/di"
	"github.com/singulatron/singulatron/localtron/logger"
	"github.com/singulatron/singulatron/localtron/router"

	_ "github.com/singulatron/singulatron/localtron/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

const port = router.DefaultPort

// @title           Singulatron
// @version         0.2
// @description     Run and develop self-hosted AI apps. Your programmable in-house GPT. The Firebase for the AI age.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://superplatform.ai/
// @contact.email  sales@singulatron.com

// @license.name  AGPL v3.0
// @license.url   https://www.gnu.org/licenses/agpl-3.0.html

// @host      localhost:58231
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

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

	router, err := di.BigBang(di.Options{
		Test: false,
	})
	if err != nil {
		logger.Error("Cannot make universe", slog.Any("error", err))
		os.Exit(1)
	}

	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	srv := &http.Server{
		Handler: router,
	}

	logger.Info("Server started", slog.String("port", port))
	err = http.ListenAndServe(":58231", srv.Handler)
	if err != nil {
		logger.Error("HTTP listen failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
