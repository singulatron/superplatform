package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/singulatron/superplatform/cli/commands/env"
	"github.com/singulatron/superplatform/cli/commands/login"
	"github.com/singulatron/superplatform/cli/commands/run"
	"github.com/singulatron/superplatform/cli/commands/whoami"
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
	addLoginCommands(rootCmd)
	addWhoamiCommands(rootCmd)

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
		Args:  cobra.ExactArgs(1),
		RunE:  env.Remove,
	}

	var envListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all environments",
		RunE:  env.List,
	}

	var envCurrentCmd = &cobra.Command{
		Use:   "current",
		Short: "Display current environment",
		RunE:  env.Current,
	}

	envCmd.AddCommand(envAddCmd)
	envCmd.AddCommand(envRemoveCmd)
	envCmd.AddCommand(envListCmd)
	envCmd.AddCommand(envCurrentCmd)

	rootCmd.AddCommand(envCmd)
}

func addRunCommands(rootCmd *cobra.Command) {
	var runCmd = &cobra.Command{
		Use:   "run [filePath]",
		Args:  cobra.ExactArgs(1),
		Short: "Run service(s) found from a JSON or YAML file.",
		RunE:  run.Run,
	}

	rootCmd.AddCommand(runCmd)
}

func addLoginCommands(rootCmd *cobra.Command) {
	var runCmd = &cobra.Command{
		Use:   "login [userName] [password]",
		Args:  cobra.ExactArgs(2),
		Short: "Log in to the currently selected env.",
		RunE:  login.Login,
	}

	rootCmd.AddCommand(runCmd)
}

func addWhoamiCommands(rootCmd *cobra.Command) {
	var runCmd = &cobra.Command{
		Use:   "whoami",
		Args:  cobra.ExactArgs(0),
		Short: "Display the user currently logged in",
		RunE:  whoami.Whoami,
	}

	rootCmd.AddCommand(runCmd)
}
