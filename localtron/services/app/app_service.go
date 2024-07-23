/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package appservice

import (
	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

type AppService struct {
	configService   configtypes.ConfigServiceI
	userService     usertypes.UserServiceI
	firehoseService firehosetypes.FirehoseServiceI
}

func NewAppService(
	cs configtypes.ConfigServiceI,
	fs firehosetypes.FirehoseServiceI,
	userService usertypes.UserServiceI,
) (*AppService, error) {

	service := &AppService{
		configService:   cs,
		firehoseService: fs,
		userService:     userService,
	}

	return service, nil
}
