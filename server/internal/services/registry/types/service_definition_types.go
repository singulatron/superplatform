/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

type ServiceDefinition struct {
	ServiceSlug string          `json:"serviceSlug,omitempty" example:"user-svc" binding:"required"` // The User Svc slug of the service whose instance is being registered.
	APISpecs    []APISpec       `json:"apiSpecs,omitempty"`                                          // API Specs such as OpenAPI definitions etc.
	Clients     []Client        `json:"clients,omitempty"`
	Containers  []ContainerSpec `json:"containers,omitempty"` // Container specifications for Docker, K8s, etc.                                        // Programming language clients.
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

type ContainerSpec struct {
	ImageSpec ImageSpec   `json:"imageSpec,omitempty"` // Container image specifications.
	Runtime   RuntimeType `json:"runtime,omitempty"`   // Runtime environment (e.g., Docker, K8s).

	// HostPort is the port on the host machine that will be mapped to the container's port
	HostPort int `json:"hostPort" example:"8081"`

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

type RegisterServiceResponse struct {
}
