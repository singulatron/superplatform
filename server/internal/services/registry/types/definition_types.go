/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

// Definition can be roughly thought of as a container image + some additional details
// like clients, api definition etc.
type Definition struct {
	Id string `json:"id,omitempty" binding:"required"`

	// Container specifications for Docker, K8s, etc.
	Image *ImageSpec `json:"image,omitempty" binding:"required"`

	// API Specs such as OpenAPI definitions etc.
	APISpecs []APISpec `json:"apiSpecs,omitempty"`

	// Programming language clients such as on npm or GitHub.
	Clients []Client `json:"clients,omitempty"`
}

func (s Definition) GetId() string {
	return s.Id
}

type ImageSpec struct {
	// Name is the container image name/URL to use for the container
	Name string `json:"name" example:"nginx:latest" binding:"required"`

	// Port is the port number that the container will listen on internally
	Port int `json:"port" example:"8080" binding:"required"`
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

type SaveDefinitionRequest struct {
	Definition Definition `json:"serviceDefinition,omitempty"`
}

type SaveDefinitionResponse struct {
}

type ListDefinitionsRequest struct{}

type ListDefinitionsResponse struct {
	Definitions []*Definition `json:"serviceDefinitions,omitempty"`
}
