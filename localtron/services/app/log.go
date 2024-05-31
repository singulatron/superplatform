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
package appservice

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/singulatron/singulatron/localtron/lib"
	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
)

const (
	prod  = "https://api.commonagi.com/app/log"
	local = "http://127.0.0.1:8080/app/log"
)

func (a *AppService) manageLogs() {
	for {
		select {
		case <-a.TriggerSend:
			a.sendLogs()
		case <-a.Timer.C:
			a.sendLogs()
		}
	}
}

func (a *AppService) sendLogs() {
	a.logMutex.Lock()
	logsToSend := make([]apptypes.Log, len(a.LogBuffer))
	copy(logsToSend, a.LogBuffer)
	a.LogBuffer = nil
	a.Timer.Reset(10 * time.Second)
	a.logMutex.Unlock()

	if len(logsToSend) > 0 {
		err := a.sendToServer(logsToSend)
		if err != nil {
			lib.Logger.Info("Failed to send logs")
		}
	}
}

func (a *AppService) sendToServer(logs []apptypes.Log) error {
	jsonData, err := json.Marshal(apptypes.LogRequest{
		Logs: logs,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", prod, bytes.NewBuffer(jsonData)) // Replace with your actual URL
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server responded with status code: %d", resp.StatusCode)
	}

	return nil
}

func (a *AppService) Log(logs []apptypes.Log) error {
	a.logMutex.Lock()
	defer a.logMutex.Unlock()

	if a.clientId == "" {
		return errors.New("app service: no client id")
	}

	for i := range logs {
		logs[i].ClientId = a.clientId
		logs[i].Platform = runtime.GOOS
	}

	a.LogBuffer = append(a.LogBuffer, logs...)
	if len(a.LogBuffer) >= 100 {
		a.TriggerSend <- true
	}

	return nil
}
