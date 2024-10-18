/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

import (
	"errors"
	"fmt"
	"time"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type ServiceInstance struct {
	ID            string    `json:"id,omitempty" example:"https://api.com:999/user-svc" binding:"required"` // Required: ID of the service instance
	ServiceSlug   string    `json:"serviceSlug,omitempty" example:"user-svc" binding:"required"`            // The User Svc slug of the service whose instance is being registered.
	URL           string    `json:"url,omitempty" example:"https://myserver.com:5981" binding:"required"`   // Full address URL of the service instance.
	Path          string    `json:"path,omitempty" example:"/your-svc"`                                     // Path of the service instance address. Optional (e.g., "/api")
	NodeURL       string    `json:"nodeUrl,omitempty" example:"https://myserver.com:58231"`                 // URL of the Singulatron daemon
	LastHeartbeat time.Time `json:"lastHeartbeat,omitempty"`                                                // Last time the instance gave a sign of life

	// URL alternative fields

	Scheme string `json:"scheme,omitempty" example:"https"`      // Scheme of the service instance address. Required if URL is not provided.
	Host   string `json:"host,omitempty" example:"myserver.com"` // Host of the service instance address. Required if URL is not provided
	IP     string `json:"ip,omitempty" example:"8.8.8.8"`        // IP of the service instance address. Optional: to register by IP instead of host
	Port   int    `json:"port,omitempty" example:"8080"`         // Port of the service instance address. Required if URL is not provided

}

func (s *ServiceInstance) GetId() string {
	return s.ID
}

func (s *ServiceInstance) DeriveID() string {
	if s.URL != "" {
		return fmt.Sprintf("%s/%s", s.URL, s.ServiceSlug)
	}

	var constructedURL string
	if s.Host != "" {
		constructedURL = fmt.Sprintf("%s://%s:%d", s.Scheme, s.Host, s.Port)
	} else {
		constructedURL = fmt.Sprintf("%s://%s:%d", s.Scheme, s.IP, s.Port)
	}

	return fmt.Sprintf("%s/%s", constructedURL, s.ServiceSlug)
}

var ErrNotFound = errors.New("service not found")

// RegisterServiceInstanceRequest represents the request body to register a service.
//
// The user must provide either:
// 1. A full URL (e.g., "https://myserver.com:5981") -OR-
// 2. A combination of the following fields:
//   - scheme: "http" or "https" (required if URL is not provided)
//   - host: the domain of the service (required if URL is not provided)
//   - port: the port number (required if URL is not provided)
//
// Additionally, if both host and port are provided, they cannot both be specified at the same time.
// The IP field is optional and can be used for registration by IP instead of host.
type RegisterServiceInstanceRequest struct {
	ServiceSlug string `json:"slug,omitempty" example:"user-svc" binding:"required"` // The User Svc slug of the service whose instance is being registered.
	URL         string `json:"url,omitempty" example:"https://myserver.com:5981"`    // Full address URL of the service instance.
	Scheme      string `json:"scheme,omitempty" example:"https"`                     // Scheme of the service instance address. Required if URL is not provided.
	Host        string `json:"host,omitempty" example:"myserver.com"`                // Host of the service instance address. Required if URL is not provided
	IP          string `json:"ip,omitempty" example:"8.8.8.8"`                       // IP of the service instance address. Optional: to register by IP instead of host
	Port        int    `json:"port,omitempty" example:"8080"`                        // Port of the service instance address. Required if URL is not provided
	Path        string `json:"path,omitempty" example:"/your-svc"`                   // Path of the service instance address. Optional (e.g., "/api")
}

type RegisterServiceInstanceResponse struct {
}

type ListServiceInstancesResponse struct {
	Instances []ServiceInstance `json:"instances,omitempty"`
}
