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
	"fmt"
	"os"

	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/datastore"
	"github.com/singulatron/superplatform/sdk/go/lock"
	"github.com/singulatron/superplatform/sdk/go/router"
	deploy "github.com/singulatron/superplatform/server/internal/services/deploy/types"
)

type DeployService struct {
	router *router.Router
	lock   lock.DistributedLock

	credentialStore datastore.DataStore
	deploymentStore datastore.DataStore
}

func NewDeployService(
	address string,
	az string,
	region string,
	router *router.Router,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any,
	) (datastore.DataStore, error)) (*DeployService, error) {

	nodeUrl := address
	var err error

	if nodeUrl == "" {
		nodeUrl, err = os.Hostname()
		if err != nil {
			return nil, err
		}
		nodeUrl = fmt.Sprintf("%v:%v", nodeUrl, "58231")
	}

	credentialStore, err := datastoreFactory("deploySvcCredentials", &sdk.Credential{})
	if err != nil {
		return nil, err
	}
	deploymentStore, err := datastoreFactory("deploySvcDeployments", &deploy.Deployment{})
	if err != nil {
		return nil, err
	}

	service := &DeployService{
		router:          router,
		lock:            lock,
		credentialStore: credentialStore,
		deploymentStore: deploymentStore,
	}

	return service, nil
}

func (ns *DeployService) Start() error {
	ctx := context.Background()
	ns.lock.Acquire(ctx, "deploy-service-start")
	defer ns.lock.Release(ctx, "deploy-service-start")

	token, err := sdk.RegisterService("deploy-svc", "Deploy Service", ns.router, ns.credentialStore)
	if err != nil {
		return err
	}
	ns.router = ns.router.SetBearerToken(token)

	return ns.registerPermissions()
}
