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
package modelservice

import "fmt"

func (ms *ModelService) MakeDefault(modelId string) error {
	stat, err := ms.Status(modelId)
	if err != nil {
		return err
	}
	if !stat.SelectedExists {
		return fmt.Errorf("cannot set model as it is not downloaded yet")
	}

	conf, err := ms.configService.GetConfig()
	if err != nil {
		return err
	}
	conf.Model.CurrentModelId = modelId
	return ms.configService.SaveConfig(conf)
}
