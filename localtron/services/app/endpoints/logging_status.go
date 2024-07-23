/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package appendpoints

import (
	"encoding/json"
	"net/http"

	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
)

func LoggingStatus(w http.ResponseWriter, r *http.Request, ds apptypes.AppServiceI) {
	stat, err := ds.LoggingStatus()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(stat)
	w.Write(jsonData)
}
