/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package appendpoints

import (
	"encoding/json"
	"net/http"

	appservice "github.com/singulatron/singulatron/localtron/services/app"
)

func LoggingStatus(w http.ResponseWriter, r *http.Request, ds *appservice.AppService) {
	stat, err := ds.LoggingStatus()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(stat)
	w.Write(jsonData)
}
