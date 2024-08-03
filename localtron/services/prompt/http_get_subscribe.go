/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/singulatron/singulatron/localtron/clients/llm"
	"github.com/singulatron/singulatron/localtron/logger"

	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Subscribe streams prompt responses to the client
// @Summary Subscribe to Prompt
// @Description Subscribe to prompt responses via Server-Sent Events (SSE)
// @Tags Prompt Service
// @Param threadId path string true "Thread ID"
// @Success 200 {string} string "Streaming response"
// @Failure 400 {object} prompttypes.ErrorResponse "Missing threadId parameter"
// @Failure 401 {object} prompttypes.ErrorResponse "Unauthorized"
// @Router /prompt-service/{threadId}/subscribe [get]
func (p *PromptService) GetSubscribe(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := p.router.AsRequestMaker(r).Post(r.Context(), "user-service", fmt.Sprintf("/permission/%v/is-authorized", prompttypes.PermissionPromptStream.Id), &usertypes.IsAuthorizedRequest{}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	if vars["threadId"] == "" {
		http.Error(w, "Missing threadId path parameter", http.StatusBadRequest)
		return
	}
	threadId := vars["threadId"]

	w.Header().Set("Content-Type", "text/event-stream")

	subscriber := make(chan *llm.CompletionResponse)
	p.Subscribe(threadId, subscriber)
	defer p.Unsubscribe(threadId, subscriber)

	// Use context to handle client disconnection
	ctx := r.Context()
	go func() {
		<-ctx.Done()
		p.Unsubscribe(threadId, subscriber)
	}()

	for resp := range subscriber {
		resp.Model = "" // Redact model from response
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Printf("Failed to marshal JSON: %v", err)
			continue
		}

		if _, writeErr := w.Write([]byte("data: " + string(jsonResp) + "\n")); writeErr != nil {
			log.Printf("Failed to write streaming response: %v", writeErr)
			break // Exit the loop on write errors
		}

		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		} else {
			logger.Warn("Warning: ResponseWriter does not support flushing, streaming might be delayed")
		}
	}
}
