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
package firehoseendpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/singulatron/singulatron/localtron/lib"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	firehosetypes "github.com/singulatron/singulatron/localtron/services/firehose/types"

	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

func Subscribe(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	fs *firehoseservice.FirehoseService,
) {
	err := userService.IsAuthorized(firehosetypes.PermissionFirehoseView.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	eventsChannel := make(chan []firehosetypes.Event)
	subscriberID := fs.Subscribe(func(events []firehosetypes.Event) {
		eventsChannel <- events
	})
	defer fs.Unsubscribe(subscriberID)

	ctx := r.Context()

	go func() {
		defer func() {
			recover()
		}()
		<-ctx.Done()
		fs.Unsubscribe(subscriberID)
		close(eventsChannel)
	}()

	for {
		select {
		// case <-time.After(time.Second * 8):
		// 	panic("timeout test")
		case events, ok := <-eventsChannel:
			if !ok {
				lib.Logger.Info("Events channel closed unexpectedly")
				return
			}

			for _, event := range events {
				toFrontend := firehosetypes.FrontendEvent{
					Name: event.Name(),
					Data: event,
				}

				jsonResp, err := json.Marshal(toFrontend)
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
