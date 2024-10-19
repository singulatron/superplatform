package env

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/singulatron/superplatform/cli/config"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if len(conf.Environments) == 0 {
		fmt.Println("No environments found.")
		return nil
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "SELECTED\tNAME\tURL\tDESCRIPTION")

	for name, env := range conf.Environments {
		selected := ""
		if conf.SelectedEnvironment == name {
			selected = "*"
		}

		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t\n", selected, name, env.URL, env.Description)
	}

	return nil
}
