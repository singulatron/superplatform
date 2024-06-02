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
	"log"
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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configService, err := configservice.NewConfigService()
	if err != nil {
		log.Fatal(err)
	}
	configService.ConfigDirectory = path.Join(homeDir, singulatronFolder)
	if os.Getenv("SINGULATRON_CONFIG_PATH") != "" {
		configService.ConfigDirectory = os.Getenv("SINGULATRON_CONFIG_PATH")
	}
	err = configService.Start()
	if err != nil {
		log.Fatal(err)
	}

	singulatronFolder := path.Join(homeDir, singulatronFolder)
	err = os.MkdirAll(singulatronFolder, 0755)
	if err != nil {
		log.Fatal(err)
	}

	downloadFolder := path.Join(singulatronFolder, "downloads")
	err = os.MkdirAll(downloadFolder, 0755)
	if err != nil {
		log.Fatal(err)
	}

	firehoseService, err := firehoseservice.NewFirehoseService()
	if err != nil {
		log.Fatal(err)
	}

	downloadService, err := downloadservice.NewDownloadService(firehoseService)
	if err != nil {
		log.Fatal(err)
	}

	downloadService.DefaultFolder = downloadFolder
	downloadService.StateFilePath = path.Join(singulatronFolder, "downloads.json")
	err = downloadService.Start()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	router.HandleFunc("/docker/info", appl(func(w http.ResponseWriter, r *http.Request) {
		dockerendpoints.Info(w, r, dockerService)
	}))

	modelService, err := modelservice.NewModelService(downloadService, configService, dockerService)
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/model/status", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.Status(w, r, modelService)
	}))
	router.HandleFunc("/model/start", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.Start(w, r, modelService)
	}))

	router.HandleFunc("/config/get", appl(func(w http.ResponseWriter, r *http.Request) {
		configendpoints.Get(w, r, configService)
	}))

	appService, err := appservice.NewAppService(configService, firehoseService)
	if err != nil {
		log.Fatal(err)
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
	log.Fatal(http.ListenAndServe(":58231", router))
}

func applicator(mws []middlewares.Middleware) func(http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		for _, middleware := range mws {
			h = middleware(h)
		}

		return h
	}
}
