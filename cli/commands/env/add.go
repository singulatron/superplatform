package env

import (
	"fmt"

	"github.com/singulatron/superplatform/cli/config"
	"github.com/singulatron/superplatform/cli/types"
	"github.com/spf13/cobra"
)

// Add envName http://address.com:8090 "A description"
func Add(cmd *cobra.Command, args []string) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	shortName := args[0]
	url := args[1]
	longDesc := args[2]

	env, ok := conf.Environments[shortName]
	if !ok {
		conf.Environments[shortName] = &types.Environment{
			ShortName:   shortName,
			URL:         url,
			Description: longDesc,
		}
		return config.SaveConfig(conf)
	}

	env.URL = url
	env.Description = longDesc

	return config.SaveConfig(conf)
}
