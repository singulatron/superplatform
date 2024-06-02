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
package downloadservice

import (
	"encoding/json"
	"log/slog"
	"os"
	"path"
	"sync"
	"time"

	"github.com/singulatron/singulatron/localtron/lib"
	types "github.com/singulatron/singulatron/localtron/services/download/types"
)

type DownloadService struct {
	downloads     map[string]*types.Download
	lock          sync.Mutex
	StateFilePath string
	DefaultFolder string
	hasChanged    bool
}

func NewDownloadService() (*DownloadService, error) {
	home, _ := os.UserHomeDir()
	ret := &DownloadService{
		StateFilePath: path.Join(home, "singulatron_downloads.json"),
		downloads:     make(map[string]*types.Download),
	}

	return ret, nil
}

func (dm *DownloadService) Start() error {
	err := dm.loadState()
	if err != nil {
		return err
	}

	for _, download := range dm.downloads {
		if download.Status == types.DownloadStatusInProgress {
			err = dm.Do(download.URL, path.Dir(download.FilePath))
			if err != nil {
				return err
			}
		}
	}

	go dm.periodicSaveState()

	return err
}

func (dm *DownloadService) loadState() error {
	dm.lock.Lock()
	defer dm.lock.Unlock()

	_, err := os.Stat(dm.StateFilePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(dm.StateFilePath), 0755)
		if err != nil {
			return err
		}
		err = os.WriteFile(dm.StateFilePath, []byte("{}"), 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	data, err := os.ReadFile(dm.StateFilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &dm.downloads)
}

func (ds *DownloadService) markChanged() {
	ds.lock.Lock()
	defer ds.lock.Unlock()
	ds.hasChanged = true
}

func (ds *DownloadService) markChangedWithoutLock() {
	ds.hasChanged = true
}

func (ds *DownloadService) saveState() error {
	ds.lock.Lock()
	data, err := json.MarshalIndent(ds.downloads, "", "  ")
	if err != nil {
		ds.lock.Unlock()
		return err
	}
	ds.hasChanged = false
	ds.lock.Unlock()

	err = os.WriteFile(ds.StateFilePath, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DownloadService) periodicSaveState() {
	for {
		time.Sleep(1 * time.Second) // Control the throttle rate here
		ds.lock.Lock()
		if ds.hasChanged {
			ds.lock.Unlock()
			if err := ds.saveState(); err != nil {
				lib.Logger.Error("Failed to save state", slog.String("error", err.Error()))
			}
		} else {
			ds.lock.Unlock()
		}
	}
}

func (dm *DownloadService) GetDownload(url string) (*types.Download, bool) {
	dm.lock.Lock()
	defer dm.lock.Unlock()

	v, ok := dm.downloads[url]
	return v, ok
}

//
// Event
//

const EventDownloadStatusChangeName = "downloadStatusChange"

type EventDownloadStatusChange struct {
}

func (e EventDownloadStatusChange) Name() string {
	return EventDownloadStatusChangeName
}
