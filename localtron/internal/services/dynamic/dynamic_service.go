/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dynamicservice

import (
	"context"
	"errors"

	"github.com/google/uuid"
	sdk "github.com/singulatron/singulatron/sdk/go"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/router"

	dynamictypes "github.com/singulatron/singulatron/localtron/internal/services/dynamic/types"

	clients "github.com/singulatron/singulatron/clients/go"
)

type DynamicService struct {
	router          *router.Router
	store           datastore.DataStore
	credentialStore datastore.DataStore
	publicKey       string
	client          *clients.APIClient
}

func NewDynamicService(
	router *router.Router,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*DynamicService, error) {
	store, err := datastoreFactory("genericSvcObjects", &dynamictypes.Object{})
	if err != nil {
		return nil, err
	}
	credentialStore, err := datastoreFactory("genericSvcCredentials", &sdk.Credential{})
	if err != nil {
		return nil, err
	}

	service := &DynamicService{
		credentialStore: credentialStore,
		router:          router,
		store:           store,
	}

	return service, nil
}

func (g *DynamicService) Start() error {
	g.client = clients.NewAPIClient(&clients.Configuration{
		Servers: clients.ServerConfigurations{
			{
				URL:         g.router.Address(),
				Description: "Default server",
			},
		},
	})

	pk, _, err := g.client.UserSvcAPI.GetPublicKey(context.Background()).Execute()
	if err != nil {
		return err
	}
	g.publicKey = *pk.PublicKey

	token, err := sdk.RegisterService("dynamic-svc", "Dynamic Svc", g.router, g.credentialStore)
	if err != nil {
		return err
	}
	g.router = g.router.SetBearerToken(token)

	return g.registerPermissions()
}

func (g *DynamicService) create(request *dynamictypes.CreateObjectRequest) error {
	if request.Object.Id == "" {
		request.Object.Id = uuid.NewString()
	}
	return g.store.Create(request.Object)
}

func (g *DynamicService) createMany(request *dynamictypes.CreateManyRequest) error {
	objectIs := []datastore.Row{}
	for _, object := range request.Objects {
		objectIs = append(objectIs, object)
	}

	return g.store.CreateMany(objectIs)
}

func (g *DynamicService) upsert(writers []string, request *dynamictypes.UpsertObjectRequest) error {
	vI, found, err := g.store.Query(
		datastore.Id(request.Object.Id),
	).FindOne()
	if err != nil {
		return err
	}

	if found {
		v := vI.(*dynamictypes.Object)

		if !intersects(writers, v.Writers) {
			return errors.New("unauthorized")
		}
	}

	return g.store.Upsert(request.Object)
}

func intersects(slice1, slice2 []string) bool {
	elemMap := make(map[string]struct{})
	for _, elem := range slice1 {
		elemMap[elem] = struct{}{}
	}

	for _, elem := range slice2 {
		if _, found := elemMap[elem]; found {
			return true
		}
	}

	return false
}

func (g *DynamicService) upsertMany(request *dynamictypes.UpsertManyRequest) error {
	objectIs := []datastore.Row{}
	for _, object := range request.Objects {
		objectIs = append(objectIs, object)
	}
	return g.store.UpsertMany(objectIs)
}

func (g *DynamicService) update(tableName string, userId string, conditions []datastore.Filter, object *dynamictypes.Object) error {
	if len(conditions) == 0 {
		return errors.New("no conditions")
	}

	conditions = append(conditions, datastore.Equals(datastore.Field("table"), tableName))
	conditions = append(conditions, datastore.Equals(
		datastore.Field("userId"),
		userId,
	))

	return g.store.Query(
		conditions...,
	).Update(object)
}

func (g *DynamicService) delete(tableName string, userId string, conditions []datastore.Filter) error {
	if len(conditions) == 0 {
		return errors.New("no conditions")
	}

	conditions = append(conditions, datastore.Equals(datastore.Field("table"), tableName))
	conditions = append(conditions, datastore.Equals(
		datastore.Field("userId"),
		userId,
	))

	return g.store.Query(
		conditions...,
	).Delete()
}
