/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptservice_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/singulatron/singulatron/localtron/clients/llm"
	"github.com/singulatron/singulatron/localtron/di"
	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAddPrompt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ds := downloadtypes.NewMockDownloadServiceI(ctrl)
	ds.EXPECT().Do("https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_S.gguf", gomock.Any())
	ds.EXPECT().SetDefaultFolder(gomock.Any())
	ds.EXPECT().SetStateFilePath(gomock.Any())
	ds.EXPECT().Start()

	lc := llm.NewMockClientI(ctrl)

	ms := modeltypes.NewMockModelServiceI(ctrl)
	ms.EXPECT().Start(gomock.Any())
	ms.EXPECT().GetModels().Return([]*modeltypes.Model{{
		Id: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_S.gguf",
		},
		PlatformId:     "llama-cpp",
		Name:           "Mistral",
		Parameters:     "7B",
		Flavour:        "Instruct",
		Version:        "v0.2",
		Quality:        "Q3_K_S",
		Extension:      "GGUF",
		FullName:       "Mistral 7B Instruct v0.2 Q3_K_S",
		Size:           3.16,
		MaxRam:         5.66,
		QuantComment:   "very small, high quality loss",
		Description:    "hi",
		PromptTemplate: "[INST] {prompt} [/INST]",
	}}, nil)
	ms.EXPECT().Status(gomock.Any()).Return(&modeltypes.ModelStatus{
		AssetsReady: true,
		Running:     true,
		Address:     "127.0.0.1:8888",
	}, nil)
	ms.EXPECT().GetPlatformByModelId(gomock.Any()).Return(&modeltypes.Platform{
		Id: modeltypes.PlatformLlamaCpp.Id,
	}, nil)

	responses := []*llm.CompletionResponse{
		{
			Choices: []struct {
				Text         string      `json:"text,omitempty"`
				Index        int         `json:"index,omitempty"`
				Logprobs     interface{} `json:"logprobs,omitempty"`
				FinishReason string      `json:"finish_reason,omitempty"`
			}{
				{Text: "Hi, I'm fine", FinishReason: ""},
			},
		},
		{
			Choices: []struct {
				Text         string      `json:"text,omitempty"`
				Index        int         `json:"index,omitempty"`
				Logprobs     interface{} `json:"logprobs,omitempty"`
				FinishReason string      `json:"finish_reason,omitempty"`
			}{
				{Text: ", how are you", FinishReason: "stop"},
			},
		},
	}

	lc.EXPECT().
		PostCompletionsStreamed(gomock.Any(), gomock.Any()).
		DoAndReturn(func(req llm.PostCompletionsRequest, callback llm.StreamCallback) error {
			go func() {
				for i := range responses {
					// without this sleep the test hangs forever
					time.Sleep(1 * time.Millisecond)
					callback(responses[i])
				}

			}()
			return nil
		})

	universe, err := di.BigBang(di.UniverseOptions{
		Test: true,
		Pre: di.Universe{
			DownloadService: ds,
			ModelService:    ms,
			LLMClient:       lc,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, universe)

	cs := universe.ConfigService

	conf, err := cs.GetConfig()
	require.NoError(t, err)

	models, err := ms.GetModels()
	require.NoError(t, err)
	var model *modeltypes.Model
	for _, v := range models {
		if v.Id == conf.Model.CurrentModelId {
			model = v
		}
	}

	require.Equal(t, true, model.Id != "")

	err = ds.Do(model.Assets["MODEL"], "")
	require.NoError(t, err)

	err = ms.Start("")
	require.NoError(t, err)

	ps := universe.PromptService

	prompt, err := ps.AddPrompt(context.Background(), &prompttypes.AddPromptRequest{
		PromptCreateFields: prompttypes.PromptCreateFields{
			Sync:   true,
			Prompt: "Hi there, how are you?",
		},
	})

	require.NoError(t, err)
	require.Equal(t, true, strings.Contains(prompt.Answer, "how"))
}
