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
	"github.com/singulatron/singulatron/localtron/router"

	streammanager "github.com/singulatron/singulatron/localtron/services/prompt/sub/stream_manager"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

type PromptService struct {
	llmCLient llm.ClientI
	router    *router.Router

	*streammanager.StreamManager

	promptsStore    datastore.DataStore
	credentialStore datastore.DataStore

	runMutex sync.Mutex
	trigger  chan bool
}

func NewPromptService(
	router *router.Router,
	llmClient llm.ClientI,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*PromptService, error) {
	promptsStore, err := datastoreFactory("prompts", &prompttypes.Prompt{})
	if err != nil {
		return nil, err
	}

	credentialStore, err := datastoreFactory("prompt_credentials", &usertypes.Credential{})
	if err != nil {
		return nil, err
	}

	service := &PromptService{
		llmCLient: llmClient,
		router:    router,

		StreamManager: streammanager.NewStreamManager(),

		promptsStore:    promptsStore,
		credentialStore: credentialStore,

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

	go service.processPrompts()

	return service, nil
}

func (cs *PromptService) Start() error {
	token, err := usertypes.RegisterService("prompt-service", "Prompt Service", cs.router, cs.credentialStore)
	if err != nil {
		return err
	}
	cs.router = cs.router.SetBearerToken(token)

	return cs.registerPermissions()
}
