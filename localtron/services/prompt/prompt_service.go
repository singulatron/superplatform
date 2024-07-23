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

	"github.com/singulatron/singulatron/localtron/clients/llm"
	"github.com/singulatron/singulatron/localtron/datastore"

	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	configtypes "github.com/singulatron/singulatron/localtron/services/config/types"
	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
	streammanager "github.com/singulatron/singulatron/localtron/services/prompt/sub/stream_manager"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

type PromptService struct {
	userService     usertypes.UserServiceI
	modelService    modeltypes.ModelServiceI
	chatService     chattypes.ChatServiceI
	firehoseService firehosetypes.FirehoseServiceI
	llmCLient       llm.ClientI

	*streammanager.StreamManager

	promptsStore datastore.DataStore

	runMutex sync.Mutex
	trigger  chan bool
}

func NewPromptService(
	cs configtypes.ConfigServiceI,
	userService usertypes.UserServiceI,
	modelService modeltypes.ModelServiceI,
	chatService chattypes.ChatServiceI,
	firehoseService firehosetypes.FirehoseServiceI,
	llmClient llm.ClientI,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*PromptService, error) {
	promptsStore, err := datastoreFactory("prompts", &prompttypes.Prompt{})
	if err != nil {
		return nil, err
	}

	service := &PromptService{
		userService:     userService,
		modelService:    modelService,
		chatService:     chatService,
		firehoseService: firehoseService,
		llmCLient:       llmClient,

		StreamManager: streammanager.NewStreamManager(),

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
