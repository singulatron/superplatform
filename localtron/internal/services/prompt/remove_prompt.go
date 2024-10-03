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
	"log/slog"

	"github.com/singulatron/singulatron/sdk/go/datastore"
	"github.com/singulatron/singulatron/sdk/go/logger"

	firehosetypes "github.com/singulatron/singulatron/localtron/internal/services/firehose/types"
	prompttypes "github.com/singulatron/singulatron/localtron/internal/services/prompt/types"
)

func (p *PromptService) removePrompt(promptId string) error {
	logger.Info("Removing prompt",
		slog.String("promptId", promptId),
	)

	err := p.promptsStore.Query(
		datastore.Id(promptId),
	).Delete()

	if err != nil {
		return err
	}

	ev := prompttypes.EventPromptRemoved{
		PromptId: promptId,
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

	return nil
}
