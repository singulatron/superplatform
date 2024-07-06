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
	"sync"

	"github.com/singulatron/singulatron/localtron/datastore"

	chatservice "github.com/singulatron/singulatron/localtron/services/chat"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	modelservice "github.com/singulatron/singulatron/localtron/services/model"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	storefactoryservice "github.com/singulatron/singulatron/localtron/services/store_factory"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

type PromptService struct {
	userService     *userservice.UserService
	modelService    *modelservice.ModelService
	appService      *chatservice.ChatService
	firehoseService *firehoseservice.FirehoseService

	StreamManager *StreamManager

	promptsStore datastore.DataStore[*prompttypes.Prompt]

	runMutex sync.Mutex
	trigger  chan bool
}

func NewPromptService(
	cs *configservice.ConfigService,
	userService *userservice.UserService,
	modelService *modelservice.ModelService,
	appService *chatservice.ChatService,
	firehoseService *firehoseservice.FirehoseService,

) (*PromptService, error) {
	promptsStore, err := storefactoryservice.GetStore[*prompttypes.Prompt]("prompts")
	if err != nil {
		return nil, err
	}

	service := &PromptService{
		userService:     userService,
		modelService:    modelService,
		appService:      appService,
		firehoseService: firehoseService,

		StreamManager: NewStreamManager(),

		promptsStore: promptsStore,

		trigger: make(chan bool, 1),
	}

	prompts, err := service.promptsStore.Query(
		datastore.Equal(datastore.Field("status"), prompttypes.PromptStatusRunning),
	).Find()
	if err != nil {
		return nil, err
	}
	promptIds := []string{}
	for _, prompt := range prompts {
		promptIds = append(promptIds, prompt.Id)
	}

	err = service.promptsStore.Query(
		datastore.Equal(datastore.Field("id"), promptIds),
	).UpdateFields(map[string]any{
		"status": prompttypes.PromptStatusScheduled,
	})
	if err != nil {
		return nil, err
	}

	service.registerPermissions()

	go service.processPrompts()

	return service, nil
}
