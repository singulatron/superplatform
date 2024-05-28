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
package dockerendpoints

import (
	"encoding/json"
	"net/http"

	dockerservice "github.com/singulatron/singulatron/localtron/services/docker"
)

func Info(w http.ResponseWriter, req *http.Request, dm *dockerservice.DockerService) {
	di, err := dm.Info()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{
		"info": di,
	})
	w.Write(jsonData)
}
