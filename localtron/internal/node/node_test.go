//go:build dist
// +build dist

package node

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/singulatron/singulatron/localtron/internal/di"
	node_types "github.com/singulatron/singulatron/localtron/internal/node/types"
	"github.com/singulatron/singulatron/sdk/go/test"
	"github.com/stretchr/testify/require"
)

func TestStart(t *testing.T) {
	hs1 := &di.HandlerSwitcher{}
	server1 := httptest.NewServer(hs1)
	defer server1.Close()

	options1 := &di.Options{
		NodeOptions: node_types.Options{
			DbPrefix: "teststart",
			Address:  server1.URL,
		},
		Test: true,
		Url:  server1.URL,
	}
	universe1, starterFunc1, err := di.BigBang(options1)
	require.NoError(t, err)

	hs1.UpdateHandler(universe1)
	err = starterFunc1()
	require.NoError(t, err)

	hs2 := &di.HandlerSwitcher{}
	server2 := httptest.NewServer(hs1)
	defer server1.Close()

	options2 := &di.Options{
		NodeOptions: node_types.Options{
			DbPrefix: "teststart",
			Address:  server2.URL,
		},
		Test: true,
		Url:  server2.URL,
	}
	universe2, starterFunc2, err := di.BigBang(options2)
	require.NoError(t, err)

	hs2.UpdateHandler(universe2)
	err = starterFunc2()
	require.NoError(t, err)

	// List nodes

	adminClient, _, err := test.AdminClient(server1.URL)
	require.NoError(t, err)

	rsp, _, err := adminClient.RegistrySvcAPI.ListNodes(context.Background()).Body(nil).Execute()
	require.NoError(t, err)
	require.Equal(t, len(rsp.Nodes), 2)
}
