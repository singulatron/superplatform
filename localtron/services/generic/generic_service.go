/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package genericservice

import (
	"errors"

	"github.com/singulatron/singulatron/localtron/datastore"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	storefactoryservice "github.com/singulatron/singulatron/localtron/services/store_factory"
	userservice "github.com/singulatron/singulatron/localtron/services/user"

	generictypes "github.com/singulatron/singulatron/localtron/services/generic/types"
)

type GenericService struct {
	configService   *configservice.ConfigService
	userService     *userservice.UserService
	firehoseService *firehoseservice.FirehoseService

	store datastore.DataStore[*generictypes.GenericObject]
}

func NewGenericService(
	cs *configservice.ConfigService,
	fs *firehoseservice.FirehoseService,
	userService *userservice.UserService,
) (*GenericService, error) {
	store, err := storefactoryservice.GetStore[*generictypes.GenericObject]("generic")
	if err != nil {
		return nil, err
	}

	service := &GenericService{
		configService:   cs,
		firehoseService: fs,
		userService:     userService,

		store: store,
	}

	err = service.registerPermissions()
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (g *GenericService) Create(tableName string, userId string, object *generictypes.GenericObject) error {
	object.Table = tableName
	object.UserId = userId
	return g.store.Create(object)
}

func (g *GenericService) CreateMany(tableName string, userId string, objects []*generictypes.GenericObject) error {
	for _, object := range objects {
		object.Table = tableName
		object.UserId = userId
	}
	return g.store.CreateMany(objects)
}

func (g *GenericService) Upsert(tableName string, userId string, object *generictypes.GenericObject) error {
	v, found, err := g.store.Query(
		datastore.Id(object.Id),
	).FindOne()
	if err != nil {
		return err
	}
	if found && v.UserId != userId {
		return errors.New("unauthorized")
	}
	object.Table = tableName
	object.UserId = userId
	return g.store.Upsert(object)
}

func (g *GenericService) UpsertMany(tableName string, userId string, objects []*generictypes.GenericObject) error {
	for _, object := range objects {
		object.Table = tableName
	}
	return g.store.UpsertMany(objects)
}

func (g *GenericService) Find(tableName string, userId string, conditions []datastore.Condition) ([]*generictypes.GenericObject, error) {
	if len(conditions) == 0 {
		return nil, errors.New("no conditions")
	}

	conditions = append(conditions, datastore.Equal(datastore.Field("table"), tableName))

	return g.store.Query(
		conditions[0], conditions[1:]...,
	).Find()
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
