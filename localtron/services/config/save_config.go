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
	"io/ioutil"
	"path"

	"github.com/pkg/errors"
	types "github.com/singulatron/singulatron/localtron/services/config/types"
	"gopkg.in/yaml.v2"
)

func (cs *ConfigService) SaveConfig(config types.Config) error {
	cs.configFileMutex.Lock()
	defer cs.configFileMutex.Unlock()

	cs.config = config

	data, err := yaml.Marshal(cs.config)
	if err != nil {
		return errors.Wrap(err, "error saving config")
	}
	if err := ioutil.WriteFile(path.Join(cs.ConfigDirectory, cs.ConfigFileName), data, 0644); err != nil {
		return errors.Wrap(err, "error writing config file")
	}

	cs.eventCallback(types.EventConfigUpdate{})

	return nil
}
