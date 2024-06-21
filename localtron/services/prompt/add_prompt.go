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
	"log/slog"
	"time"

	"github.com/pkg/errors"
	"github.com/singulatron/singulatron/localtron/logger"

	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

func (p *PromptService) AddPrompt(prompt *prompttypes.Prompt) error {
	prompt.Status = prompttypes.PromptStatusScheduled
	now := timeNow()
	prompt.CreatedAt = now
	prompt.UpdatedAt = now

	err := p.promptsStore.Create(prompt)
	if err != nil {
		return err
	}

	logger.Info("Created prompt",
		slog.String("promptId", prompt.Id),
	)

	threadId := prompt.ThreadId
	if threadId == "" {
		threadId = prompt.Id
	}

	_, threadExists, err := p.appService.GetChatThread(threadId)
	if err != nil {
		return errors.Wrap(err, "cannot get thread")
	}

	if !threadExists {
		logger.Info("Creating thread", slog.String("threadId", threadId))

		// threads can be created when a message is sent
		now := time.Now().Format(time.RFC3339Nano)

		thread := &apptypes.ChatThread{
			Id:        prompt.ThreadId,
			UserIds:   []string{prompt.UserId},
			CreatedAt: now,
			UpdatedAt: now,
		}

		if thread.Title == "" {
			if len(prompt.Prompt) > 100 {
				thread.Title = prompt.Prompt[:100]
			} else {
				thread.Title = prompt.Prompt
			}
		}

		_, err := p.appService.AddChatThread(thread)
		if err != nil {
			return errors.Wrap(err, "failed to add thread")
		}
	}

	p.firehoseService.Publish(prompttypes.EventPromptAdded{
		PromptId: prompt.Id,
	})

	go p.triggerPromptProcessing()
	return nil
}

func (p *PromptService) triggerPromptProcessing() {
	select {
	case p.trigger <- true:
		logger.Debug("Prompt trigger signal sent")
	default:
		logger.Debug("Prompt trigger signal skipped, already pending")
	}
}
