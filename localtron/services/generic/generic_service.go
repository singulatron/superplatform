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

	"github.com/singulatron/singulatron/localtron/datastore"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	userservice "github.com/singulatron/singulatron/localtron/services/user"

	generictypes "github.com/singulatron/singulatron/localtron/services/generic/types"
)

type GenericService struct {
	configService   *configservice.ConfigService
	userService     *userservice.UserService
	firehoseService *firehoseservice.FirehoseService

	store datastore.DataStore
}

func NewGenericService(
	cs *configservice.ConfigService,
	fs *firehoseservice.FirehoseService,
	userService *userservice.UserService,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*GenericService, error) {
	store, err := datastoreFactory("generic", &generictypes.GenericObject{})
	if err != nil {
		return nil, err
	}

	service := &GenericService{
		configService:   cs,
		firehoseService: fs,
		userService:     userService,
		store:           store,
	}

	err = service.registerPermissions()
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (g *GenericService) Create(object *generictypes.GenericObject) error {
	return g.store.Create(object)
}

func (g *GenericService) CreateMany(objects []*generictypes.GenericObject) error {
	objectIs := []datastore.Row{}
	for _, object := range objects {
		objectIs = append(objectIs, object)
	}

	return g.store.CreateMany(objectIs)
}

func (g *GenericService) Upsert(object *generictypes.GenericObject) error {
	vI, found, err := g.store.Query(
		datastore.Id(object.Id),
	).FindOne()
	if err != nil {
		return err
	}

	if found {
		v := vI.(*generictypes.GenericObject)
		if v.UserId != object.UserId {
			return errors.New("unauthorized")
		}
	}

	return g.store.Upsert(object)
}

func (g *GenericService) UpsertMany(objects []*generictypes.GenericObject) error {
	objectIs := []datastore.Row{}
	for _, object := range objects {
		objectIs = append(objectIs, object)
	}
	return g.store.UpsertMany(objectIs)
}

func (g *GenericService) Update(tableName string, userId string, conditions []datastore.Condition, object *generictypes.GenericObject) error {
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

func (g *GenericService) Delete(tableName string, userId string, conditions []datastore.Condition) error {
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
