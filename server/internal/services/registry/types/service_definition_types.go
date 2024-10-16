/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

type ServiceDefinition struct {
	ServiceSlug string    `json:"serviceSlug,omitempty" example:"user-svc" binding:"required"` // The User Svc slug of the service whose instance is being registered.
	APISpecs    []APISpec `json:"apiSpecs,omitempty"`                                          // API Specs such as OpenAPI definitions etc.
	Clients     []Client  `json:"clients,omitempty"`                                           // Programming language clients.
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

type RegisterServiceRequest struct {
	Hostname string `json:"hostname"`       // Hostname of the node
	IP       string `json:"ip,omitempty"`   // IP of the node. Optional: If not provided, resolved by hostname.
	GPUs     []*GPU `json:"gpus,omitempty"` // List of GPUs available on the node
}

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
