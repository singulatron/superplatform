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

	appservice "github.com/singulatron/singulatron/localtron/services/app"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	modelservice "github.com/singulatron/singulatron/localtron/services/model"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

type PromptService struct {
	modelService    *modelservice.ModelService
	appService      *appservice.AppService
	firehoseService *firehoseservice.FirehoseService

	StreamManager *StreamManager

	currentPrompt      *prompttypes.Prompt
	currentPromptMutex sync.Mutex

	promptsToProcess      []*prompttypes.Prompt
	promptsToProcessMutex sync.Mutex

	trigger chan bool
}

func NewPromptService(
	modelService *modelservice.ModelService,
	appService *appservice.AppService,
	firehoseService *firehoseservice.FirehoseService,
) *PromptService {
	service := &PromptService{
		modelService:    modelService,
		appService:      appService,
		firehoseService: firehoseService,

		StreamManager: NewStreamManager(),

		promptsToProcess: []*prompttypes.Prompt{},
		trigger:          make(chan bool, 1),
	}
	go service.processPrompts()
	return service
}
