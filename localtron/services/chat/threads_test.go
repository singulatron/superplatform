package chatservice_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	chatservice "github.com/singulatron/singulatron/localtron/services/chat"
	chattypes "github.com/singulatron/singulatron/localtron/services/chat/types"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

func TestMessageCreatesThread(t *testing.T) {
	cs, err := configservice.NewConfigService()
	require.NoError(t, err)
	us, err := userservice.NewUserService(cs)
	require.NoError(t, err)
	fs, err := firehoseservice.NewFirehoseService(us)
	require.NoError(t, err)
	as, err := chatservice.NewChatService(cs, fs, us)
	require.NoError(t, err)

	t.Run("no thread id", func(t *testing.T) {
		err := as.AddMessage(&chattypes.Message{
			Id:      uuid.New().String(),
			Content: "hi there",
		})
		require.Error(t, err)
	})

	t.Run("thread does not exist", func(t *testing.T) {
		err := as.AddMessage(&chattypes.Message{
			Id:       uuid.New().String(),
			ThreadId: "1",
			Content:  "hi there",
		})
		require.Error(t, err)
	})

	t.Run("no user id", func(t *testing.T) {
		tid := uuid.New().String()
		title := "Test Thread Title"
		_, err := as.AddThread(&chattypes.Thread{
			Id:    tid,
			Title: title,
		})
		require.Error(t, err)
	})

	userId := "usr-1"
	var threadId string

	t.Run("create thread", func(t *testing.T) {
		tid := uuid.New().String()
		title := "Test Thread Title"
		thread, err := as.AddThread(&chattypes.Thread{
			Id:      tid,
			Title:   title,
			UserIds: []string{userId},
		})
		require.NoError(t, err)
		require.Equal(t, tid, thread.Id)
		require.Equal(t, title, thread.Title)
		threadId = thread.Id
	})

	t.Run("no user id", func(t *testing.T) {
		err := as.AddMessage(&chattypes.Message{
			Id:       uuid.New().String(),
			ThreadId: threadId,
			Content:  "hi there",
		})
		require.NoError(t, err)
	})
}
