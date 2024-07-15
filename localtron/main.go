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
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/singulatron/singulatron/localtron/di"
	"github.com/singulatron/singulatron/localtron/logger"
)

const port = "58231"

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

	universe, err := di.BigBang(di.UniverseOptions{
		Test: false,
	})
	if err != nil {
		logger.Error("Cannot make universe", slog.Any("error", err))
		os.Exit(1)
	}

	router := di.HttpHandler(universe)

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
