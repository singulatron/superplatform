/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package generictypes

import "github.com/singulatron/singulatron/localtron/datastore"

type GenericServiceI interface {
	Find(options FindOptions) ([]*GenericObject, error)
	Create(object *GenericObjectCreateFields) error
	CreateMany(objects []*GenericObjectCreateFields) error
	Upsert(object *GenericObjectCreateFields) error
	UpsertMany(objects []*GenericObjectCreateFields) error
	Update(tableName string, userId string, conditions []datastore.Condition, object *GenericObject) error
	Delete(tableName string, userId string, conditions []datastore.Condition) error
}
