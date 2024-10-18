/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package downloadservice

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"path"
	"sync"
	"time"

	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/datastore"
	"github.com/singulatron/superplatform/sdk/go/lock"
	"github.com/singulatron/superplatform/sdk/go/logger"
	"github.com/singulatron/superplatform/sdk/go/router"
	types "github.com/singulatron/superplatform/server/internal/services/download/types"
	firehosetypes "github.com/singulatron/superplatform/server/internal/services/firehose/types"
)

type DownloadService struct {
	router *router.Router
	dlock  lock.DistributedLock

	downloads map[string]*types.Download
	lock      sync.Mutex

	StateFilePath string
	DefaultFolder string
	hasChanged    bool

	// for testing purposes
	SyncDownloads bool

	credentialStore datastore.DataStore
}

func NewDownloadService(
	router *router.Router,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*DownloadService, error) {
	home, _ := os.UserHomeDir()

	credentialStore, err := datastoreFactory("downloadSvcCredentials", &sdk.Credential{})
	if err != nil {
		return nil, err
	}

	ret := &DownloadService{
		credentialStore: credentialStore,
		router:          router,
		dlock:           lock,

		StateFilePath: path.Join(home, "downloads.json"),
		downloads:     make(map[string]*types.Download),
	}

	return ret, nil
}

func (dm *DownloadService) SetDefaultFolder(s string) {
	dm.DefaultFolder = s
}

func (dm *DownloadService) SetStateFilePath(s string) {
	dm.StateFilePath = s
}

func (dm *DownloadService) Start() error {
	ctx := context.Background()
	dm.dlock.Acquire(ctx, "download-svc-start")
	defer dm.dlock.Release(ctx, "download-svc-start")

	token, err := sdk.RegisterService("download-svc", "Download Service", dm.router, dm.credentialStore)
	if err != nil {
		return err
	}
	dm.router = dm.router.SetBearerToken(token)

	err = dm.registerPermissions()
	if err != nil {
		return err
	}

	err = dm.loadState()
	if err != nil {
		return err
	}

	for _, download := range dm.downloads {
		if download.Status == types.DownloadStatusInProgress {
			err = dm.do(download.URL, path.Dir(download.FilePath))
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

	ds.router.Post(context.Background(), "firehose-svc", "/event", firehosetypes.EventPublishRequest{
		Event: &firehosetypes.Event{
			Name: types.EventDownloadStatusChangeName,
		},
	}, nil)

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
				logger.Error("Failed to save state", slog.String("error", err.Error()))
			}
		} else {
			ds.lock.Unlock()
		}
	}
}

func (dm *DownloadService) getDownload(url string) (*types.Download, bool) {
	dm.lock.Lock()
	defer dm.lock.Unlock()

	v, ok := dm.downloads[url]
	return v, ok
}
