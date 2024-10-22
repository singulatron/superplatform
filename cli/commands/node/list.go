package node

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/singulatron/superplatform/cli/config"
	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return fmt.Errorf("Cannot get env url: '%v'", err)
	}

	cf := sdk.NewApiClientFactory(url)

	rsp, _, err := cf.Client(sdk.WithToken(token)).RegistrySvcAPI.ListNodes(ctx).Execute()
	if err != nil {
		fmt.Errorf("Failed to list nodes: '%v'", err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "URL\tLAST HEARTBEAT")

	for _, node := range rsp.Nodes {
		const layout = "2006-01-02T15:04:05.999999999-07:00"
		ago := ""
		if node.LastHeartbeat != nil {
			t, err := time.Parse(layout, *node.LastHeartbeat)
			if err != nil {
				return err
			}

			duration := time.Since(t)
			ago = roundDuration(duration)

		}
		fmt.Fprintf(writer, "%s\t%s\n", *node.Url, ago)
	}

	return nil
}

func roundDuration(d time.Duration) string {
	// You can adjust the rounding logic as needed
	if d < time.Minute {
		// Round to nearest second
		return fmt.Sprintf("%.0fs", d.Seconds())
	} else if d < time.Hour {
		// Round to nearest minute
		return fmt.Sprintf("%dm", int(d.Minutes()))
	} else {
		// Round to nearest hour
		return fmt.Sprintf("%dh", int(d.Hours()))
	}
}
