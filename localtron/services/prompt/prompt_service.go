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
package promptservice

import (
	"path"
	"sync"
	"time"

	"github.com/singulatron/singulatron/localtron/lib"
	appservice "github.com/singulatron/singulatron/localtron/services/app"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	modelservice "github.com/singulatron/singulatron/localtron/services/model"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

type PromptService struct {
	userService     *userservice.UserService
	modelService    *modelservice.ModelService
	appService      *appservice.AppService
	firehoseService *firehoseservice.FirehoseService

	PromptsFilePath string
	StreamManager   *StreamManager

	promptsMem  *lib.MemoryStore[*prompttypes.Prompt]
	promptsFile *lib.StateManager[*prompttypes.Prompt]

	runMutex sync.Mutex
	trigger  chan bool
}

func NewPromptService(
	cs *configservice.ConfigService,
	userService *userservice.UserService,
	modelService *modelservice.ModelService,
	appService *appservice.AppService,
	firehoseService *firehoseservice.FirehoseService,
) (*PromptService, error) {

	promptsPath := path.Join(cs.ConfigDirectory, "data", "prompts")
	pm := lib.NewMemoryStore[*prompttypes.Prompt]()

	service := &PromptService{
		userService:     userService,
		modelService:    modelService,
		appService:      appService,
		firehoseService: firehoseService,

		PromptsFilePath: promptsPath,
		StreamManager:   NewStreamManager(),

		promptsMem:  pm,
		promptsFile: lib.NewStateManager[*prompttypes.Prompt](pm, promptsPath),

		trigger: make(chan bool, 1),
	}

	err := service.promptsFile.LoadState()
	if err != nil {
		return nil, err
	}

	service.promptsMem.Foreach(func(i int, item *prompttypes.Prompt) {
		if item.Status == prompttypes.PromptStatusRunning {
			item.Status = prompttypes.PromptStatusScheduled
		}
	})

	service.registerPermissions()
	service.promptsFile.MarkChanged()

	go service.processPrompts()
	go service.promptsFile.PeriodicSaveState(2 * time.Second)

	return service, nil
}
