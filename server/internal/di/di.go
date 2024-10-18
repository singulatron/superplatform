package di

import (
	"log/slog"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/gorilla/mux"
	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/clients/llm"
	"github.com/singulatron/superplatform/sdk/go/datastore"
	"github.com/singulatron/superplatform/sdk/go/datastore/localstore"
	"github.com/singulatron/superplatform/sdk/go/lock"
	distlock "github.com/singulatron/superplatform/sdk/go/lock/local"
	"github.com/singulatron/superplatform/sdk/go/logger"
	"github.com/singulatron/superplatform/sdk/go/middlewares"
	"github.com/singulatron/superplatform/sdk/go/router"
	node_types "github.com/singulatron/superplatform/server/internal/node/types"
	chatservice "github.com/singulatron/superplatform/server/internal/services/chat"
	configservice "github.com/singulatron/superplatform/server/internal/services/config"
	deployservice "github.com/singulatron/superplatform/server/internal/services/deploy"
	dockerservice "github.com/singulatron/superplatform/server/internal/services/docker"
	downloadservice "github.com/singulatron/superplatform/server/internal/services/download"
	dynamicservice "github.com/singulatron/superplatform/server/internal/services/dynamic"
	firehoseservice "github.com/singulatron/superplatform/server/internal/services/firehose"
	modelservice "github.com/singulatron/superplatform/server/internal/services/model"
	policyservice "github.com/singulatron/superplatform/server/internal/services/policy"
	promptservice "github.com/singulatron/superplatform/server/internal/services/prompt"
	registryservice "github.com/singulatron/superplatform/server/internal/services/registry"
	userservice "github.com/singulatron/superplatform/server/internal/services/user"
)

const singulatronFolder = ".singulatron"

