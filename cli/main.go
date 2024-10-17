package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/singulatron/superplatform/cli/commands/env"
	"github.com/singulatron/superplatform/cli/commands/run"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "sup",
		Short: "Superplatform CLI",

		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	addEnvCommands(rootCmd)
	addRunCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addEnvCommands(rootCmd *cobra.Command) {
	var envCmd = &cobra.Command{
		Use:   "env",
		Short: "Manage environments",
	}

	var envAddCmd = &cobra.Command{
		Use:   "add [name] [url] [description]",
		Short: "Add a new environment",
		Args:  cobra.RangeArgs(2, 3),
		RunE:  env.Add,
	}

	var envRemoveCmd = &cobra.Command{
		Use:   "remove [name]",
		Short: "Remove an environment",
		Args:  cobra.ExactArgs(1), // Requires exactly 1 argument
		RunE:  env.Remove,
	}

	var envListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all environments",
		RunE:  env.List,
	}

	envCmd.AddCommand(envAddCmd)
	envCmd.AddCommand(envRemoveCmd)
	envCmd.AddCommand(envListCmd)

	rootCmd.AddCommand(envCmd)
}

func addRunCommands(rootCmd *cobra.Command) {
	var runCmd = &cobra.Command{
		Use:   "run [filePath]",
		Args:  cobra.ExactArgs(1),
		Short: "Run service(s) found from a JSON or YAML file.",
		RunE:  run.Run,
	}

	// Add 'run' command to the root command
	rootCmd.AddCommand(runCmd)
}
