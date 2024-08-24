/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package policyservice

import (
	"sync"

	sdk "github.com/singulatron/singulatron/sdk/go"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/router"

	policytypes "github.com/singulatron/singulatron/localtron/internal/services/policy/types"
)

type PolicyService struct {
	router *router.Router

	instancesStore  datastore.DataStore
	credentialStore datastore.DataStore

	instances []*policytypes.Instance

	rateLimiters sync.Map // Map to store rate limiters
	mutex        sync.Mutex
}

func NewPolicyService(
	router *router.Router,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*PolicyService, error) {

	instancesStore, err := datastoreFactory("policySvcInstances", &policytypes.Instance{})
	if err != nil {
		return nil, err
	}

	credentialStore, err := datastoreFactory("policySvcCredentials", &sdk.Credential{})
	if err != nil {
		return nil, err
	}

	service := &PolicyService{
		router:          router,
		instancesStore:  instancesStore,
		credentialStore: credentialStore,
	}

	return service, nil
}

func (cs *PolicyService) Start() error {
	instances, err := cs.instancesStore.Query().Find()
	if err != nil {
		return err
	}
	for _, instanceI := range instances {
		instance := instanceI.(*policytypes.Instance)
		cs.instances = append(cs.instances, instance)
	}

	token, err := sdk.RegisterService("policy-svc", "Policy Service", cs.router, cs.credentialStore)
	if err != nil {
		return err
	}
	cs.router = cs.router.SetBearerToken(token)

	return cs.registerPermissions()
}
