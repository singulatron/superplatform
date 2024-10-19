package run

import (
	"fmt"

	"github.com/singulatron/superplatform/cli/config"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	fmt.Println("Run is not implemented yet")

	_, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	return nil
}
