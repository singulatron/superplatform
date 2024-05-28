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
package promptendpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/singulatron/singulatron/localtron/lib"
	"github.com/singulatron/singulatron/localtron/llm"
	promptservice "github.com/singulatron/singulatron/localtron/services/prompt"
)

func Subscribe(w http.ResponseWriter, r *http.Request, promptService *promptservice.PromptService) {
	threadId := r.URL.Query().Get("threadId")

	if threadId == "" {
		http.Error(w, "Missing threadId parameter", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")

	subscriber := make(chan *llm.CompletionResponse)
	promptService.StreamManager.Subscribe(threadId, subscriber)
	defer promptService.StreamManager.Unsubscribe(threadId, subscriber)

	// Use context to handle client disconnection
	ctx := r.Context()
	go func() {
		<-ctx.Done()
		promptService.StreamManager.Unsubscribe(threadId, subscriber)
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
			lib.Logger.Warn("Warning: ResponseWriter does not support flushing, streaming might be delayed")
		}
	}
}
