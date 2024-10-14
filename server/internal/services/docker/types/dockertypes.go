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
	// Name is the name of the container
	Name string `json:"name,omitempty"`

	// Hash is a unique identifier for the container
	Hash string `json:"hash,omitempty"`

	// Envs are environment variables to set in the container
	Envs []string `json:"envs,omitempty"`

	// Labels are metadata labels associated with the container
	Labels map[string]string `json:"labels,omitempty"`

	// PersistentPaths are paths that should be persisted across container restarts
	PersistentPaths []string `json:"persistentPaths,omitempty"`

	// GPUEnabled specifies if GPU support is enabled
	GPUEnabled bool `json:"gpuEnabled,omitempty"`

	// Assets maps environment variable names to file URLs.
	// Example: {"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf"}
	// These files are downloaded by the Download Svc and mounted in the container.
	// The environment variable `MODEL` will point to the local file path in the container.
	Assets map[string]string `json:"assets,omitempty"`
}

type LaunchContainerRequest struct {
	// Image is the Docker image to use for the container
	Image string `json:"image" example:"nginx:latest" binding:"required"`

	// Port is the port number that the container will expose
	Port int `json:"port" example:"8080" binding:"required"`

	// HostPort is the port on the host machine that will be mapped to the container's port
	HostPort int `json:"hostPort" example:"8081"`

	// Options provides additional options for launching the container
	Options *LaunchContainerOptions `json:"options"`
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
