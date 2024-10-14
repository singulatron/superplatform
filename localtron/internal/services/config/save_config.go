/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package configservice

import (
	"context"
	"io/ioutil"
	"path"

	"github.com/pkg/errors"
	"github.com/singulatron/superplatform/sdk/go/logger"
	types "github.com/singulatron/superplatform/server/internal/services/config/types"
	firehosetypes "github.com/singulatron/superplatform/server/internal/services/firehose/types"
	"gopkg.in/yaml.v2"
)

func (cs *ConfigService) saveConfig(config types.Config) error {
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

	ev := types.EventConfigUpdate{}
	err = cs.router.Post(context.Background(), "firehose-svc", "/event", firehosetypes.EventPublishRequest{
		Event: &firehosetypes.Event{
			Name: ev.Name(),
			Data: ev,
		},
	}, nil)
	if err != nil {
		logger.Error("Failed to publish: %v", err)
	}

	return nil
}
