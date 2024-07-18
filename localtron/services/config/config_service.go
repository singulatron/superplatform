/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package configservice

import (
	"fmt"
	"io/ioutil"
	"log/slog"
	"os"
	"path"
	"sync"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	types "github.com/singulatron/singulatron/localtron/services/config/types"
	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"

	"github.com/singulatron/singulatron/localtron/logger"
)

const DefaultModelId = `huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf`

type ConfigService struct {
	// import cycle doesn't alllow use to have
	// the firehose service here
	EventCallback func(firehosetypes.Event)
	// defined like this to avoid passing in the user service
	UpsertPermission func(id, name, description string) (*usertypes.Permission, error)
	// defined like this to avoid passing in the user service
	AddPermissionToRole func(roleId, permissionId string) error

	ConfigDirectory   string
	ConfigFileName    string
	config            types.Config
	configFileMutex   sync.Mutex
	clientIdFileMutex sync.Mutex
	clientId          string
}

func NewConfigService() (*ConfigService, error) {
	cs := &ConfigService{
		ConfigFileName: "config.yaml",
	}

	return cs, nil
}

func (cs *ConfigService) Start() error {
	if cs.ConfigDirectory == "" {
		return fmt.Errorf("config service is missing a config directory option")
	}
	err := cs.registerPermissions()
	if err != nil {
		return err
	}

	err = cs.loadConfig()
	if err != nil {
		return err
	}
	return nil
}

func (cs *ConfigService) loadConfig() error {
	cs.configFileMutex.Lock()
	defer cs.configFileMutex.Unlock()

	if _, err := os.Stat(cs.ConfigDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(cs.ConfigDirectory, os.ModePerm); err != nil {
			return errors.Wrap(err, "error creating config directory")
		}
	}

	if _, err := os.Stat(path.Join(cs.ConfigDirectory, cs.ConfigFileName)); err == nil {
		data, err := ioutil.ReadFile(path.Join(cs.ConfigDirectory, cs.ConfigFileName))
		if err != nil {
			return errors.Wrap(err, "failed to read config")
		}

		if err := yaml.Unmarshal(data, &cs.config); err != nil {
			return errors.Wrap(err, "failed to unmarshal config")
		}
	} else {
		logger.Debug("Config file does not exist", slog.String("path", path.Join(cs.ConfigDirectory, cs.ConfigFileName)))
		cs.config = types.Config{}
	}

	if cs.config.Download.DownloadFolder == "" {
		cs.config.Download.DownloadFolder = path.Join(cs.ConfigDirectory, "downloads")
	}

	if cs.config.Model.CurrentModelId == "" {
		cs.config.Model.CurrentModelId = DefaultModelId
	}

	return nil
}
