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

	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/llm"
	"github.com/singulatron/singulatron/localtron/logger"

	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

const (
	maxRetries    = 5
	baseDelay     = 1 * time.Second
	promptTimeout = 1 * time.Minute
)

// a blocking method, call it in a goroutine
func (p *PromptService) processPrompts() {
	ticker := time.NewTicker(2000 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		case <-p.trigger:
		}

		err := p.processNextPrompt()
		if err != nil {
			logger.Error("Error processing prompt",
				slog.String("error", err.Error()),
			)
		}
	}
}

func (p *PromptService) processNextPrompt() error {
	p.runMutex.Lock()
	defer p.runMutex.Unlock()

	runningPrompts, err := p.promptsStore.Query(
		datastore.Equal("status", prompttypes.PromptStatusRunning),
	).Find()
	if err != nil {
		return err
	}

	hasRunning := false
	runningPromptId := ""
	for _, runningPrompt := range runningPrompts {
		if runningPrompt.LastRun.Before(time.Now().Add(-promptTimeout)) {
			logger.Info("Setting prompt as timed out",
				slog.String("promptId", runningPrompt.Id),
			)

			runningPrompt.Status = prompttypes.PromptStatusErrored
			runningPrompt.Error = "timed out"
			err = p.promptsStore.Query(
				datastore.Id(runningPrompt.Id),
			).Update(runningPrompt)
			if err != nil {
				return err
			}
			continue
		}
		hasRunning = true
		runningPromptId = runningPrompt.Id
	}

	if hasRunning {
		logger.Debug("Prompt is already running",
			slog.String("promptId", runningPromptId),
		)
		return nil
	}

	currentPrompt, err := selectPrompt(p.promptsStore)
	if err != nil {
		return err
	}
	if currentPrompt == nil {
		return nil
	}

	return p.processPrompt(currentPrompt)
}

func (p *PromptService) processPrompt(currentPrompt *prompttypes.Prompt) (err error) {
	defer func() {
		if r := recover(); r != nil {
			currentPrompt.Error = fmt.Sprintf("%v", r)
			currentPrompt.Status = prompttypes.PromptStatusErrored
			return
		}

		if err != nil {
			currentPrompt.Error = err.Error()
			currentPrompt.Status = prompttypes.PromptStatusErrored
		} else {
			currentPrompt.Status = prompttypes.PromptStatusCompleted
		}

		logger.Info("Prompt finished",
			slog.String("promptId", currentPrompt.Id),
			slog.String("status", string(currentPrompt.Status)),
		)

		err = p.promptsStore.Query(
			datastore.Id(currentPrompt.Id),
		).Update(currentPrompt)
		if err != nil {
			logger.Error("Error updating prompt",
				slog.String("promptId", currentPrompt.Id),
				slog.String("error", err.Error()),
			)
		}
	}()

	logger.Info("Picking up prompt from queue",
		slog.String("promptId", currentPrompt.Id),
	)

	p.firehoseService.Publish(prompttypes.EventPromptProcessingStarted{
		PromptId: currentPrompt.Id,
	})

	defer p.firehoseService.Publish(prompttypes.EventPromptProcessingFinished{
		PromptId: currentPrompt.Id,
		Error:    errToString(err),
	})

	currentPrompt.LastRun = time.Now()
	currentPrompt.Error = ""
	currentPrompt.Status = prompttypes.PromptStatusRunning
	currentPrompt.RunCount++

	err = p.appService.AddChatMessage(&apptypes.ChatMessage{
		// not a fan of taking the prompt id but at least it makes this idempotent
		// in case prompts get retried over and over again
		Id:        currentPrompt.Id,
		ThreadId:  currentPrompt.ThreadId,
		UserId:    currentPrompt.UserId,
		Content:   currentPrompt.Prompt,
		CreatedAt: time.Now().Format(time.RFC3339),
	})
	if err != nil {
		return err
	}

	stat, err := p.modelService.Status(currentPrompt.ModelId)
	if err != nil {
		return errors.Wrap(err, "error getting model status")
	}
	if !stat.Running {
		return fmt.Errorf("model '%v' is not running", currentPrompt.ModelId)
	}
	if stat.Address == "" {
		return errors.Wrap(err, "missing model address")
	}
	if !strings.HasPrefix(stat.Address, "http") {
		stat.Address = "http://" + stat.Address
	}
	llmClient := llm.Client{
		LLMAddress: stat.Address,
	}

	fullPrompt := currentPrompt.Prompt
	if currentPrompt.Template != "" {
		fullPrompt = strings.Replace(currentPrompt.Template, "{prompt}", currentPrompt.Prompt, -1)
	}

	start := time.Now()
	var responseCount int
	var mu sync.Mutex

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				mu.Lock()
				logger.Debug("LLM is streaming",
					slog.String("promptId", currentPrompt.Id),
					slog.Float64("responsesPerSecond", float64(responseCount/int(time.Since(start).Seconds()))),
					slog.Int("totalResponses", responseCount),
				)
				mu.Unlock()
			case <-done:
				return
			}
		}
	}()

	err = llmClient.PostCompletionsStreamed(llm.PostCompletionsRequest{
		Prompt:    fullPrompt,
		Stream:    true,
		MaxTokens: 1000000,
	}, func(resp *llm.CompletionResponse) {
		mu.Lock()
		responseCount++
		mu.Unlock()

		p.StreamManager.Broadcast(currentPrompt.ThreadId, resp)

		if len(resp.Choices) > 0 && resp.Choices[0].FinishReason == "stop" {
			err := p.appService.AddChatMessage(&apptypes.ChatMessage{
				Id:       uuid.New().String(),
				ThreadId: currentPrompt.ThreadId,
				Content:  llmResponseToText(p.StreamManager.history[currentPrompt.ThreadId]),
			})
			if err != nil {
				logger.Error("Error when saving chat message after broadcast",
					slog.String("error", err.Error()))
				return
			}

			delete(p.StreamManager.history, currentPrompt.ThreadId)
		}
	})

	done <- true

	logger.Debug("Finished streaming LLM",
		slog.String("error", fmt.Sprintf("%v", err)),
	)
	if err != nil {
		return errors.Wrap(err, "error streaming llm")
	}

	return nil
}
