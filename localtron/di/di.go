package di

import (
	"log/slog"
	"net/http"
	"os"
	"path"
	"sync"

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
	genericservice "github.com/singulatron/singulatron/localtron/services/generic"
	modelservice "github.com/singulatron/singulatron/localtron/services/model"
	nodeservice "github.com/singulatron/singulatron/localtron/services/node"
	promptservice "github.com/singulatron/singulatron/localtron/services/prompt"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

const singulatronFolder = ".singulatron"

type Options struct {
	Url              string
	Test             bool
	LLMClient        llm.ClientI
	Router           *router.Router
	DatastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
}

func BigBang(options *Options) (*http.ServeMux, func() error, error) {
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
	configService.SetDatastoreFactory(options.DatastoreFactory)

	if options.Router == nil {
		router, err := router.NewRouter(options.DatastoreFactory)
		if err != nil {
			logger.Error("Creating router failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		router.SetDefaultAddress(options.Url)
		options.Router = router
	}

	configService.SetRouter(options.Router)

	userService, err := userservice.NewUserService(
		options.Router,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("User service start failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

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

	chatService, err := chatservice.NewChatService(
		options.Router,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Chat service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	promptService, err := promptservice.NewPromptService(
		options.Router,
		options.LLMClient,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Prompt service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	genericService, err := genericservice.NewGenericService(
		options.Router,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Generic service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	nodeService, err := nodeservice.NewNodeService(options.Router)
	if err != nil {
		logger.Error("Node service creation failed", slog.String("error", err.Error()))
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
		firehoseService.Subscribe(w, r)
	}))

	router.HandleFunc("/download/do", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadService.Do(w, r)
	}))

	router.HandleFunc("/download/pause", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadService.Pause(w, r)
	}))

	router.HandleFunc("/download/list", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadService.List(w, r)
	}))

	router.HandleFunc("/docker/info", appl(func(w http.ResponseWriter, r *http.Request) {
		dockerService.Info(w, r)
	}))

	router.HandleFunc("/model/status", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.Status(w, r)
	}))
	router.HandleFunc("/model/get-models", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.GetModels(w, r)
	}))
	router.HandleFunc("/model/start", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.PostStart(w, r)
	}))
	router.HandleFunc("/model/make-default", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.MakeDefault(w, r)
	}))

	router.HandleFunc("/config/get", appl(func(w http.ResponseWriter, r *http.Request) {
		configService.Get(w, r)
	}))

	router.HandleFunc("/chat/message/add", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.AddMessage(w, r)
	}))

	router.HandleFunc("/chat/message/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.DeleteMessage(w, r)
	}))

	router.HandleFunc("/chat/messages", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.GetMessages(w, r)
	}))

	router.HandleFunc("/chat/thread/add", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.AddThread(w, r)
	}))

	router.HandleFunc("/chat/thread/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.DeleteThread(w, r)
	}))

	router.HandleFunc("/chat/threads", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.GetThreads(w, r)
	}))

	router.HandleFunc("/chat/thread", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.GetThread(w, r)
	}))

	router.HandleFunc("/chat/thread/update", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.UpdateThread(w, r)
	}))

	router.HandleFunc("/prompt/add", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.PostAdd(w, r)
	}))

	router.HandleFunc("/prompt/remove", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.RemovePrompt(w, r)
	}))

	router.HandleFunc("/prompt/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.GetSubscribe(w, r)
	}))

	router.HandleFunc("/prompt/list", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.GetPrompts(w, r)
	}))

	router.HandleFunc("/user/login", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.Login(w, r)
	}))
	router.HandleFunc("/user/read-user-by-token", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ReadUserByToken(w, r)
	}))
	router.HandleFunc("/user/get-users", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.GetUsers(w, r)
	}))
	router.HandleFunc("/user/save-profile", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.SaveProfile(w, r)
	}))
	router.HandleFunc("/user/change-password", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ChangePassword(w, r)
	}))
	router.HandleFunc("/user/change-password-admin", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ChangePasswordAdmin(w, r)
	}))
	router.HandleFunc("/user/create-user", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.CreateUser(w, r)
	}))
	router.HandleFunc("/user/delete-user", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.DeleteUser(w, r)
	}))
	router.HandleFunc("/user/get-roles", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.GetRoles(w, r)
	}))
	router.HandleFunc("/user/delete-role", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.DeleteRole(w, r)
	}))
	router.HandleFunc("/user/get-permissions", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.GetPermissions(w, r)
	}))
	router.HandleFunc("/user/set-role-permissions", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.SetRolePermissions(w, r)
	}))
	router.HandleFunc("/user/upsert-permission", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.UpsertPermission(w, r)
	}))
	router.HandleFunc("/user/register", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.Register(w, r)
	}))
	router.HandleFunc("/user/add-permission-to-role", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.AddPermissionToRole(w, r)
	}))

	router.HandleFunc("/generic/create", appl(func(w http.ResponseWriter, r *http.Request) {
		genericService.Create(w, r)
	}))
	router.HandleFunc("/generic/update", appl(func(w http.ResponseWriter, r *http.Request) {
		genericService.Update(w, r)
	}))
	router.HandleFunc("/generic/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		genericService.Delete(w, r)
	}))
	router.HandleFunc("/generic/find", appl(func(w http.ResponseWriter, r *http.Request) {
		genericService.Find(w, r)
	}))
	router.HandleFunc("/generic/upsert", appl(func(w http.ResponseWriter, r *http.Request) {
		genericService.Upsert(w, r)
	}))

	router.HandleFunc("/node/list", appl(func(w http.ResponseWriter, r *http.Request) {
		nodeService.List(w, r)
	}))

	return router, func() error {
		err = configService.Start()
		if err != nil {
			return err
		}
		err = downloadService.Start()
		if err != nil {
			return err
		}
		err = firehoseService.Start()
		if err != nil {
			return err
		}
		err = dockerService.Start()
		if err != nil {
			return err
		}
		err = modelService.Start()
		if err != nil {
			return err
		}
		err = chatService.Start()
		if err != nil {
			return err
		}

		return nil
	}, nil
}

func applicator(mws []middlewares.Middleware) func(http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		for _, middleware := range mws {
			h = middleware(h)
		}

		return h
	}
}

type HandlerSwitcher struct {
	mu      sync.RWMutex
	handler http.Handler
}

func (hs *HandlerSwitcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hs.mu.RLock()
	defer hs.mu.RUnlock()
	if hs.handler != nil {
		hs.handler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (hs *HandlerSwitcher) UpdateHandler(handler http.Handler) {
	hs.mu.Lock()
	defer hs.mu.Unlock()
	hs.handler = handler
}
