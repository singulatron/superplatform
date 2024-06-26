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
	"runtime/debug"

	"github.com/singulatron/singulatron/localtron/logger"
	"github.com/singulatron/singulatron/localtron/middlewares"

	dockerservice "github.com/singulatron/singulatron/localtron/services/docker"
	dockerendpoints "github.com/singulatron/singulatron/localtron/services/docker/endpoints"
	storefactoryservice "github.com/singulatron/singulatron/localtron/services/store_factory"

	userservice "github.com/singulatron/singulatron/localtron/services/user"
	userendpoints "github.com/singulatron/singulatron/localtron/services/user/endpoints"

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
	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"

	genericservice "github.com/singulatron/singulatron/localtron/services/generic"
	genericendpoints "github.com/singulatron/singulatron/localtron/services/generic/endpoints"
)

const singulatronFolder = ".singulatron"
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

	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Error("Homedir creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	configService, err := configservice.NewConfigService()
	if err != nil {
		logger.Error("Config service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
	configService.EventCallback = func(event firehosetypes.Event) {
		logger.Debug("Received event from config before firehose is set up",
			slog.String("eventName", event.Name()),
		)
	}
	configService.ConfigDirectory = path.Join(homeDir, singulatronFolder)
	if os.Getenv("SINGULATRON_CONFIG_PATH") != "" {
		configService.ConfigDirectory = os.Getenv("SINGULATRON_CONFIG_PATH")
	}
	storefactoryservice.LocalStorePath = path.Join(configService.ConfigDirectory, "data")
	err = os.MkdirAll(storefactoryservice.LocalStorePath, 0755)
	if err != nil {
		logger.Error("Creating data folder failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	userService, err := userservice.NewUserService(
		configService,
	)
	if err != nil {
		logger.Error("User service start failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
	// hacks to avoid import cycles
	configService.UpsertPermission = userService.UpsertPermission
	configService.AddPermissionToRole = userService.AddPermissionToRole

	err = configService.Start()
	if err != nil {
		logger.Error("Config service start failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	firehoseService, err := firehoseservice.NewFirehoseService(userService)
	if err != nil {
		logger.Error("Firehose service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
	configService.EventCallback = firehoseService.Publish

	singulatronFolder := path.Join(homeDir, singulatronFolder)
	err = os.MkdirAll(singulatronFolder, 0755)
	if err != nil {
		logger.Error("Config folder creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	downloadFolder := path.Join(singulatronFolder, "downloads")
	err = os.MkdirAll(downloadFolder, 0755)
	if err != nil {
		logger.Error("Downloads folder creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	downloadService, err := downloadservice.NewDownloadService(firehoseService, userService)
	if err != nil {
		logger.Error("Download service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	downloadService.DefaultFolder = downloadFolder
	downloadService.StateFilePath = path.Join(singulatronFolder, "downloads.json")
	err = downloadService.Start()
	if err != nil {
		logger.Error("Download service start failed", slog.String("error", err.Error()))
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
		firehoseendpoints.Subscribe(w, r, userService, firehoseService)
	}))

	router.HandleFunc("/download/do", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadendpoints.Do(w, r, userService, downloadService)
	}))

	router.HandleFunc("/download/pause", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadendpoints.Pause(w, r, userService, downloadService)
	}))

	router.HandleFunc("/download/list", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadendpoints.List(w, r, userService, downloadService)
	}))

	dockerService, err := dockerservice.NewDockerService(downloadService, userService, configService)
	if err != nil {
		logger.Error("Docker service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	router.HandleFunc("/docker/info", appl(func(w http.ResponseWriter, r *http.Request) {
		dockerendpoints.Info(w, r, userService, dockerService)
	}))

	modelService, err := modelservice.NewModelService(downloadService, userService, configService, dockerService)
	if err != nil {
		logger.Error("Model service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	router.HandleFunc("/model/status", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.Status(w, r, userService, modelService)
	}))
	router.HandleFunc("/model/get-models", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.GetModels(w, r, userService, modelService)
	}))
	router.HandleFunc("/model/start", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.Start(w, r, userService, modelService)
	}))
	router.HandleFunc("/model/make-default", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.MakeDefault(w, r, userService, modelService)
	}))

	router.HandleFunc("/config/get", appl(func(w http.ResponseWriter, r *http.Request) {
		configendpoints.Get(w, r, userService, configService)
	}))

	appService, err := appservice.NewAppService(
		configService,
		firehoseService,
		userService,
	)
	if err != nil {
		logger.Error("App service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

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
		appendpoints.AddChatMessage(w, r, userService, appService)
	}))

	router.HandleFunc("/chat/message/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.DeleteChatMessage(w, r, userService, appService)
	}))

	router.HandleFunc("/chat/messages", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.GetChatMessages(w, r, userService, appService)
	}))

	router.HandleFunc("/chat/thread/add", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.AddChatThread(w, r, userService, appService)
	}))

	router.HandleFunc("/chat/thread/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.DeleteChatThread(w, r, userService, appService)
	}))

	router.HandleFunc("/chat/threads", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.GetChatThreads(w, r, userService, appService)
	}))

	router.HandleFunc("/chat/thread", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.GetChatThread(w, r, userService, appService)
	}))

	router.HandleFunc("/chat/thread/update", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.UpdateChatThread(w, r, userService, appService)
	}))

	promptService, err := promptservice.NewPromptService(
		configService,
		userService,
		modelService,
		appService,
		firehoseService,
	)
	if err != nil {
		logger.Error("Prompt service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	router.HandleFunc("/prompt/add", appl(func(w http.ResponseWriter, r *http.Request) {
		promptendpoints.Add(w, r, userService, promptService)
	}))

	router.HandleFunc("/prompt/remove", appl(func(w http.ResponseWriter, r *http.Request) {
		promptendpoints.Remove(w, r, userService, promptService)
	}))

	router.HandleFunc("/prompt/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		promptendpoints.Subscribe(w, r, userService, promptService)
	}))

	router.HandleFunc("/prompt/list", appl(func(w http.ResponseWriter, r *http.Request) {
		promptendpoints.List(w, r, userService, promptService)
	}))

	router.HandleFunc("/user/login", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.Login(w, r, userService)
	}))
	router.HandleFunc("/user/read-user-by-token", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.ReadUserByToken(w, r, userService)
	}))
	router.HandleFunc("/user/get-users", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.GetUsers(w, r, userService)
	}))
	router.HandleFunc("/user/save-profile", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.SaveProfile(w, r, userService)
	}))
	router.HandleFunc("/user/change-password", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.ChangePassword(w, r, userService)
	}))
	router.HandleFunc("/user/create-user", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.CreateUser(w, r, userService)
	}))
	router.HandleFunc("/user/delete-user", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.DeleteUser(w, r, userService)
	}))
	router.HandleFunc("/user/get-roles", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.GetRoles(w, r, userService)
	}))
	router.HandleFunc("/user/delete-role", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.DeleteRole(w, r, userService)
	}))
	router.HandleFunc("/user/get-permissions", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.GetPermissions(w, r, userService)
	}))
	router.HandleFunc("/user/set-role-permissions", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.SetRolePermissions(w, r, userService)
	}))

	genericService, err := genericservice.NewGenericService(
		configService,
		firehoseService,
		userService,
	)
	if err != nil {
		logger.Error("Generic service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	router.HandleFunc("/generic/create", appl(func(w http.ResponseWriter, r *http.Request) {
		genericendpoints.Create(w, r, userService, genericService)
	}))
	router.HandleFunc("/generic/update", appl(func(w http.ResponseWriter, r *http.Request) {
		genericendpoints.Update(w, r, userService, genericService)
	}))
	router.HandleFunc("/generic/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		genericendpoints.Delete(w, r, userService, genericService)
	}))
	router.HandleFunc("/generic/find", appl(func(w http.ResponseWriter, r *http.Request) {
		genericendpoints.Find(w, r, userService, genericService)
	}))
	router.HandleFunc("/generic/upsert", appl(func(w http.ResponseWriter, r *http.Request) {
		genericendpoints.Upsert(w, r, userService, genericService)
	}))

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

func applicator(mws []middlewares.Middleware) func(http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		for _, middleware := range mws {
			h = middleware(h)
		}

		return h
	}
}
