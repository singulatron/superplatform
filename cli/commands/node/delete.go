package node

import (
	"fmt"

	"github.com/singulatron/superplatform/cli/config"
	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/spf13/cobra"
)

// Delete [nodeUrl]
func Delete(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	nodeUrl := args[0]

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return fmt.Errorf("Cannot get env url: '%v'", err)
	}

	cf := sdk.NewApiClientFactory(url)

	_, err = cf.Client(sdk.WithToken(token)).RegistrySvcAPI.DeleteNode(ctx, nodeUrl).Execute()
	if err != nil {
		return fmt.Errorf("Error deleting node: '%v'", err)
	}

	return nil
}
