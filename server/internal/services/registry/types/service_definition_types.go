/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

type ServiceDefinition struct {
	// No id as there is only one definition per service slug.

	ServiceSlug string     `json:"serviceSlug,omitempty" example:"user-svc" binding:"required"` // The User Svc slug of the service whose instance is being registered.
	APISpecs    []APISpec  `json:"apiSpecs,omitempty"`                                          // API Specs such as OpenAPI definitions etc.
	Clients     []Client   `json:"clients,omitempty"`
	Image       *ImageSpec `json:"image,omitempty"` // Container specifications for Docker, K8s, etc.                                        // Programming language clients.
}

func (s ServiceDefinition) GetId() string {
	return s.ServiceSlug
}

type APISpec struct {
	URL          string            `json:"url,omitempty"`          // URL to the OpenAPI file or other API definition
	ProtocolType string            `json:"protocolType,omitempty"` // Protocol type (e.g., OpenAPI, Swagger, etc.)
	Version      string            `json:"version,omitempty"`      // Version of the API specification (e.g., 3.0.0)
	Metadata     map[string]string `json:"metadata,omitempty"`     // Additional metadata about the API (e.g., author, license, etc.)
}

type Client struct {
	Language Language `json:"language,omitempty" example:"JavaScript" binding:"required"`               // Programming language.
	URL      string   `json:"url,omitempty" example:"https://example.com/client.js" binding:"required"` // The URL of the client.
}

type Image struct {
	// Image name/URL of the container
	Image string `json:"image" example:"nginx:latest" binding:"required"`

	// Runtime environment (e.g., Docker, K8s)
	Runtime RuntimeType `json:"runtime,omitempty"`

	// Port is the port number that the container will expose
	Port int `json:"port" example:"8080" binding:"required"`

	// PersistentPaths are paths inside the container which should be persisted across container restarts
	PersistentPaths []string `json:"persistentPaths,omitempty"`

	// GPUEnabled specifies if GPU support is enabled
	GPUEnabled bool `json:"gpuEnabled,omitempty"`
}

type ImageSpec struct {
	// Image is the Docker image to use for the container
	Image string `json:"image" example:"nginx:latest" binding:"required"`

	// Port is the port number that the container will expose
	Port int `json:"port" example:"8080" binding:"required"`
}

type RuntimeType string

const (
	RuntimeDocker     RuntimeType = "Docker"
	RuntimeContainerd RuntimeType = "containerd"
	RuntimeCRIO       RuntimeType = "CRI-O"
)

type Language string

const (
	JavaScript Language = "JavaScript"
	Python     Language = "Python"
	Java       Language = "Java"
	CSharp     Language = "C#"
	CPlusPlus  Language = "C++"
	Ruby       Language = "Ruby"
	Go         Language = "Go"
	Swift      Language = "Swift"
	PHP        Language = "PHP"
	TypeScript Language = "TypeScript"
	Kotlin     Language = "Kotlin"
	Scala      Language = "Scala"
	Perl       Language = "Perl"
	Rust       Language = "Rust"
	Haskell    Language = "Haskell"
	Clojure    Language = "Clojure"
	Elixir     Language = "Elixir"
	ObjectiveC Language = "Objective-C"
	FSharp     Language = "F#"
)

type SaveServiceDefinitionRequest struct {
	ServiceDefinition ServiceDefinition `json:"serviceDefinition,omitempty"`
}

type SaveServiceDefinitionResponse struct {
}

type ListServiceDefinitionsRequest struct{}

type ListServiceDefinitionsResponse struct {
	ServiceDefinitions []*ServiceDefinition `json:"serviceDefinitions,omitempty"`
}
