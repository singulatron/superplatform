package firehoseservice_test

import (
	"bufio"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	client "github.com/singulatron/singulatron/clients/go"
	"github.com/singulatron/singulatron/localtron/internal/di"
	firehose "github.com/singulatron/singulatron/localtron/internal/services/firehose/types"
	"github.com/singulatron/singulatron/sdk/go/test"
	"github.com/stretchr/testify/require"
)

func TestFirehoseSubscription(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)
	err = starterFunc()
	require.NoError(t, err)

	cl, adminToken, err := test.AdminClient(server.URL)
	require.NoError(t, err)

	firehoseSvc := cl.FirehoseSvcAPI

	t.Run("firehose subscription", func(t *testing.T) {
		event := &client.FirehoseSvcEvent{
			Name: client.PtrString("test-event"),
			Data: map[string]any{"hi": "there"},
		}

		eventChannel := make(chan *firehose.Event, 1)

		go func() {
			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL+"/firehose-svc/subscribe", nil)

			require.NoError(t, err)
			req.Header.Set("Authorization", "Bearer "+adminToken)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			scanner := bufio.NewScanner(resp.Body)
			for scanner.Scan() {
				line := scanner.Text()

				if !strings.HasPrefix(line, "data: ") {
					return
				}

				ev := &firehose.Event{}
				jsonData := strings.TrimSpace(strings.ReplaceAll(line, "data: ", ""))
				if jsonData == "" {
					return
				}

				err = json.Unmarshal([]byte(jsonData), &ev)
				require.NoError(t, err)

				eventChannel <- ev
			}

			require.NoError(t, scanner.Err())
		}()

		_, err := firehoseSvc.PublishEvent(context.Background()).Event(client.FirehoseSvcEventPublishRequest{
			Event: event,
		}).Execute()
		require.NoError(t, err)

		receivedEvent := <-eventChannel

		require.Equal(t, *event.Name, receivedEvent.Name)
		require.Equal(t, event.Data, receivedEvent.Data)
	})
}
