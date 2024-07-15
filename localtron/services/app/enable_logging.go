/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package appservice

import (
	"github.com/pkg/errors"
)

func (a *AppService) EnableLogging() error {
	conf, err := a.configService.GetConfig()
	if err != nil {
		return errors.Wrap(err, "cannot get config")
	}
	conf.App.LoggingDisabled = false

	err = a.configService.SaveConfig(conf)
	if err != nil {
		return errors.Wrap(err, "cannot save config")
	}

	return nil
}
