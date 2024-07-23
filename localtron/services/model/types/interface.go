/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package modeltypes

type ModelServiceI interface {
	Status(modelId string) (*ModelStatus, error)
	Start(modelId string) error
	MakeDefault(modelId string) error
	GetModels() ([]*Model, error)
}
