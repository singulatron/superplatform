package deployment

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/singulatron/superplatform/cli/config"
	openapi "github.com/singulatron/superplatform/clients/go"
	sdk "github.com/singulatron/superplatform/sdk/go"
	"github.com/spf13/cobra"
)

// Save /home/user/deploymentA.yaml
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
		return fmt.Errorf("Failed to stat deployment file: '%v'", err)
	}
	if fileInfo.Size() == 0 {
		return fmt.Errorf("Service deployment file is empty at '%v'", filePath)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Failed to read deployment file: '%v'", err)
	}

	deployment := openapi.DeploySvcDeployment{}

	if err := yaml.Unmarshal(data, &deployment); err != nil {
		return fmt.Errorf("Failed to decode deployment file: '%v'", err)
	}

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return fmt.Errorf("Cannot get env url: '%v'", err)
	}

	cf := sdk.NewApiClientFactory(url)

	_, _, err = cf.Client(sdk.WithToken(token)).DeploySvcAPI.SaveDeployment(ctx).Body(openapi.DeploySvcSaveDeploymentRequest{
		Deployment: &deployment,
	}).Execute()
	if err != nil {
		fmt.Errorf("Failed to save deployment: '%v'", err)
	}

	return nil
}
