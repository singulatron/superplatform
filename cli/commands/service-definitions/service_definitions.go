package service_definitions

import "github.com/spf13/cobra"

func AddServiceDefinitionCommands(rootCmd *cobra.Command) {
	var envCmd = &cobra.Command{
		Use:     "service-definition",
		Aliases: []string{"sd", "service-definition"},
		Short:   "Manage service definitions",
	}

	var envSaveCmd = &cobra.Command{
		Use:   "save [filePath]",
		Args:  cobra.ExactArgs(1),
		Short: "Save service definition(s) found in a JSON or YAML file",
		RunE:  Save,
	}

	var envDeleteCmd = &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a service definition",
		Args:  cobra.ExactArgs(1),
		RunE:  Delete,
	}

	var envListCmd = &cobra.Command{
		Use:   "list",
		Short: "List service definitions",
		RunE:  List,
	}

	envCmd.AddCommand(envSaveCmd)
	envCmd.AddCommand(envDeleteCmd)
	envCmd.AddCommand(envListCmd)

	rootCmd.AddCommand(envCmd)
}
