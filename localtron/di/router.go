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
package di

import (
	"net/http"

	"github.com/singulatron/singulatron/localtron/middlewares"

	appendpoints "github.com/singulatron/singulatron/localtron/services/app/endpoints"
	chatendpoints "github.com/singulatron/singulatron/localtron/services/chat/endpoints"
	configendpoints "github.com/singulatron/singulatron/localtron/services/config/endpoints"
	dockerendpoints "github.com/singulatron/singulatron/localtron/services/docker/endpoints"
	downloadendpoints "github.com/singulatron/singulatron/localtron/services/download/endpoints"
	firehoseendpoints "github.com/singulatron/singulatron/localtron/services/firehose/endpoints"
	genericendpoints "github.com/singulatron/singulatron/localtron/services/generic/endpoints"
	modelendpoints "github.com/singulatron/singulatron/localtron/services/model/endpoints"
	nodeendpoints "github.com/singulatron/singulatron/localtron/services/node/endpoints"
	promptendpoints "github.com/singulatron/singulatron/localtron/services/prompt/endpoints"
	userendpoints "github.com/singulatron/singulatron/localtron/services/user/endpoints"
)

func HttpHandler(universe *Universe) http.Handler {
	mws := []middlewares.Middleware{
		middlewares.ThrottledLogger,
		middlewares.Recover,
		middlewares.CORS,
		middlewares.GzipDecodeMiddleware,
	}
	appl := applicator(mws)

	router := http.NewServeMux()

	router.HandleFunc("/firehose/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		firehoseendpoints.Subscribe(w, r, universe.UserService, universe.FirehoseService)
	}))

	router.HandleFunc("/download/do", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadendpoints.Do(w, r, universe.UserService, universe.DownloadService)
	}))

	router.HandleFunc("/download/pause", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadendpoints.Pause(w, r, universe.UserService, universe.DownloadService)
	}))

	router.HandleFunc("/download/list", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadendpoints.List(w, r, universe.UserService, universe.DownloadService)
	}))

	router.HandleFunc("/docker/info", appl(func(w http.ResponseWriter, r *http.Request) {
		dockerendpoints.Info(w, r, universe.UserService, universe.DockerService)
	}))

	router.HandleFunc("/model/status", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.Status(w, r, universe.UserService, universe.ModelService)
	}))
	router.HandleFunc("/model/get-models", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.GetModels(w, r, universe.UserService, universe.ModelService)
	}))
	router.HandleFunc("/model/start", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.Start(w, r, universe.UserService, universe.ModelService)
	}))
	router.HandleFunc("/model/make-default", appl(func(w http.ResponseWriter, r *http.Request) {
		modelendpoints.MakeDefault(w, r, universe.UserService, universe.ModelService)
	}))

	router.HandleFunc("/config/get", appl(func(w http.ResponseWriter, r *http.Request) {
		configendpoints.Get(w, r, universe.UserService, universe.ConfigService)
	}))

	router.HandleFunc("/app/log/disable", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.DisableLogging(w, r, universe.AppService)
	}))

	router.HandleFunc("/app/log/enable", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.EnableLogging(w, r, universe.AppService)
	}))

	router.HandleFunc("/app/log/status", appl(func(w http.ResponseWriter, r *http.Request) {
		appendpoints.LoggingStatus(w, r, universe.AppService)
	}))

	router.HandleFunc("/chat/message/add", appl(func(w http.ResponseWriter, r *http.Request) {
		chatendpoints.AddMessage(w, r, universe.UserService, universe.ChatService)
	}))

	router.HandleFunc("/chat/message/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		chatendpoints.DeleteMessage(w, r, universe.UserService, universe.ChatService)
	}))

	router.HandleFunc("/chat/messages", appl(func(w http.ResponseWriter, r *http.Request) {
		chatendpoints.GetMessages(w, r, universe.UserService, universe.ChatService)
	}))

	router.HandleFunc("/chat/thread/add", appl(func(w http.ResponseWriter, r *http.Request) {
		chatendpoints.AddThread(w, r, universe.UserService, universe.ChatService)
	}))

	router.HandleFunc("/chat/thread/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		chatendpoints.DeleteThread(w, r, universe.UserService, universe.ChatService)
	}))

	router.HandleFunc("/chat/threads", appl(func(w http.ResponseWriter, r *http.Request) {
		chatendpoints.GetThreads(w, r, universe.UserService, universe.ChatService)
	}))

	router.HandleFunc("/chat/thread", appl(func(w http.ResponseWriter, r *http.Request) {
		chatendpoints.GetThread(w, r, universe.UserService, universe.ChatService)
	}))

	router.HandleFunc("/chat/thread/update", appl(func(w http.ResponseWriter, r *http.Request) {
		chatendpoints.UpdateThread(w, r, universe.UserService, universe.ChatService)
	}))

	router.HandleFunc("/prompt/add", appl(func(w http.ResponseWriter, r *http.Request) {
		promptendpoints.Add(w, r, universe.UserService, universe.PromptService)
	}))

	router.HandleFunc("/prompt/remove", appl(func(w http.ResponseWriter, r *http.Request) {
		promptendpoints.Remove(w, r, universe.UserService, universe.PromptService)
	}))

	router.HandleFunc("/prompt/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		promptendpoints.Subscribe(w, r, universe.UserService, universe.PromptService)
	}))

	router.HandleFunc("/prompt/list", appl(func(w http.ResponseWriter, r *http.Request) {
		promptendpoints.List(w, r, universe.UserService, universe.PromptService)
	}))

	router.HandleFunc("/user/login", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.Login(w, r, universe.UserService)
	}))
	router.HandleFunc("/user/read-user-by-token", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.ReadUserByToken(w, r, universe.UserService)
	}))
	router.HandleFunc("/user/get-users", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.GetUsers(w, r, universe.UserService)
	}))
	router.HandleFunc("/user/save-profile", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.SaveProfile(w, r, universe.UserService)
	}))
	router.HandleFunc("/user/change-password", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.ChangePassword(w, r, universe.UserService)
	}))
	router.HandleFunc("/user/change-password-admin", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.ChangePasswordAdmin(w, r, universe.UserService)
	}))
	router.HandleFunc("/user/create-user", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.CreateUser(w, r, universe.UserService)
	}))
	router.HandleFunc("/user/delete-user", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.DeleteUser(w, r, universe.UserService)
	}))
	router.HandleFunc("/user/get-roles", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.GetRoles(w, r, universe.UserService)
	}))
	router.HandleFunc("/user/delete-role", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.DeleteRole(w, r, universe.UserService)
	}))
	router.HandleFunc("/user/get-permissions", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.GetPermissions(w, r, universe.UserService)
	}))
	router.HandleFunc("/user/set-role-permissions", appl(func(w http.ResponseWriter, r *http.Request) {
		userendpoints.SetRolePermissions(w, r, universe.UserService)
	}))

	router.HandleFunc("/generic/create", appl(func(w http.ResponseWriter, r *http.Request) {
		genericendpoints.Create(w, r, universe.UserService, universe.GenericService)
	}))
	router.HandleFunc("/generic/update", appl(func(w http.ResponseWriter, r *http.Request) {
		genericendpoints.Update(w, r, universe.UserService, universe.GenericService)
	}))
	router.HandleFunc("/generic/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		genericendpoints.Delete(w, r, universe.UserService, universe.GenericService)
	}))
	router.HandleFunc("/generic/find", appl(func(w http.ResponseWriter, r *http.Request) {
		genericendpoints.Find(w, r, universe.UserService, universe.GenericService)
	}))
	router.HandleFunc("/generic/upsert", appl(func(w http.ResponseWriter, r *http.Request) {
		genericendpoints.Upsert(w, r, universe.UserService, universe.GenericService)
	}))

	router.HandleFunc("/node/list", appl(func(w http.ResponseWriter, r *http.Request) {
		nodeendpoints.List(w, r, universe.UserService, universe.NodeService)
	}))

	return router
}

func applicator(mws []middlewares.Middleware) func(http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		for _, middleware := range mws {
			h = middleware(h)
		}

		return h
	}
}
