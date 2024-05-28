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
package configendpoints

import (
	"encoding/json"
	"net/http"

	configservice "github.com/singulatron/singulatron/localtron/services/config"
	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
)

func Get(w http.ResponseWriter, req *http.Request, cs *configservice.ConfigService) {
	conf, err := cs.GetConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(configtypes.ConfigGetResponse{
		Config: &conf,
	})
	w.Write(jsonData)
}
