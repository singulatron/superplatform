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
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/singulatron/singulatron/localtron/lib"
	"github.com/singulatron/singulatron/localtron/llm"
	appservice "github.com/singulatron/singulatron/localtron/services/app"
	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
	modelservice "github.com/singulatron/singulatron/localtron/services/model"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

type PromptService struct {
	modelService *modelservice.ModelService
	appService   *appservice.AppService

	StreamManager *StreamManager

	currentPrompt    *prompttypes.Prompt
	promptsToProcess []*prompttypes.Prompt
	mutex            sync.Mutex
	trigger          chan bool
}

func NewPromptService(modelService *modelservice.ModelService, appService *appservice.AppService) *PromptService {
	service := &PromptService{
		modelService: modelService,
		appService:   appService,

		StreamManager: NewStreamManager(),

		promptsToProcess: []*prompttypes.Prompt{},
		trigger:          make(chan bool, 1),
	}
	go service.processPrompts()
	return service
}

func (p *PromptService) processPrompts() {
	ticker := time.NewTicker(2000 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		case <-p.trigger: // Listen for immediate processing signals
		}

		p.mutex.Lock()
		if p.currentPrompt == nil && len(p.promptsToProcess) > 0 {
			p.currentPrompt = p.promptsToProcess[0]
			p.currentPrompt.IsBeingProcessed = true
			p.promptsToProcess = p.promptsToProcess[1:]
			lib.Logger.Info("Picking up prompt from queue", slog.String("promptId", p.currentPrompt.Id))

			go func() {
				if r := recover(); r != nil {
					lib.Logger.Error("Recovered from prompt process")
				}
				err := p.processPromptWrapper()
				if err != nil {
					lib.Logger.Error("Prompt process errored", slog.String("error", err.Error()))
				}
			}()

		}
		p.mutex.Unlock()
	}
}

func (p *PromptService) TriggerPromptProcessing() {
	select {
	case p.trigger <- true:
		lib.Logger.Debug("Prompt trigger signal sent")
	default:
		lib.Logger.Debug("Prompt trigger signal skipped, already pending")
	}
}

func (p *PromptService) processPromptWrapper() error {
	err := p.processPrompt()
	if err != nil {
		lib.Logger.Error("Prompt process errored, putting prompt back to queue",
			slog.String("error", err.Error()),
		)
		// put the prompt back to the queue
		p.mutex.Lock()
		p.currentPrompt.IsBeingProcessed = false
		p.promptsToProcess = append([]*prompttypes.Prompt{p.currentPrompt}, p.promptsToProcess...)
		p.currentPrompt = nil
		p.mutex.Unlock()
	}
	p.mutex.Lock()
	p.currentPrompt = nil
	p.mutex.Unlock()
	return nil
}

func (p *PromptService) processPrompt() error {
	// @todo make this idempotent - on failures and retries
	// a bunch of messages will be generated...
	err := p.appService.AddChatMessage(&apptypes.ChatMessage{
		Id:             uuid.New().String(),
		ThreadId:       p.currentPrompt.ThreadId,
		IsUserMessage:  true,
		MessageContent: p.currentPrompt.Message,
		Time:           time.Now().Format(time.RFC3339),
	})
	if err != nil {
		return err
	}

	stat, err := p.modelService.Status(p.currentPrompt.ModelId)
	if err != nil {
		return errors.Wrap(err, "error getting model status")
	}
	if !stat.Running {
		return fmt.Errorf("model '%v' is not running", p.currentPrompt.ModelId)
	}
	if stat.ModelAddress == "" {
		return errors.Wrap(err, "missing model address")
	}
	if !strings.HasPrefix(stat.ModelAddress, "http") {
		stat.ModelAddress = "http://" + stat.ModelAddress
	}
	llmClient := llm.Client{
		LLMAddress: stat.ModelAddress,
	}

	err = llmClient.PostCompletionsStreamed(llm.PostCompletionsRequest{
		Prompt:    p.currentPrompt.Prompt,
		Stream:    true,
		MaxTokens: 200,
	}, func(resp *llm.CompletionResponse) {
		p.StreamManager.Broadcast(p.currentPrompt.ThreadId, resp)
		if len(resp.Choices) > 0 && resp.Choices[0].FinishReason == "stop" {
			err := p.appService.AddChatMessage(&apptypes.ChatMessage{
				Id:             uuid.New().String(),
				ThreadId:       p.currentPrompt.ThreadId,
				MessageContent: llmResponseToText(p.StreamManager.history[p.currentPrompt.ThreadId]),
			})
			if err != nil {
				lib.Logger.Error("Error when saving chat message after broadcast",
					slog.String("error", err.Error()))
				return
			}

			delete(p.StreamManager.history, p.currentPrompt.ThreadId)
		}
	})
	if err != nil {
		return errors.Wrap(err, "error streaming llm")
	}

	return nil
}

func llmResponseToText(responses []*llm.CompletionResponse) string {
	var result strings.Builder

	first := true
	for _, v := range responses {
		if len(v.Choices) == 0 {
			continue
		}
		choice := v.Choices[0]

		var textToAdd string
		if strings.Contains(result.String(), "```") {
			// Handling for inline code formatting if the resulting string is already within a code block
			count := strings.Count(result.String(), "```")
			if count%2 == 1 { // If the count of ``` is odd, we are inside a code block
				textToAdd = choice.Text // No escaping needed inside code block
			} else {
				textToAdd = escapeHtml(choice.Text) // Apply HTML escaping when outside code blocks
			}
		} else {
			textToAdd = escapeHtml(choice.Text) // Apply HTML escaping if there is no code block
		}

		if first {
			textToAdd = strings.TrimLeft(textToAdd, " ")
			first = false
		}

		result.WriteString(textToAdd)

		if choice.FinishReason == "stop" {
			break
		}
	}

	return result.String()
}

func escapeHtml(input string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&#39;",
	)
	return replacer.Replace(input)
}
