//go:build dist
// +build dist

package node

import (
	"context"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/singulatron/singulatron/localtron/internal/di"
	node_types "github.com/singulatron/singulatron/localtron/internal/node/types"
	sdk "github.com/singulatron/singulatron/sdk/go"
	"github.com/singulatron/singulatron/sdk/go/test"
	"github.com/stretchr/testify/require"
)

func TestStart(t *testing.T) {
	hs1 := &di.HandlerSwitcher{}
	server1 := httptest.NewServer(hs1)
	defer server1.Close()

	dbprefix := sdk.Id("node_start")

	options1 := node_types.Options{
		Db:       "postgres",
		DbPrefix: dbprefix,
		Address:  server1.URL,
	}

	universe1, starterFunc1, err := Start(options1)
	require.NoError(t, err)

	hs1.UpdateHandler(universe1)
	err = starterFunc1()
	require.NoError(t, err)

	hs2 := &di.HandlerSwitcher{}
	server2 := httptest.NewServer(hs1)
	defer server1.Close()

	options2 := node_types.Options{
		Db:       "postgres",
		DbPrefix: dbprefix,
		Address:  server2.URL,
	}
	universe2, starterFunc2, err := Start(options2)
	require.NoError(t, err)

	hs2.UpdateHandler(universe2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		err := starterFunc1()
		wg.Done()
		require.NoError(t, err)

	}()

	go func() {
		err := starterFunc2()
		wg.Done()
		require.NoError(t, err)
	}()

	wg.Wait()

	// List nodes

	c := 0
	for {
		time.Sleep(100 * time.Millisecond)
		c++

		adminClient, _, err := test.AdminClient(server1.URL)
		require.NoError(t, err)

		rsp, _, err := adminClient.RegistrySvcAPI.ListNodes(context.Background()).Body(nil).Execute()
		require.NoError(t, err)

		if len(rsp.Nodes) == 2 {
			break
		}
		if c > 10 {
			require.Equal(t, 2, len(rsp.Nodes))
		}

	}
}
