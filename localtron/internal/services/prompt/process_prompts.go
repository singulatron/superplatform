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
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/singulatron/singulatron/sdk/go/clients/llm"
	"github.com/singulatron/singulatron/sdk/go/clients/stable_diffusion"
	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/logger"

	apptypes "github.com/singulatron/singulatron/localtron/internal/services/chat/types"
	chattypes "github.com/singulatron/singulatron/localtron/internal/services/chat/types"
	configtypes "github.com/singulatron/singulatron/localtron/internal/services/config/types"
	firehosetypes "github.com/singulatron/singulatron/localtron/internal/services/firehose/types"
	modeltypes "github.com/singulatron/singulatron/localtron/internal/services/model/types"
	prompttypes "github.com/singulatron/singulatron/localtron/internal/services/prompt/types"
)

var TimeNow = time.Now

const (
	maxRetries    = 5
	BaseDelay     = 1 * time.Second
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
		datastore.Equal(datastore.Field("status"), prompttypes.PromptStatusRunning),
	).Find()
	if err != nil {
		return err
	}

	hasRunning := false
	runningPromptId := ""
	for _, runningPromptI := range runningPrompts {
		runningPrompt := runningPromptI.(*prompttypes.Prompt)

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

	currentPrompt, err := SelectPrompt(p.promptsStore)
	if err != nil {
		return err
	}
	if currentPrompt == nil {
		return nil
	}

	return p.processPrompt(currentPrompt)
}

