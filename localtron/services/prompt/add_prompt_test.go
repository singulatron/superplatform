/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptservice_test

import (
	"testing"

	"github.com/singulatron/singulatron/localtron/di"
	modeltypes "github.com/singulatron/singulatron/localtron/services/model/types"
	"github.com/stretchr/testify/require"
)

func TestAddPrompt(t *testing.T) {
	universe, err := di.BigBang(di.UniverseOptions{
		Test: true,
	})
	require.NoError(t, err)
	require.NotNil(t, universe)

	cs := universe.ConfigService

	conf, err := cs.GetConfig()
	require.NoError(t, err)

	ms := universe.ModelService
	models, err := ms.GetModels()
	require.NoError(t, err)
	var model *modeltypes.Model
	for _, v := range models {
		if v.Id == conf.Model.CurrentModelId {
			model = v
		}
	}

	require.Equal(t, true, model.Id != "")

	//ds := universe.DownloadService
	//ds.SyncDownloads = true
	//
	//err = ds.Do(model.Assets["MODEL"], "")
	//require.NoError(t, err)
	//
	//err = ms.Start("")
	//require.NoError(t, err)
	//
	//ps := universe.PromptService
	//prompt, err := ps.AddPrompt(context.Background(), &prompttypes.Prompt{
	//	Sync:   true,
	//	Prompt: "Hi there, how are you?",
	//})
	//require.NoError(t, err)
	//require.Equal(t, true, strings.Contains(prompt.Answer, "how"))
}
