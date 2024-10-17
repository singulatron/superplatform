package env

import (
	"fmt"

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

	fmt.Println("Environments:")
	for name, env := range conf.Environments {
		fmt.Printf("Name: %s\nURL: %s\nDescription: %s\n\n", name, env.URL, env.Description)
	}

	return nil
}
