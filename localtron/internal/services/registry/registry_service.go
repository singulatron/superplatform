/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package registryservice

import (
	"os"

	registry "github.com/singulatron/singulatron/localtron/internal/services/registry/types"
	sdk "github.com/singulatron/singulatron/sdk/go"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/router"
)

type RegistryService struct {
	Hostname string

	router *router.Router

	credentialStore      datastore.DataStore
	serviceInstanceStore datastore.DataStore
}

func NewRegistryService(router *router.Router, datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)) (*RegistryService, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	credentialStore, err := datastoreFactory("registrySvcCredentials", &sdk.Credential{})
	if err != nil {
		return nil, err
	}
	serviceInstanceStore, err := datastoreFactory("registrySvcServiceInstances", &registry.ServiceInstance{})
	if err != nil {
		return nil, err
	}

	service := &RegistryService{
		Hostname:             hostname,
		router:               router,
		credentialStore:      credentialStore,
		serviceInstanceStore: serviceInstanceStore,
	}

	return service, nil
}

func (ns *RegistryService) Start() error {
	return ns.registerPermissions()
}
