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
	"path"

	"github.com/singulatron/singulatron/localtron/middlewares"

	dockerservice "github.com/singulatron/singulatron/localtron/services/docker"
	dockerendpoints "github.com/singulatron/singulatron/localtron/services/docker/endpoints"

	modelservice "github.com/singulatron/singulatron/localtron/services/model"
	modelendpoints "github.com/singulatron/singulatron/localtron/services/model/endpoints"

	downloadservice "github.com/singulatron/singulatron/localtron/services/download"
	downloadendpoints "github.com/singulatron/singulatron/localtron/services/download/endpoints"

	configservice "github.com/singulatron/singulatron/localtron/services/config"
	configendpoints "github.com/singulatron/singulatron/localtron/services/config/endpoints"

	appservice "github.com/singulatron/singulatron/localtron/services/app"
	appendpoints "github.com/singulatron/singulatron/localtron/services/app/endpoints"

	promptservice "github.com/singulatron/singulatron/localtron/services/prompt"
	promptendpoints "github.com/singulatron/singulatron/localtron/services/prompt/endpoints"

	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	firehoseendpoints "github.com/singulatron/singulatron/localtron/services/firehose/endpoints"

	"github.com/singulatron/singulatron/localtron/lib"
)

const singulatronFolder = ".singulatron"
const port = "58231"

func main() {
	defer func() {
		if r := recover(); r != nil {
			lib.Logger.Error("Panic in main", slog.String("trace", fmt.Sprintf("%v", r)))
			os.Exit(1)
		}
	}()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		lib.Logger.Error("Homedir creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	firehoseService, err := firehoseservice.NewFirehoseService()
	if err != nil {
		lib.Logger.Error("Firehose service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	configService, err := configservice.NewConfigService(firehoseService)
	if err != nil {
		lib.Logger.Error("Config service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
	configService.ConfigDirectory = path.Join(homeDir, singulatronFolder)
	if os.Getenv("SINGULATRON_CONFIG_PATH") != "" {
		configService.ConfigDirectory = os.Getenv("SINGULATRON_CONFIG_PATH")
	}
	err = configService.Start()
	if err != nil {
		lib.Logger.Error("Config service start failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	singulatronFolder := path.Join(homeDir, singulatronFolder)
	err = os.MkdirAll(singulatronFolder, 0755)
	if err != nil {
		lib.Logger.Error("Config folder creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	downloadFolder := path.Join(singulatronFolder, "downloads")
	err = os.MkdirAll(downloadFolder, 0755)
	if err != nil {
		lib.Logger.Error("Downloads folder creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	downloadService, err := downloadservice.NewDownloadService(firehoseService)
	if err != nil {
		lib.Logger.Error("Download service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	downloadService.DefaultFolder = downloadFolder
	downloadService.StateFilePath = path.Join(singulatronFolder, "downloads.json")
	err = downloadService.Start()
	if err != nil {
		lib.Logger.Error("Download service start failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	mws := []middlewares.Middleware{
		middlewares.ThrottledLogger,
		middlewares.Recover,
		middlewares.CORS,
		middlewares.GzipDecodeMiddleware,
	}
	appl := applicator(mws)

	router := http.NewServeMux()

	router.HandleFunc("/firehose/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		firehoseendpoints.Subscribe(w, r, firehoseService)
	}))

	router.HandleFunc("/download/do", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadendpoints.Do(w, r, downloadService)
	}))

	router.HandleFunc("/download/pause", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadendpoints.Pause(w, r, downloadService)
	}))

	router.HandleFunc("/download/list", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadendpoints.List(w, r, downloadService)
	}))

	dockerService, err := dockerservice.NewDockerService(downloadService)
	if err != nil {
		lib.Logger.Error("Docker service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	router.HandleFunc("/docker/info", appl(func(w http.ResponseWriter, r *http.Request) {
		dockerendpoints.Info(w, r, dockerService)
	}))

	modelService, err := modelservice.NewModelService(downloadService, configService, dockerService)
	if err != nil {
		lib.Logger.Error("Model service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	router.HandleFunc("/model/status", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.Status(w, r, modelService)
	}))
	router.HandleFunc("/model/start", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.Start(w, r, modelService)
	}))
	router.HandleFunc("/model/make-default", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.MakeDefault(w, r, modelService)
	}))

	router.HandleFunc("/config/get", appl(func(w http.ResponseWriter, r *http.Request) {
		configendpoints.Get(w, r, configService)
	}))

	appService, err := appservice.NewAppService(configService, firehoseService)
	if err != nil {
		lib.Logger.Error("App service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	router.HandleFunc("/app/log", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.Log(w, r, appService)
	}))

	router.HandleFunc("/app/log/disable", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.DisableLogging(w, r, appService)
	}))

	router.HandleFunc("/app/log/enable", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.EnableLogging(w, r, appService)
	}))

	router.HandleFunc("/app/log/status", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.LoggingStatus(w, r, appService)
	}))

	router.HandleFunc("/chat/message/add", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.AddChatMessage(w, r, appService)
	}))

	router.HandleFunc("/chat/message/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.DeleteChatMessage(w, r, appService)
	}))

	router.HandleFunc("/chat/messages", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.GetChatMessages(w, r, appService)
	}))

	router.HandleFunc("/chat/thread/add", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.AddChatThread(w, r, appService)
	}))

	router.HandleFunc("/chat/thread/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.DeleteChatThread(w, r, appService)
	}))

	router.HandleFunc("/chat/threads", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.GetChatThreads(w, r, appService)
	}))

	router.HandleFunc("/chat/thread", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.GetChatThread(w, r, appService)
	}))

	router.HandleFunc("/chat/thread/update", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.UpdateChatThread(w, r, appService)
	}))

	promptService := promptservice.NewPromptService(modelService, appService, firehoseService)

	router.HandleFunc("/prompt/add", appl(func(w http.ResponseWriter, r *http.Request) {
		promptendpoints.Add(w, r, promptService)
	}))

	router.HandleFunc("/prompt/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		promptendpoints.Subscribe(w, r, promptService)
	}))

	router.HandleFunc("/prompt/list", appl(func(w http.ResponseWriter, r *http.Request) {
		promptendpoints.List(w, r, promptService)
	}))

	lib.Logger.Info("Server started", slog.String("port", port))
	err = http.ListenAndServe(":58231", router)
	if err != nil {
		lib.Logger.Error("HTTP listen failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func applicator(mws []middlewares.Middleware) func(http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		for _, middleware := range mws {
			h = middleware(h)
		}

		return h
	}
}
