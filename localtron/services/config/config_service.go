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

	"github.com/singulatron/singulatron/localtron/router"
	types "github.com/singulatron/singulatron/localtron/services/config/types"

	"github.com/singulatron/singulatron/localtron/logger"
)

const DefaultModelId = `huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf`

type ConfigService struct {
	router *router.Router

	ConfigDirectory string
	ConfigFileName  string
	config          types.Config
	configFileMutex sync.Mutex
}

func NewConfigService(router *router.Router) (*ConfigService, error) {
	cs := &ConfigService{
		ConfigFileName: "config.yaml",
		router:         router,
	}

	return cs, nil
}

func (cs *ConfigService) GetConfigDirectory() string {
	return cs.ConfigDirectory
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
