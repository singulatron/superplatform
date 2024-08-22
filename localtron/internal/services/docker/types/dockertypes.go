/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package docker_svc

type ErrorResponse struct {
	Error string `json:"error"`
}

type LaunchInfo struct {
	NewContainerStarted bool
	PortNumber          int
}

type ModelLaunchRequest struct{}

type OnModelLaunch struct {
	Error *string `json:"error,omitempty"`
}

type GetInfoResponse struct {
	Info *DockerInfo `json:"info"`
}

type DockerInfo struct {
	HasDocker           bool    `json:"hasDocker"`
	DockerDaemonAddress *string `json:"dockerDaemonAddress,omitempty"`
	Error               *string `json:"error,omitempty"`
}

type LaunchContainerOptions struct {
	Name            string            `json:"name,omitempty"`
	Hash            string            `json:"hash,omitempty"`
	Envs            []string          `json:"envs,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	PersistentPaths []string          `json:"persistentPaths,omitempty"`
	GPUEnabled      bool              `json:"gpuEnabled,omitempty"`
	// Asset is a map of envar name to file URL.
	// eg. {"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf"}
	// This file will be downloaded with the Download Svc and the local file will be mounted in the container
	// and the envar `MODEL=/local/path/to/file` will be available in the container launched by the Docker Svc.
	Assets map[string]string `json:"assets,omitempty"`
}

type LaunchContainerRequest struct {
	Image    string                  `json:"image"`
	Port     int                     `json:"port"`
	HostPort int                     `json:"hostPort"`
	Options  *LaunchContainerOptions `json:"options"`
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

type ContainerIsRunningRequest struct {
	Hash string `json:"hash"`
}

type ContainerIsRunningResponse struct {
	IsRunning bool `json:"isRunning"`
}

type GetDockerHostRequest struct{}

type GetDockerHostResponse struct {
	Host string `json:"host"`
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
