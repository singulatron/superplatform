package di

import (
	"log/slog"
	"os"
	"path"

	"github.com/singulatron/singulatron/localtron/clients/llm"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/datastore/localstore"
	"github.com/singulatron/singulatron/localtron/logger"
	appservice "github.com/singulatron/singulatron/localtron/services/app"
	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
	chatservice "github.com/singulatron/singulatron/localtron/services/chat"
	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
	dockerservice "github.com/singulatron/singulatron/localtron/services/docker"
	dockertypes "github.com/singulatron/singulatron/localtron/services/docker/types"
	downloadservice "github.com/singulatron/singulatron/localtron/services/download"
	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"
	genericservice "github.com/singulatron/singulatron/localtron/services/generic"
	generictypes "github.com/singulatron/singulatron/localtron/services/generic/types"
	modelservice "github.com/singulatron/singulatron/localtron/services/model"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
	nodeservice "github.com/singulatron/singulatron/localtron/services/node"
	nodetypes "github.com/singulatron/singulatron/localtron/services/node/types"
	promptservice "github.com/singulatron/singulatron/localtron/services/prompt"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

const singulatronFolder = ".singulatron"

type Universe struct {
	ConfigService   configtypes.ConfigServiceI
	PromptService   prompttypes.PromptServiceI
	UserService     usertypes.UserServiceI
	FirehoseService firehosetypes.FirehoseServiceI
	ChatService     chattypes.ChatServiceI
	GenericService  generictypes.GenericServiceI
	ModelService    modeltypes.ModelServiceI
	DownloadService downloadtypes.DownloadServiceI
	AppService      apptypes.AppServiceI
	DockerService   dockertypes.DockerServiceI
	NodeService     nodetypes.NodeServiceI

	LLMClient llm.ClientI
}

type UniverseOptions struct {
	Test             bool
	Pre              Universe
	DatastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
}

func BigBang(options UniverseOptions) (*Universe, error) {
	universe := &Universe{}

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

	if options.Pre.ConfigService != nil {
		universe.ConfigService = options.Pre.ConfigService
	} else {
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

		universe.ConfigService = configService
	}

	if options.DatastoreFactory == nil {
		localStorePath := path.Join(universe.ConfigService.GetConfigDirectory(), "data")
		err = os.MkdirAll(localStorePath, 0755)
		if err != nil {
			logger.Error("Creating data folder failed", slog.String("error", err.Error()))
			os.Exit(1)
		}

		options.DatastoreFactory = func(tableName string, instance any) (datastore.DataStore, error) {
			return localstore.NewLocalStore(instance, path.Join(localStorePath, tableName))
		}
	}

	if options.Pre.UserService != nil {
		universe.UserService = options.Pre.UserService
	} else {
		userService, err := userservice.NewUserService(
			universe.ConfigService,
			options.DatastoreFactory,
		)
		if err != nil {
			logger.Error("User service start failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		universe.UserService = userService
	}

	// hacks to avoid import cycles
	universe.ConfigService.SetUpsertPermissionFunc(universe.UserService.UpsertPermission)
	universe.ConfigService.SetAddPermissionToRoleFunc(universe.UserService.AddPermissionToRole)

	err = universe.ConfigService.Start()
	if err != nil {
		logger.Error("Config service start failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if options.Pre.FirehoseService != nil {
		universe.FirehoseService = options.Pre.FirehoseService
	} else {
		firehoseService, err := firehoseservice.NewFirehoseService(universe.UserService)
		if err != nil {
			logger.Error("Firehose service creation failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		universe.FirehoseService = firehoseService
	}

	universe.ConfigService.SetEventCallback(universe.FirehoseService.Publish)

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

	if options.Pre.DownloadService != nil {
		universe.DownloadService = options.Pre.DownloadService
	} else {
		downloadService, err := downloadservice.NewDownloadService(universe.FirehoseService, universe.UserService)
		if err != nil {
			logger.Error("Download service creation failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		universe.DownloadService = downloadService
	}

	universe.DownloadService.SetDefaultFolder(downloadFolder)
	universe.DownloadService.SetStateFilePath(path.Join(singulatronFolder, "downloads.json"))

	err = universe.DownloadService.Start()
	if err != nil {
		logger.Error("Download service start failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if options.Pre.DockerService != nil {
		universe.DockerService = options.Pre.DockerService
	} else {
		dockerService, err := dockerservice.NewDockerService(
			universe.DownloadService,
			universe.UserService,
			universe.ConfigService,
		)
		if err != nil {
			logger.Error("Docker service creation failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		universe.DockerService = dockerService
	}

	if options.Pre.ModelService != nil {
		universe.ModelService = options.Pre.ModelService
	} else {
		modelService, err := modelservice.NewModelService(
			universe.DownloadService,
			universe.UserService,
			universe.ConfigService,
			universe.DockerService,
			options.DatastoreFactory,
		)
		if err != nil {
			logger.Error("Model service creation failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		universe.ModelService = modelService
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

	return universe, nil
}
