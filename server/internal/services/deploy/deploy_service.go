/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package deployservice

import (
	"context"

	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/datastore"
	"github.com/singulatron/superplatform/sdk/go/lock"
	deploy "github.com/singulatron/superplatform/server/internal/services/deploy/types"
)

type DeployService struct {
	clientFactory sdk.ClientFactory

	lock  lock.DistributedLock
	token string

	credentialStore datastore.DataStore
	deploymentStore datastore.DataStore
}

func NewDeployService(
	clientFactory sdk.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any,
	) (datastore.DataStore, error)) (*DeployService, error) {

	credentialStore, err := datastoreFactory("deploySvcCredentials", &sdk.Credential{})
	if err != nil {
		return nil, err
	}
	deploymentStore, err := datastoreFactory("deploySvcDeployments", &deploy.Deployment{})
	if err != nil {
		return nil, err
	}

	service := &DeployService{
		clientFactory: clientFactory,

		lock:            lock,
		credentialStore: credentialStore,
		deploymentStore: deploymentStore,
	}

	return service, nil
}

func (ns *DeployService) Start() error {
	ctx := context.Background()
	ns.lock.Acquire(ctx, "deploy-svc-start")
	defer ns.lock.Release(ctx, "deploy-svc-start")

	token, err := sdk.RegisterServiceNoRouter(ns.clientFactory.Client().UserSvcAPI, "deploy-svc", "Deploy Service", ns.credentialStore)
	if err != nil {
		return err
	}

	ns.token = token

	go ns.loop()

	return ns.registerPermissions()
}
