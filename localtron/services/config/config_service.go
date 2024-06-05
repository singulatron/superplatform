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
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"

	"github.com/singulatron/singulatron/localtron/lib"
)

const defaultModelId = `https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_S.gguf`

type ConfigService struct {
	firehoseService *firehoseservice.FirehoseService

	ConfigDirectory   string
	ConfigFileName    string
	config            types.Config
	configFileMutex   sync.Mutex
	clientIdFileMutex sync.Mutex
	clientId          string
}

func NewConfigService(firehoseService *firehoseservice.FirehoseService) (*ConfigService, error) {
	cs := &ConfigService{
		firehoseService: firehoseService,

		ConfigFileName: "config.yaml",
	}

	return cs, nil
}

func (cs *ConfigService) Start() error {
	if cs.ConfigDirectory == "" {
		return fmt.Errorf("config service is missing a config directory option")
	}
	err := cs.loadConfig()
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
		lib.Logger.Debug("Config file does not exist", slog.String("path", path.Join(cs.ConfigDirectory, cs.ConfigFileName)))
		cs.config = types.Config{}
	}

	if cs.config.Download.DownloadFolder == "" {
		cs.config.Download.DownloadFolder = path.Join(cs.ConfigDirectory, "downloads")
	}

	if cs.config.Model.CurrentModelId == "" {
		cs.config.Model.CurrentModelId = defaultModelId
	}

	return nil
}
