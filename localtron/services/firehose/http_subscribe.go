/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package firehoseservice

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/singulatron/singulatron/localtron/logger"

	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

// Subscribe subscribes to a firehose stream and streams events to the client
// @Summary Subscribe
// @Description Establish a subscription to firehose events and stream them to the client in real-time.
// @Tags firehose
// @Accept json
// @Produce text/event-stream
// @Success 200 {string} string "Event data"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /firehose/subscribe [get]
func (p *FirehoseService) Subscribe(
	w http.ResponseWriter,
	r *http.Request,
) {
	rsp := &usertypes.IsAuthorizedResponse{}
	err := p.router.Post(r.Context(), "user", "/is-authorized", &usertypes.IsAuthorizedRequest{
		PermissionId: firehosetypes.PermissionFirehoseView.Id,
	}, rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !rsp.Authorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	eventsChannel := make(chan []*firehosetypes.Event)
	subscriberID := p.subscribe(func(events []*firehosetypes.Event) {
		eventsChannel <- events
	})
	defer p.unsubscribe(subscriberID)

	ctx := r.Context()

	go func() {
		defer func() {
			recover()
		}()
		<-ctx.Done()
		p.unsubscribe(subscriberID)
		close(eventsChannel)
	}()

	for {
		select {
		// case <-time.After(time.Second * 8):
		// 	panic("timeout test")
		case events, ok := <-eventsChannel:
			if !ok {
				logger.Info("Events channel closed unexpectedly")
				return
			}

			for _, event := range events {
				jsonResp, err := json.Marshal(event)
				if err != nil {
					log.Printf("Failed to marshal event: %v", err)
					continue
				}

				if _, err := w.Write([]byte("data: " + string(jsonResp) + "\n\n")); err != nil {
					log.Printf("Failed to write event to client: %v", err)
					return
				}

				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				} else {
					log.Println("Warning: ResponseWriter does not support flushing, streaming might be delayed")
				}
			}
		}
	}
}
