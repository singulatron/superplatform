/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package registry_svc

import (
	"testing"
)

// TestGenerateID tests the GenerateID method of the ServiceInstance struct.
func TestGenerateID(t *testing.T) {
	tests := []struct {
		name       string
		instance   ServiceInstance
		expectedID string
	}{
		{
			name: "With URL",
			instance: ServiceInstance{
				URL:         "https://myserver.com:5981",
				ServiceSlug: "user-svc",
			},
			expectedID: "https://myserver.com:5981/user-svc",
		},
		{
			name: "With IP",
			instance: ServiceInstance{
				Scheme:      "https",
				IP:          "192.168.1.1",
				Port:        999,
				ServiceSlug: "user-svc",
			},
			expectedID: "https://192.168.1.1:999/user-svc",
		},
		{
			name: "With Host",
			instance: ServiceInstance{
				Scheme:      "http",
				Host:        "api.com",
				Port:        80,
				ServiceSlug: "user-svc",
			},
			expectedID: "http://api.com:80/user-svc",
		},
		{
			name: "With Host and Port",
			instance: ServiceInstance{
				Scheme:      "http",
				Host:        "api.com",
				Port:        8080,
				ServiceSlug: "user-svc",
			},
			expectedID: "http://api.com:8080/user-svc",
		},
		{
			name: "With Missing URL, IP, and Host",
			instance: ServiceInstance{
				Scheme:      "http",
				Port:        8080,
				ServiceSlug: "user-svc",
			},
			expectedID: "http://:8080/user-svc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := tt.instance.DeriveID()

			if id != tt.expectedID {
				t.Errorf("DeriveID() = %v, want %v", id, tt.expectedID)
			}
		})
	}
}