func (p *PromptService) processPrompt(currentPrompt *prompttypes.Prompt) (err error) {

	updateCurr := func() {
		logger.Info("Prompt finished",
			slog.String("promptId", currentPrompt.Id),
			slog.String("status", string(currentPrompt.Status)),
			slog.Any("error", err),
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

		err = p.promptsStore.Query(
			datastore.Id(currentPrompt.Id),
		).Update(currentPrompt)
		if err != nil {
			logger.Error("Error updating prompt",
				slog.String("promptId", currentPrompt.Id),
				slog.String("error", err.Error()),
			)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			currentPrompt.Error = fmt.Sprintf("%v", r)
			currentPrompt.Status = prompttypes.PromptStatusErrored
			updateCurr()
		}

		if err != nil {
			currentPrompt.Error = err.Error()
			currentPrompt.Status = prompttypes.PromptStatusErrored
		} else {
			currentPrompt.Status = prompttypes.PromptStatusCompleted
		}

		updateCurr()
	}()

	logger.Info("Picking up prompt from queue",
		slog.String("promptId", currentPrompt.Id),
	)

	defer func() {
		ev := prompttypes.EventPromptProcessingFinished{
			PromptId: currentPrompt.Id,
			Error:    errToString(err),
		}
		err = p.router.Post(context.Background(), "firehose-svc", "/publish", firehosetypes.PublishRequest{
			Event: &firehosetypes.Event{
				Name: ev.Name(),
				Data: ev,
			},
		}, nil)
		if err != nil {
			logger.Error("Failed to publish: %v", err)
		}
	}()

	currentPrompt.LastRun = time.Now()
	currentPrompt.Error = ""
	currentPrompt.Status = prompttypes.PromptStatusRunning
	currentPrompt.RunCount++

	err = p.promptsStore.Upsert(currentPrompt)
	if err != nil {
		return errors.Wrap(err, "error updating currently running prompt")
	}

	ev := prompttypes.EventPromptProcessingStarted{
		PromptId: currentPrompt.Id,
	}
	err = p.router.Post(context.Background(), "firehose-svc", "/publish", firehosetypes.PublishRequest{
		Event: &firehosetypes.Event{
			Name: ev.Name(),
			Data: ev,
		},
	}, nil)
	if err != nil {
		logger.Error("Failed to publish: %v", err)
	}

	addMessageReq := &apptypes.AddMessageRequest{
		Message: &apptypes.Message{
			// not a fan of taking the prompt id but at least it makes this idempotent
			// in case prompts get retried over and over again
			Id:        currentPrompt.Id,
			ThreadId:  currentPrompt.ThreadId,
			UserId:    currentPrompt.UserId,
			Content:   currentPrompt.Prompt,
			CreatedAt: time.Now(),
		},
	}

	err = p.router.Post(context.Background(), "chat-svc", fmt.Sprintf("/thread/%v/message", currentPrompt.ThreadId), addMessageReq, nil)
	if err != nil {
		return err
	}

	modelId := currentPrompt.ModelId
	if modelId == "" {
		//getConfigReq := configtypes.GetConfigRequest{}
		getConfigRsp := configtypes.GetConfigResponse{}
		err := p.router.Get(context.Background(), "config-svc", "/config", nil, &getConfigRsp)
		if err != nil {
			return err
		}
		modelId = getConfigRsp.Config.Model.CurrentModelId
	}

	statusRsp := modeltypes.StatusResponse{}
	err = p.router.Get(context.Background(), "model-svc", fmt.Sprintf("/model/%v/status", url.PathEscape(modelId)), nil, &statusRsp)
	if err != nil {

		return err
	}

	stat := statusRsp.Status
	if !stat.Running {
		return fmt.Errorf("model '%v' is not running", modelId)
	}
	if stat.Address == "" {
		return errors.Wrap(err, "missing model address")
	}
	if !strings.HasPrefix(stat.Address, "http") {
		stat.Address = "http://" + stat.Address
	}

	fullPrompt := currentPrompt.Prompt
	if currentPrompt.Template != "" {
		fullPrompt = strings.Replace(currentPrompt.Template, "{prompt}", currentPrompt.Prompt, -1)
	}

	err = p.processPlatform(stat.Address, modelId, fullPrompt, currentPrompt)

	logger.Debug("Finished streaming LLM",
		slog.String("error", fmt.Sprintf("%v", err)),
	)
	if err != nil {
		return errors.Wrap(err, "error streaming llm")
	}

	return nil
}

func (p *PromptService) processPlatform(address string, modelId string, fullPrompt string, currentPrompt *prompttypes.Prompt) error {
	getModelRsp := modeltypes.GetModelResponse{}
	err := p.router.Get(context.Background(), "model-svc", fmt.Sprintf("/model/%v", url.PathEscape(modelId)), nil, &getModelRsp)
	if err != nil {
		return err
	}

	switch getModelRsp.Platform.Id {
	case modeltypes.PlatformLlamaCpp.Id:
		return p.processLlamaCpp(address, fullPrompt, currentPrompt)
	case modeltypes.PlatformStableDiffusion.Id:
		return p.processStableDiffusion(address, fullPrompt, currentPrompt)
	}

	return fmt.Errorf("cannot find platform %v", getModelRsp.Platform.Id)
}

func (p *PromptService) processStableDiffusion(address string, fullPrompt string, currentPrompt *prompttypes.Prompt) error {
	sd := stable_diffusion.Client{
		Address: address,
	}

	req := stable_diffusion.PredictRequest{
		FnIndex: 1,
		Params: stable_diffusion.StableDiffusionParams{
			Prompt:        fullPrompt,
			NumImages:     1,
			Steps:         50,
			Width:         512,
			Height:        512,
			GuidanceScale: 7.5,
			Seed:          0,
			Flag1:         false,
			Flag2:         false,
			Scheduler:     "PNDM",
			Rate:          0.25,
		},
	}
	req.ConvertParamsToData()

	rsp, err := sd.Predict(req)
	if err != nil {
		return err
	}

	if len(rsp.Data) == 0 {
		return errors.New("no image in response")
	}

	imgUrl := stable_diffusion.FileURL(address, rsp.Data[0].FileData[0].Name)

	base64String, err := stable_diffusion.GetImageAsBase64(imgUrl)
	if err != nil {
		return err
	}
	if len(base64String) == 0 {
		return errors.New("empty image acquired")
	}

	asset := &apptypes.Asset{
		Id:      uuid.New().String(),
		Content: base64String,
	}

	upsertReq := chattypes.UpsertAssetsRequest{
		Assets: []*apptypes.Asset{
			asset,
		},
	}
	upsertRsp := chattypes.UpsertAssetsResponse{}
	err = p.router.Post(context.Background(), "chat", "/upsert-assets", upsertReq, &upsertRsp)
	if err != nil {
		return err
	}

	addMsgReq := chattypes.AddMessageRequest{
		Message: &apptypes.Message{
			Id:       uuid.New().String(),
			ThreadId: currentPrompt.ThreadId,
			Content:  "Sure, here is your image",
			AssetIds: []string{asset.Id},
		},
	}
	addMsgRsp := chattypes.AddMessageResponse{}
	err = p.router.Post(context.Background(), "chat", "/message/add", addMsgReq, &addMsgRsp)
	if err != nil {
		logger.Error("Error when saving chat message after image generation",
			slog.String("error", err.Error()))
		return err
	}

	return nil
}

func (p *PromptService) processLlamaCpp(address string, fullPrompt string, currentPrompt *prompttypes.Prompt) error {
	var llmClient llm.ClientI
	if p.llmCLient != nil {
		llmClient = p.llmCLient
	} else {
		llmClient = &llm.Client{
			LLMAddress: address,
		}
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

	err := llmClient.PostCompletionsStreamed(llm.PostCompletionsRequest{
		Prompt:    fullPrompt,
		Stream:    true,
		MaxTokens: 1000000,
	}, func(resp *llm.CompletionResponse) {
		mu.Lock()
		responseCount++
		mu.Unlock()

		p.StreamManager.Broadcast(currentPrompt.ThreadId, resp)

		if len(resp.Choices) > 0 && resp.Choices[0].FinishReason == "stop" {
			defer func() {
				done <- true
			}()

			addMsgReq := chattypes.AddMessageRequest{
				Message: &apptypes.Message{
					Id:       uuid.New().String(),
					ThreadId: currentPrompt.ThreadId,
					Content:  llmResponseToText(p.StreamManager.History[currentPrompt.ThreadId]),
				},
			}
			addMsgRsp := chattypes.AddMessageResponse{}
			err := p.router.Post(context.Background(), "chat-svc", fmt.Sprintf("/thread/%v/message", currentPrompt.ThreadId), addMsgReq, &addMsgRsp)
			if err != nil {
				logger.Error("Error when saving chat message after broadcast",
					slog.String("error", err.Error()))
				return
			}

			delete(p.StreamManager.History, currentPrompt.ThreadId)

		}
	})

	return err
}
