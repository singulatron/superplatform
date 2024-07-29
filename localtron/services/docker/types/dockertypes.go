/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dockertypes

type LaunchOptions struct {
	Name       string
	Envs       []string
	Labels     map[string]string
	HostBinds  []string
	GPUEnabled bool
	Hash       string
}

type LaunchInfo struct {
	NewContainerStarted bool
	PortNumber          int
}

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

type LaunchContainerRequest struct {
	Image    string         `json:"image"`
	Port     int            `json:"port"`
	HostPort int            `json:"hostPort"`
	Options  *LaunchOptions `json:"options"`
}

type LaunchContainerResponse struct {
	Info *LaunchInfo `json:"info"`
}

type GetContainerSummaryRequest struct {
	Hash  string `json:"hash"`
	Lines int    `json:"lines"`
}

type GetContainerSummaryResponse struct {
	Summary string `json:"summary"`
}

type HashIsRunningRequest struct {
	Hash string `json:"hash"`
}

type HashIsRunningResponse struct {
	IsRunning bool `json:"isRunning"`
}

type GetDockerHostRequest struct{}

type GetDockerHostResponse struct {
	Host string `json:"host"`
}
