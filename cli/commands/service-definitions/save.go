package service_definitions

import (
	"fmt"
	"os"

	"github.com/singulatron/superplatform/cli/config"
	openapi "github.com/singulatron/superplatform/clients/go"
	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Save /home/user/serviceA.yaml
func Save(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	filePath := args[0]

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("Cannot apply nonexistent file at '%v'", filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Failed to open file: '%v'", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Failed to stat service definition file: '%v'", err)
	}
	if fileInfo.Size() == 0 {
		return fmt.Errorf("Service definition file is empty at '%v'", filePath)
	}

	serviceDefinition := openapi.RegistrySvcServiceDefinition{}

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&serviceDefinition); err != nil {
		return fmt.Errorf("Failed to decode service definition file: '%v'", err)
	}

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return fmt.Errorf("Cannot get env url: '%v'", err)
	}

	cf := sdk.NewApiClientFactory(url)

	_, _, err = cf.Client(sdk.WithToken(token)).RegistrySvcAPI.SaveServiceDefinition(ctx).Request(openapi.RegistrySvcSaveServiceDefinitionRequest{
		ServiceDefinition: &serviceDefinition,
	}).Execute()
	if err != nil {
		fmt.Errorf("Failed to save service definition: '%v'", err)
	}

	return nil
}
