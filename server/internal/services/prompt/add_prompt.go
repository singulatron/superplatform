/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptservice

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/singulatron/superplatform/sdk/go/clients/llm"
	"github.com/singulatron/superplatform/sdk/go/logger"

	apptypes "github.com/singulatron/superplatform/server/internal/services/chat/types"
	chattypes "github.com/singulatron/superplatform/server/internal/services/chat/types"
	firehosetypes "github.com/singulatron/superplatform/server/internal/services/firehose/types"
	prompttypes "github.com/singulatron/superplatform/server/internal/services/prompt/types"
)

const maxThreadTitle = 100

func (p *PromptService) addPrompt(ctx context.Context, promptReq *prompttypes.AddPromptRequest, userId string) (*prompttypes.AddPromptResponse, error) {
	prompt := &prompttypes.Prompt{
		PromptCreateFields: promptReq.PromptCreateFields,
	}

	prompt.Status = prompttypes.PromptStatusScheduled
	now := TimeNow()
	prompt.CreatedAt = now
	prompt.UpdatedAt = now
	prompt.UserId = userId

	if prompt.Id == "" {
		prompt.Id = sdk.Id("prom")
	}

	if prompt.ThreadId == "" {
		prompt.ThreadId = prompt.Id
	}

	err := p.promptsStore.Create(prompt)
	if err != nil {
		return nil, err
	}

	logger.Info("Created prompt",
		slog.String("promptId", prompt.Id),
	)

	threadId := prompt.ThreadId

	//getThreadResp := apptypes.GetThreadResponse{}
	getThreadRsp := &chattypes.GetThreadResponse{}
	err = p.router.Get(ctx, "chat-svc", fmt.Sprintf("/thread/%v", threadId), nil, &getThreadRsp)
	if err != nil {
		return nil, err
	}

	if !getThreadRsp.Exists {
		logger.Info("Creating thread", slog.String("threadId", threadId))

		// threads can be created when a message is sent
		now := time.Now()

		thread := &apptypes.Thread{
			Id:        threadId,
			UserIds:   []string{userId},
			CreatedAt: now,
			UpdatedAt: now,
		}

		if thread.Title == "" {
			if len(prompt.Prompt) > maxThreadTitle {
				thread.Title = prompt.Prompt[:maxThreadTitle]
			} else {
				thread.Title = prompt.Prompt
			}
		}

		rsp := &chattypes.AddThreadResponse{}
		err = p.router.Post(context.Background(), "chat-svc", "/thread", &chattypes.AddThreadRequest{
			Thread: thread,
		}, rsp)
		if err != nil {
			return nil, err
		}
	}

	ev := prompttypes.EventPromptAdded{
		PromptId: prompt.Id,
	}

	err = p.router.Post(context.Background(), "firehose-svc", "/event", firehosetypes.EventPublishRequest{
		Event: &firehosetypes.Event{
			Name: ev.Name(),
			Data: ev,
		},
	}, nil)
	if err != nil {
		logger.Error("Failed to publish: %v", err)
	}

	go p.triggerPromptProcessing()

	rsp := &prompttypes.AddPromptResponse{}

	if prompt.Sync {
		subscriber := make(chan *llm.CompletionResponse)
		p.StreamManager.Subscribe(threadId, subscriber)

		go func() {
			<-ctx.Done()
			p.StreamManager.Unsubscribe(threadId, subscriber)
		}()

		for resp := range subscriber {
			rsp.Answer += resp.Choices[0].Text

			if resp.Choices[0].FinishReason != "" {
				return rsp, nil
			}
		}
	}

	return rsp, nil
}

func (p *PromptService) triggerPromptProcessing() {
	select {
	case p.trigger <- true:
		logger.Debug("Prompt trigger signal sent")
	default:
		logger.Debug("Prompt trigger signal skipped, already pending")
	}
}
