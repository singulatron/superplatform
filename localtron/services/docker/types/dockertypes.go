/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dockertypes

type ModelLaunchRequest struct{}

type OnModelLaunch struct {
	Error *string `json:"error,omitempty"`
}

type OnDockerInfo struct {
	HasDocker           bool    `json:"hasDocker"`
	DockerDaemonAddress *string `json:"dockerDaemonAddress,omitempty"`
	Error               *string `json:"error,omitempty"`
}

//
// Events
//

// @todo nothing to trigger this yet
const EventDockerInfoUpdatedName = "dockerInfoUpdated"

type EventDockerInfoUpdated struct {
	ThreadId string `json:"threadId"`
}

func (e EventDockerInfoUpdated) Name() string {
	return EventDockerInfoUpdatedName
}
