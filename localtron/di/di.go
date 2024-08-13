package di

import (
	"log/slog"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/gorilla/mux"
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
	HomeDir          string
}

func BigBang(options *Options) (*mux.Router, func() error, error) {
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
	options.HomeDir = homeDir

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
		if options.Url != "" {
			router.SetDefaultAddress(options.Url)
		}
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

	firehoseService, err := firehoseservice.NewFirehoseService(options.Router, options.DatastoreFactory)
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

	downloadService, err := downloadservice.NewDownloadService(options.Router, options.DatastoreFactory)
	if err != nil {
		logger.Error("Download service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	downloadService.SetDefaultFolder(downloadFolder)
	downloadService.SetStateFilePath(path.Join(singulatronFolder, "downloads.json"))

	dockerService, err := dockerservice.NewDockerService(options.Router, options.DatastoreFactory)
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

	router := mux.NewRouter().SkipClean(true).UseEncodedPath()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	router.HandleFunc("/firehose-svc/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		firehoseService.Subscribe(w, r)
	})).Methods("OPTIONS", "GET")
	router.HandleFunc("/firehose-svc/publish", appl(func(w http.ResponseWriter, r *http.Request) {
		firehoseService.Publish(w, r)
	})).Methods("OPTIONS", "POST")

	router.HandleFunc("/download-svc/download", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadService.Do(w, r)
	})).Methods("OPTIONS", "PUT")

	router.HandleFunc("/download-svc/download/{downloadId}/pause", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadService.Pause(w, r)
	})).Methods("OPTIONS", "PUT")
	router.HandleFunc("/download-svc/download/{downloadId}", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadService.Get(w, r)
	})).Methods("OPTIONS", "GET")

	router.HandleFunc("/download-svc/downloads", appl(func(w http.ResponseWriter, r *http.Request) {
		downloadService.List(w, r)
	})).Methods("OPTIONS", "POST")

	router.HandleFunc("/docker-svc/info", appl(func(w http.ResponseWriter, r *http.Request) {
		dockerService.Info(w, r)
	})).Methods("OPTIONS", "GET")
	router.HandleFunc("/docker-svc/host", appl(func(w http.ResponseWriter, r *http.Request) {
		dockerService.Host(w, r)
	})).Methods("OPTIONS", "GET")
	router.HandleFunc("/docker-svc/container", appl(func(w http.ResponseWriter, r *http.Request) {
		dockerService.LaunchContainer(w, r)
	})).Methods("OPTIONS", "PUT")
	router.HandleFunc("/docker-svc/container/{hash}/is-running", appl(func(w http.ResponseWriter, r *http.Request) {
		dockerService.HashIsRunning(w, r)
	})).Methods("OPTIONS", "GET")
	router.HandleFunc("/docker-svc/container/{hash}/summary/{numberOfLines}", appl(func(w http.ResponseWriter, r *http.Request) {
		dockerService.Summary(w, r)
	})).Methods("OPTIONS", "GET")

	router.HandleFunc("/model-svc/default-model/status", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.DefaultStatus(w, r)
	})).Methods("OPTIONS", "GET")
	router.HandleFunc("/model-svc/model/{modelId}/status", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.Status(w, r)
	})).Methods("OPTIONS", "GET")
	router.HandleFunc("/model-svc/models", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.List(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/model-svc/model/{modelId}", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.Get(w, r)
	})).Methods("OPTIONS", "GET")
	router.HandleFunc("/model-svc/default-model/start", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.StartDefault(w, r)
	})).Methods("OPTIONS", "PUT")
	router.HandleFunc("/model-svc/model/{modelId}/start", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.StartSpecific(w, r)
	})).Methods("OPTIONS", "PUT")
	router.HandleFunc("/model-svc/model/{modelId}/make-default", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.MakeDefault(w, r)
	})).Methods("OPTIONS", "PUT")

	router.HandleFunc("/config-svc/config", appl(func(w http.ResponseWriter, r *http.Request) {
		configService.Get(w, r)
	})).Methods("OPTIONS", "GET")
	router.HandleFunc("/config-svc/config", appl(func(w http.ResponseWriter, r *http.Request) {
		configService.Save(w, r)
	})).Methods("OPTIONS", "PUT")

	router.HandleFunc("/chat-svc/thread/{threadId}/message", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.AddMessage(w, r)
	})).Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/message/{messageId}", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.DeleteMessage(w, r)
	})).Methods("OPTIONS", "DELETE")

	router.HandleFunc("/chat-svc/thread/{threadId}/messages", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.GetMessages(w, r)
	})).Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/thread", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.AddThread(w, r)
	})).Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/thread/{threadId}", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.DeleteThread(w, r)
	})).Methods("OPTIONS", "DELETE")

	router.HandleFunc("/chat-svc/threads", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.GetThreads(w, r)
	})).Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/thread/{threadId}", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.GetThread(w, r)
	})).Methods("OPTIONS", "GET")

	router.HandleFunc("/chat-svc/thread/{threadId}", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.UpdateThread(w, r)
	})).Methods("OPTIONS", "PUT")

	router.HandleFunc("/prompt-svc/prompt", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.Add(w, r)
	})).Methods("OPTIONS", "POST")

	router.HandleFunc("/prompt-svc'/prompt/{promptId}", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.RemovePrompt(w, r)
	})).Methods("OPTIONS", "DELETE")

	router.HandleFunc("/prompt-svc/{threadId}/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.GetSubscribe(w, r)
	})).Methods("OPTIONS", "GET")

	router.HandleFunc("/prompt-svc/prompts", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.GetPrompts(w, r)
	})).Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/login", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.Login(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/user/by-token", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ReadUserByToken(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/users", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.GetUsers(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/user/{userId}", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.SaveProfile(w, r)
	})).Methods("OPTIONS", "PUT")
	router.HandleFunc("/user-svc/change-password", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ChangePassword(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/change-password-admin", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ChangePasswordAdmin(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/organization", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.CreateOrganization(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/user", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.CreateUser(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/user/{userId}", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.DeleteUser(w, r)
	})).Methods("OPTIONS", "DELETE")
	router.HandleFunc("/user-svc/roles", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.GetRoles(w, r)
	})).Methods("OPTIONS", "GET")
	router.HandleFunc("/user-svc/role/{roleId}", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.DeleteRole(w, r)
	})).Methods("OPTIONS", "DELETE")
	router.HandleFunc("/user-svc/role", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.CreateRole(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/permission/{permissionId}/is-authorized", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.IsAuthorized(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/role/{roleId}/permissions", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.GetPermissions(w, r)
	})).Methods("OPTIONS", "GET")
	router.HandleFunc("/user-svc/role/{roleId}/permissions", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.SetRolePermissions(w, r)
	})).Methods("OPTIONS", "PUT")
	router.HandleFunc("/user-svc/permission/{permissionId}", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.UpsertPermission(w, r)
	})).Methods("OPTIONS", "PUT")
	router.HandleFunc("/user-svc/register", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.Register(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/role/{roleId}/permission/{permissionId}", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.AddPermissionToRole(w, r)
	})).Methods("OPTIONS", "PUT")
	router.HandleFunc("/user-svc/public-key", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.GetPublicKey(w, r)
	})).Methods("OPTIONS", "GET")

	router.HandleFunc("/generic-svc/object", appl(func(w http.ResponseWriter, r *http.Request) {
		genericService.Create(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/generic-svc/objects/update", appl(func(w http.ResponseWriter, r *http.Request) {
		genericService.Update(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/generic-svc/objects/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		genericService.Delete(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/generic-svc/objects", appl(func(w http.ResponseWriter, r *http.Request) {
		genericService.Find(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/generic-svc/object/{objectId}", appl(func(w http.ResponseWriter, r *http.Request) {
		genericService.Upsert(w, r)
	})).Methods("OPTIONS", "PUT")

	router.HandleFunc("/node-svc/nodes", appl(func(w http.ResponseWriter, r *http.Request) {
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
		err = promptService.Start()
		if err != nil {
			return err
		}
		err = genericService.Start()
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
