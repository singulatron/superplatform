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
	"encoding/json"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
)

type AppService struct {
	configService   *configservice.ConfigService
	firehoseService *firehoseservice.FirehoseService

	LogBuffer   []apptypes.Log
	TriggerSend chan bool
	Timer       *time.Timer

	clientId     string
	ChatFilePath string
	chatFile     *apptypes.ChatFile

	chatFileMutex sync.Mutex
	logMutex      sync.Mutex
}

func NewAppService(
	cs *configservice.ConfigService,
	fs *firehoseservice.FirehoseService,
) (*AppService, error) {
	ci, err := cs.GetClientId()
	if err != nil {
		return nil, errors.Wrap(err, "app service canno get client id")
	}
	service := &AppService{
		configService:   cs,
		firehoseService: fs,

		LogBuffer:   make([]apptypes.Log, 0),
		TriggerSend: make(chan bool, 1),
		Timer:       time.NewTimer(10 * time.Second),

		clientId: ci,
	}
	service.ChatFilePath = path.Join(cs.ConfigDirectory, "chats.json")
	err = service.loadChatFile()
	if err != nil {
		return nil, err
	}
	go service.manageLogs()
	service.setupSignalHandler()
	return service, nil
}

func (a *AppService) setupSignalHandler() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		a.sendLogs()
		os.Exit(0)
	}()
}

func (a *AppService) loadChatFile() error {
	a.chatFileMutex.Lock()
	defer a.chatFileMutex.Unlock()

	_, err := os.Stat(a.ChatFilePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(a.ChatFilePath), 0755)
		if err != nil {
			return err
		}
		err = os.WriteFile(a.ChatFilePath, []byte(`{}`), 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	data, err := os.ReadFile(a.ChatFilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &a.chatFile)
}

func (a *AppService) saveChatFile() error {
	a.chatFileMutex.Lock()
	defer a.chatFileMutex.Unlock()

	bs, err := json.Marshal(a.chatFile)
	if err != nil {
		return err
	}

	err = os.WriteFile(a.ChatFilePath, bs, 0755)
	return err
}
