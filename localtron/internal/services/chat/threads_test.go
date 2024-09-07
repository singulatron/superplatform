package chatservice_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/singulatron/singulatron/localtron/internal/di"
	chattypes "github.com/singulatron/singulatron/localtron/internal/services/chat/types"
	sdk "github.com/singulatron/singulatron/sdk/go"
)

func TestMessageCreatesThread(t *testing.T) {
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
	router := options.Router

	err = starterFunc()
	require.NoError(t, err)

	token, err := sdk.RegisterUser(router, "someuser", "pw123", "Some name")
	require.NoError(t, err)
	router = router.SetBearerToken(token)

	t.Run("no thread id", func(t *testing.T) {
		req := &chattypes.AddMessageRequest{
			Message: &chattypes.Message{
				Id:      sdk.Id("msg"),
				Content: "hi there",
			},
		}
		err = router.Post(context.Background(), "chat-svc", "/message", req, nil)
		require.Error(t, err)
	})

	t.Run("thread does not exist", func(t *testing.T) {
		req := &chattypes.AddMessageRequest{
			Message: &chattypes.Message{
				Id:       sdk.Id("msg"),
				ThreadId: "1",
				Content:  "hi there",
			},
		}
		err = router.Post(context.Background(), "chat-svc", "/message", req, nil)
		require.Error(t, err)

	})

	t.Run("no user id should not fail", func(t *testing.T) {
		tid := sdk.Id("thr")
		title := "Test Thread Title"

		req := &chattypes.AddThreadRequest{
			Thread: &chattypes.Thread{
				Id:    tid,
				Title: title,
			},
		}

		err = router.Post(context.Background(), "chat-svc", "/thread", req, nil)
		require.NoError(t, err)
	})

	userId := "usr-1"
	var threadId string

	t.Run("create thread", func(t *testing.T) {
		tid := sdk.Id("thr")
		title := "Test Thread Title"

		req := &chattypes.AddThreadRequest{
			Thread: &chattypes.Thread{
				Id:      tid,
				Title:   title,
				UserIds: []string{userId},
			},
		}
		rsp := &chattypes.AddThreadResponse{}
		err = router.Post(context.Background(), "chat-svc", "/thread", req, rsp)
		require.NoError(t, err)

		thread := rsp.Thread

		require.Equal(t, tid, thread.Id)
		require.Equal(t, title, thread.Title)
		threadId = thread.Id
	})

	t.Run("no user id", func(t *testing.T) {
		req := chattypes.AddMessageRequest{
			Message: &chattypes.Message{
				Id:       sdk.Id("msg"),
				ThreadId: threadId,
				Content:  "hi there",
			}}
		err = router.Post(context.Background(), "chat-svc", fmt.Sprintf("/thread/%v/message", threadId), req, nil)
		require.NoError(t, err)
	})
}
