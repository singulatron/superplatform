package node

import "github.com/spf13/cobra"

func AddNodeCommands(rootCmd *cobra.Command) {
	var envCmd = &cobra.Command{
		Use:     "node",
		Aliases: []string{"nodes"},
		Short:   "Manage nodes",
	}

	var envDeleteCmd = &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a node",
		Args:  cobra.ExactArgs(1),
		RunE:  Delete,
	}

	var envListCmd = &cobra.Command{
		Use:   "list",
		Short: "List nodes",
		RunE:  List,
	}

	envCmd.AddCommand(envDeleteCmd)
	envCmd.AddCommand(envListCmd)

	rootCmd.AddCommand(envCmd)
}