type Options struct {
	// NodeOptions contains settings coming from envars
	NodeOptions node_types.Options

	// Url that will be passed down to the router when calling
	// the Singulatron daemon from itself.
	// (Inter-service calls go through the network.)
	Url string

	// Test mode if true will cause the localstore to
	// save data into random temporary folders.
	Test bool

	Lock lock.DistributedLock

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

	if options.Lock == nil {
		options.Lock = distlock.NewLocalDistributedLock()
	}

	configService, err := configservice.NewConfigService(options.Lock)
	if err != nil {
		logger.Error("Config service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	singulatronFolder := path.Join(homeDir, singulatronFolder)
	if options.NodeOptions.ConfigPath != "" {
		singulatronFolder = options.NodeOptions.ConfigPath
	}

	configService.ConfigDirectory = singulatronFolder

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

	firehoseService, err := firehoseservice.NewFirehoseService(
		options.Router,
		options.Lock,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Firehose service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

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

	downloadService, err := downloadservice.NewDownloadService(
		options.Router,
		options.Lock,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Download service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	downloadService.SetDefaultFolder(downloadFolder)
	downloadService.SetStateFilePath(path.Join(singulatronFolder, "downloads.json"))

	dockerService, err := dockerservice.NewDockerService(
		options.NodeOptions.VolumeName,
		options.Router,
		options.Lock,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Docker service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	modelService, err := modelservice.NewModelService(
		options.NodeOptions.GpuPlatform,
		options.NodeOptions.LLMHost,
		options.Router,
		options.Lock,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Model service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	chatService, err := chatservice.NewChatService(
		options.Router,
		options.Lock,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Chat service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	promptService, err := promptservice.NewPromptService(
		options.Router,
		options.LLMClient,
		options.Lock,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Prompt service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	dynamicService, err := dynamicservice.NewDynamicService(
		options.Router,
		options.Lock,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Generic service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	policyService, err := policyservice.NewPolicyService(
		options.Router,
		options.Lock,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Policy service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	registryService, err := registryservice.NewRegistryService(
		options.NodeOptions.Address,
		options.NodeOptions.Az,
		options.NodeOptions.Region,
		options.Router,
		options.Lock,
		options.DatastoreFactory,
	)
	if err != nil {
		logger.Error("Node service creation failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	clientFactory := sdk.NewApiClientFactory(router.SelfAddress())

	deployService, err := deployservice.NewDeployService(
		clientFactory,
		options.Lock,
		options.DatastoreFactory,
	)
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

	router.HandleFunc("/firehose-svc/events/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		firehoseService.Subscribe(w, r)
	})).Methods("OPTIONS", "GET")
	router.HandleFunc("/firehose-svc/event", appl(func(w http.ResponseWriter, r *http.Request) {
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

	router.HandleFunc("/chat-svc/evens", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.Events(w, r)
	})).Methods("OPTIONS", "GET")

	router.HandleFunc("/prompt-svc/prompt", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.AddPrompt(w, r)
	})).Methods("OPTIONS", "POST")

	router.HandleFunc("/prompt-svc/prompt/{promptId}", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.RemovePrompt(w, r)
	})).Methods("OPTIONS", "DELETE")

	router.HandleFunc("/prompt-svc/prompts/{threadId}/responses/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.SubscribeToPromptResponses(w, r)
	})).Methods("OPTIONS", "GET")

	router.HandleFunc("/prompt-svc/prompts", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.ListPrompts(w, r)
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
	router.HandleFunc("/user-svc/organization/{organizationId}/user", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.AddUserToOrganization(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/organization/{organizationId}/user/{userId}", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.RemoveUserFromOrganization(w, r)
	})).Methods("OPTIONS", "DELETE")
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

	router.HandleFunc("/dynamic-svc/object", appl(func(w http.ResponseWriter, r *http.Request) {
		dynamicService.Create(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/dynamic-svc/objects/update", appl(func(w http.ResponseWriter, r *http.Request) {
		dynamicService.Update(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/dynamic-svc/objects/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		dynamicService.Delete(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/dynamic-svc/objects", appl(func(w http.ResponseWriter, r *http.Request) {
		dynamicService.Query(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/dynamic-svc/object/{objectId}", appl(func(w http.ResponseWriter, r *http.Request) {
		dynamicService.Upsert(w, r)
	})).Methods("OPTIONS", "PUT")

	router.HandleFunc("/registry-svc/nodes", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.List(w, r)
	})).Methods("OPTIONS", "POST")

	router.HandleFunc("/policy-svc/check", appl(func(w http.ResponseWriter, r *http.Request) {
		policyService.Check(w, r)
	})).Methods("OPTIONS", "POST")

	router.HandleFunc("/policy-svc/instance/{instanceId}", appl(func(w http.ResponseWriter, r *http.Request) {
		policyService.UpsertInstance(w, r)
	})).Methods("OPTIONS", "PUT")

	router.HandleFunc("/registry-svc/service-instances", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.ListServiceInstances(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/registry-svc/service-definitions", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.ListServiceDefinitions(w, r)
	})).Methods("OPTIONS", "POST")
	router.HandleFunc("/registry-svc/service-instance", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.RegisterServiceInstance(w, r)
	})).Methods("OPTIONS", "PUT")
	router.HandleFunc("/registry-svc/service-definition", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.SaveServiceDefinition(w, r)
	})).Methods("OPTIONS", "PUT")
	router.HandleFunc("/registry-svc/service-instance/{id}", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.RemoveServiceInstance(w, r)
	})).Methods("OPTIONS", "DELETE")
	router.HandleFunc("/registry-svc/service-definition/{id}", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.DeleteServiceDefinition(w, r)
	})).Methods("OPTIONS", "DELETE")

	router.HandleFunc("/deploy-svc/deployment", appl(func(w http.ResponseWriter, r *http.Request) {
		deployService.SaveDeployment(w, r)
	})).Methods("OPTIONS", "PUT")
	router.HandleFunc("/deploy-svc/deployments", appl(func(w http.ResponseWriter, r *http.Request) {
		deployService.ListDeployments(w, r)
	})).Methods("OPTIONS", "POST")

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
		err = dynamicService.Start()
		if err != nil {
			return err
		}
		err = policyService.Start()
		if err != nil {
			return err
		}
		err = registryService.Start()
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
