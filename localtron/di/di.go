package di

import (
	"log/slog"
	"net/http"
	"os"
	"path"

	"github.com/singulatron/singulatron/localtron/clients/llm"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/datastore/localstore"
	"github.com/singulatron/singulatron/localtron/logger"
	"github.com/singulatron/singulatron/localtron/middlewares"
	"github.com/singulatron/singulatron/localtron/router"
	chatservice "github.com/singulatron/singulatron/localtron/services/chat"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	dockerservice "github.com/singulatron/singulatron/localtron/services/docker"
	downloadservice "github.com/singulatron/singulatron/localtron/services/download"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"
	genericservice "github.com/singulatron/singulatron/localtron/services/generic"
	modelservice "github.com/singulatron/singulatron/localtron/services/model"
	nodeservice "github.com/singulatron/singulatron/localtron/services/node"
	promptservice "github.com/singulatron/singulatron/localtron/services/prompt"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

const singulatronFolder = ".singulatron"

type Options struct {
	Test             bool
	LLMClient        llm.ClientI
	Router           *router.Router
	DatastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
}

func BigBang(options Options) (*http.ServeMux, error) {
	var homeDir string
	var err error
	if options.Test {
		homeDir, err = os.MkdirTemp("", "singulatron-")
		if err != nil {
			logger.Error("Homedir creation failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
	} else {
		homeDir, err = os.UserHomeDir()
		if err != nil {
			logger.Error("Homedir creation failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}

	configService, err := configservice.NewConfigService()
	if err != nil {
		logger.Error("Config service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
	configService.EventCallback = func(event firehosetypes.Event) {
		logger.Debug("Received event from config before firehose is set up",
			slog.String("eventName", event.Name),
		)
	}
	configService.ConfigDirectory = path.Join(homeDir, singulatronFolder)
	if os.Getenv("SINGULATRON_CONFIG_PATH") != "" {
		configService.ConfigDirectory = os.Getenv("SINGULATRON_CONFIG_PATH")
	}

	if options.DatastoreFactory == nil {
		localStorePath := path.Join(configService.GetConfigDirectory(), "data")
		err = os.MkdirAll(localStorePath, 0755)
		if err != nil {
			logger.Error("Creating data folder failed", slog.String("error", err.Error()))
			os.Exit(1)
		}

		options.DatastoreFactory = func(tableName string, instance any) (datastore.DataStore, error) {
			return localstore.NewLocalStore(instance, path.Join(localStorePath, tableName))
		}
	}

	if options.Router == nil {
		router, err := router.NewRouter(options.DatastoreFactory)
		if err != nil {
			logger.Error("Creating router failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		options.Router = router
	}

	userService, err := userservice.NewUserService(
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("User service start failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	err = configService.Start()
	if err != nil {
		logger.Error("Config service start failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	firehoseService, err := firehoseservice.NewFirehoseService(options.Router)
	if err != nil {
		logger.Error("Firehose service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

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

	downloadService, err := downloadservice.NewDownloadService(options.Router)
	if err != nil {
		logger.Error("Download service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	downloadService.SetDefaultFolder(downloadFolder)
	downloadService.SetStateFilePath(path.Join(singulatronFolder, "downloads.json"))

	downloadService.Start()
	if err != nil {
		logger.Error("Download service start failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	dockerService, err := dockerservice.NewDockerService(options.Router)
	if err != nil {
		logger.Error("Docker service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	modelService, err := modelservice.NewModelService(
		options.Router,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Model service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if options.Pre.AppService != nil {
		universe.AppService = options.Pre.AppService
	} else {
		appService, err := appservice.NewAppService(
			universe.ConfigService,
			universe.FirehoseService,
			universe.UserService,
		)
		if err != nil {
			logger.Error("App service creation failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		universe.AppService = appService
	}

	if options.Pre.ChatService != nil {
		universe.ChatService = options.Pre.ChatService
	} else {
		chatService, err := chatservice.NewChatService(
			universe.ConfigService,
			universe.FirehoseService,
			universe.UserService,
			options.DatastoreFactory,
		)
		if err != nil {
			logger.Error("Chat service creation failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		universe.ChatService = chatService
	}

	if options.Pre.PromptService != nil {
		universe.PromptService = options.Pre.PromptService
	} else {
		promptService, err := promptservice.NewPromptService(
			universe.ConfigService,
			universe.UserService,
			universe.ModelService,
			universe.ChatService,
			universe.FirehoseService,
			universe.Router,
			options.Pre.LLMClient,
			options.DatastoreFactory,
		)
		if err != nil {
			logger.Error("Prompt service creation failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		universe.PromptService = promptService
	}

	if options.Pre.GenericService != nil {
		universe.GenericService = options.Pre.GenericService
	} else {
		genericService, err := genericservice.NewGenericService(
			universe.ConfigService,
			universe.FirehoseService,
			universe.UserService,
			options.DatastoreFactory,
		)
		if err != nil {
			logger.Error("Generic service creation failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		universe.GenericService = genericService
	}

	if options.Pre.NodeService != nil {
		universe.NodeService = options.Pre.NodeService
	} else {
		nodeService, err := nodeservice.NewNodeService(universe.UserService)
		if err != nil {
			logger.Error("Node service creation failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		universe.NodeService = nodeService
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
		promptendpoints.RemovePrompt(w, r, universe.UserService, universe.PromptService)
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
