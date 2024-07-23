/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
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
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

type PromptService struct {
	userService     *userservice.UserService
	modelService    *modelservice.ModelService
	appService      *chatservice.ChatService
	firehoseService *firehoseservice.FirehoseService

	*StreamManager

	promptsStore datastore.DataStore

	runMutex sync.Mutex
	trigger  chan bool
}

func NewPromptService(
	cs *configservice.ConfigService,
	userService *userservice.UserService,
	modelService *modelservice.ModelService,
	appService *chatservice.ChatService,
	firehoseService *firehoseservice.FirehoseService,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*PromptService, error) {
	promptsStore, err := datastoreFactory("prompts", &prompttypes.Prompt{})
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

	promptIs, err := service.promptsStore.Query(
		datastore.Equal(datastore.Field("status"), prompttypes.PromptStatusRunning),
	).Find()
	if err != nil {
		return nil, err
	}
	promptIds := []string{}
	for _, promptI := range promptIs {
		promptIds = append(promptIds, promptI.(*prompttypes.Prompt).Id)
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
