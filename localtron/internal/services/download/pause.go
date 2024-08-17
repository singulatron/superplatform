/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package downloadservice

import (
	"fmt"

	downloadtypes "github.com/singulatron/singulatron/localtron/internal/services/download/types"
)

/*
Pauses a download.
*/
func (ds *DownloadService) pause(url string) error {
	ds.lock.Lock()
	defer ds.lock.Unlock()

	d, exists := ds.downloads[url]
	if !exists {
		return fmt.Errorf("url '%v' is not being downloaded", url)
	}

	d.Status = downloadtypes.DownloadStatusPaused
	ds.markChangedWithoutLock()

	return nil
}
