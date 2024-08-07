/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package genericservice

import (
	"errors"

	"github.com/google/uuid"
	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/router"
	sdk "github.com/singulatron/singulatron/localtron/sdk/go"

	generictypes "github.com/singulatron/singulatron/localtron/services/generic/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

type GenericService struct {
	router          *router.Router
	store           datastore.DataStore
	credentialStore datastore.DataStore
}

func NewGenericService(
	router *router.Router,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*GenericService, error) {
	store, err := datastoreFactory("generic", &generictypes.GenericObject{})
	if err != nil {
		return nil, err
	}
	credentialStore, err := datastoreFactory("chat_credentials", &usertypes.Credential{})
	if err != nil {
		return nil, err
	}

	service := &GenericService{
		credentialStore: credentialStore,
		router:          router,
		store:           store,
	}

	return service, nil
}

func (g *GenericService) Start() error {
	token, err := sdk.RegisterService("generic-svc", "Generic Service", g.router, g.credentialStore)
	if err != nil {
		return err
	}
	g.router = g.router.SetBearerToken(token)

	return g.registerPermissions()
}

func (g *GenericService) create(request *generictypes.CreateObjectRequest) error {
	if request.Object.Id == "" {
		request.Object.Id = uuid.NewString()
	}
	return g.store.Create(request.Object)
}

func (g *GenericService) createMany(request *generictypes.CreateManyRequest) error {
	objectIs := []datastore.Row{}
	for _, object := range request.Objects {
		objectIs = append(objectIs, object)
	}

	return g.store.CreateMany(objectIs)
}

func (g *GenericService) upsert(request *generictypes.UpsertObjectRequest) error {
	vI, found, err := g.store.Query(
		datastore.Id(request.Object.Id),
	).FindOne()
	if err != nil {
		return err
	}

	if found {
		v := vI.(*generictypes.GenericObject)
		if v.UserId != request.Object.UserId {
			return errors.New("unauthorized")
		}
	}

	return g.store.Upsert(request.Object)
}

func (g *GenericService) upsertMany(request *generictypes.UpsertManyRequest) error {
	objectIs := []datastore.Row{}
	for _, object := range request.Objects {
		objectIs = append(objectIs, object)
	}
	return g.store.UpsertMany(objectIs)
}

func (g *GenericService) update(tableName string, userId string, conditions []datastore.Condition, object *generictypes.GenericObject) error {
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

func (g *GenericService) delete(tableName string, userId string, conditions []datastore.Condition) error {
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
