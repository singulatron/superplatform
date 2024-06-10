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
package lib

import (
	"encoding/json"
	"io/ioutil"
	"log/slog"
	"os"
	"os/signal"
	"path"
	"strings"
	"sync"
	"syscall"
	"time"
)

type StateFile[T Row] struct {
	Rows []T `json:"rows"`
}

type StateManager[T Row] struct {
	key        string
	memStore   *MemoryStore[T]
	lock       sync.Mutex
	filePath   string
	hasChanged bool
}

func NewStateManager[T Row](memStore *MemoryStore[T], filePath string) *StateManager[T] {
	sm := &StateManager[T]{
		memStore: memStore,
		filePath: filePath,
	}
	sm.setupSignalHandler()
	return sm
}

func (sm *StateManager[T]) LoadState() error {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	_, err := os.Stat(sm.filePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(sm.filePath), 0755)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(sm.filePath, []byte("{}"), 0644)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(sm.filePath)
	if err != nil {
		return err
	}

	var stateFile StateFile[T]
	if strings.TrimSpace(string(data)) == "" {
		return nil
	}
	err = json.Unmarshal(data, &stateFile)
	if err != nil {
		return err
	}

	sm.memStore.Reset(stateFile.Rows)

	return nil
}

func (sm *StateManager[T]) SaveState() error {
	shallowCopy := sm.memStore.SliceCopy()

	sm.lock.Lock()
	defer sm.lock.Unlock()

	sm.memStore.Lock()
	data, err := json.Marshal(&StateFile[T]{
		Rows: shallowCopy,
	})
	sm.memStore.Unlock()
	if err != nil {
		return err
	}
	sm.hasChanged = false

	tempFilePath := sm.filePath + ".tmp"
	err = ioutil.WriteFile(tempFilePath, data, 0644)
	if err != nil {
		return err
	}

	err = os.Rename(tempFilePath, sm.filePath)
	if err != nil {
		return err
	}

	return nil
}

func (sm *StateManager[T]) MarkChanged() {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	sm.hasChanged = true
}

func (sm *StateManager[T]) PeriodicSaveState(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if sm.hasChanged {
				if err := sm.SaveState(); err != nil {
					Logger.Error("Error saving file state",
						slog.String("filePath", sm.filePath),
						slog.String("error", err.Error()),
					)
				}
			}
		}
	}
}

func (sm *StateManager[T]) setupSignalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		err := sm.SaveState()
		if err != nil {
			Logger.Error("Error saving file state on shutdown",
				slog.String("filePath", sm.filePath),
				slog.String("error", err.Error()),
			)
		}
		os.Exit(0)
	}()
}
