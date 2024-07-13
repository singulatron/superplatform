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
package statemanager

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"os"
	"os/signal"
	"path"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/singulatron/singulatron/localtron/logger"
)

type StateFile struct {
	Rows []any `json:"rows"`
}

type StateManager struct {
	lock        sync.Mutex
	filePath    string
	hasChanged  bool
	stateGetter func() []any
}

func New(stateGetter func() []any, filePath string) *StateManager {
	sm := &StateManager{
		filePath:    filePath + ".zip",
		stateGetter: stateGetter,
	}
	sm.setupSignalHandler()
	return sm
}

func (sm *StateManager) LoadState() ([]any, error) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	_, err := os.Stat(sm.filePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(sm.filePath), 0755)
		if err != nil {
			return nil, err
		}
		emptyData := []byte("{}")
		zippedEmptyData, err := zipData(sm.filePath, emptyData)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(sm.filePath, zippedEmptyData, 0644)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(sm.filePath)
	if err != nil {
		return nil, err
	}

	unzippedData, err := unzipData(data)
	if err != nil {
		return nil, err
	}

	var stateFile StateFile
	if strings.TrimSpace(string(unzippedData)) == "" {
		return nil, nil
	}
	err = json.Unmarshal(unzippedData, &stateFile)
	if err != nil {
		return nil, err
	}

	return stateFile.Rows, nil
}

func (sm *StateManager) SaveState(shallowCopy []any) error {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	data, err := json.Marshal(&StateFile{
		Rows: shallowCopy,
	})
	if err != nil {
		return err
	}

	zippedData, err := zipData(sm.filePath, data)
	if err != nil {
		return err
	}

	tempFilePath := sm.filePath + ".tmp"
	err = ioutil.WriteFile(tempFilePath, zippedData, 0644)
	if err != nil {
		return err
	}

	finalFilePath := sm.filePath
	err = os.Rename(tempFilePath, finalFilePath)
	if err != nil {
		return err
	}

	sm.hasChanged = false
	return nil
}

func (sm *StateManager) MarkChanged() {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	sm.hasChanged = true
}

func (sm *StateManager) PeriodicSaveState(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if sm.hasChanged {
				if err := sm.SaveState(sm.stateGetter()); err != nil {
					logger.Logger.Error("Error saving file state",
						slog.String("filePath", sm.filePath),
						slog.String("error", err.Error()),
					)
				}
			}
		}
	}
}

func (sm *StateManager) setupSignalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		err := sm.SaveState(sm.stateGetter())
		if err != nil {
			logger.Logger.Error("Error saving file state on shutdown",
				slog.String("filePath", sm.filePath),
				slog.String("error", err.Error()),
			)
		}
		os.Exit(0)
	}()
}

func zipData(filePath string, data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)
	writer, err := zipWriter.Create(
		fmt.Sprintf("%v.json", path.Base(filePath)),
	)
	if err != nil {
		return nil, err
	}
	_, err = writer.Write(data)
	if err != nil {
		return nil, err
	}
	err = zipWriter.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func unzipData(data []byte) ([]byte, error) {
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}
	if len(reader.File) != 1 {
		return nil, os.ErrInvalid
	}
	file := reader.File[0]
	rc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	unzippedData, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return unzippedData, nil
}
