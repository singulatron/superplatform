/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dynamicservice

import (
	"errors"

	"github.com/google/uuid"
	sdk "github.com/singulatron/singulatron/sdk/go"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/router"

	dynamictypes "github.com/singulatron/singulatron/localtron/internal/services/dynamic/types"
)

type DynamicService struct {
	router          *router.Router
	store           datastore.DataStore
	credentialStore datastore.DataStore
}

func NewDynamicService(
	router *router.Router,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*DynamicService, error) {
	store, err := datastoreFactory("genericSvcObjects", &dynamictypes.GenericObject{})
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
	token, err := sdk.RegisterService("dynamic-svc", "Generic Service", g.router, g.credentialStore)
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

func (g *DynamicService) upsert(request *dynamictypes.UpsertObjectRequest) error {
	vI, found, err := g.store.Query(
		datastore.Id(request.Object.Id),
	).FindOne()
	if err != nil {
		return err
	}

	if found {
		v := vI.(*dynamictypes.GenericObject)
		if v.UserId != request.Object.UserId {
			return errors.New("unauthorized")
		}
	}

	return g.store.Upsert(request.Object)
}

func (g *DynamicService) upsertMany(request *dynamictypes.UpsertManyRequest) error {
	objectIs := []datastore.Row{}
	for _, object := range request.Objects {
		objectIs = append(objectIs, object)
	}
	return g.store.UpsertMany(objectIs)
}

func (g *DynamicService) update(tableName string, userId string, conditions []datastore.Condition, object *dynamictypes.GenericObject) error {
	if len(conditions) == 0 {
		return errors.New("no conditions")
	}

	conditions = append(conditions, datastore.Equal(datastore.Field("table"), tableName))
	conditions = append(conditions, datastore.Equal(
		datastore.Field("userId"),
		userId,
	))

	return g.store.Query(
		conditions[0], conditions[1:]...,
	).Update(object)
}

func (g *DynamicService) delete(tableName string, userId string, conditions []datastore.Condition) error {
	if len(conditions) == 0 {
		return errors.New("no conditions")
	}

	conditions = append(conditions, datastore.Equal(datastore.Field("table"), tableName))
	conditions = append(conditions, datastore.Equal(
		datastore.Field("userId"),
		userId,
	))

	return g.store.Query(
		conditions[0], conditions[1:]...,
	).Delete()
}
