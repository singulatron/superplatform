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

	"github.com/singulatron/singulatron/localtron/datastore"
	"github.com/singulatron/singulatron/localtron/logger"

	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

func (p *PromptService) Remove(prompt *prompttypes.Prompt) error {
	if prompt.Status != prompttypes.PromptStatusScheduled {
		return nil
	}

	logger.Info("Removing prompt",
		slog.String("promptId", prompt.Id),
	)

	err := p.promptsStore.Query(
		datastore.Id(prompt.Id),
	).Delete()

	if err != nil {
		return err
	}

	p.firehoseService.Publish(prompttypes.EventPromptRemoved{
		PromptId: prompt.Id,
	})

	return nil
}
