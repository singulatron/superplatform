package config

import (
	"fmt"

	"os"
	"path/filepath"

	types "github.com/singulatron/superplatform/cli/types"

	"gopkg.in/yaml.v2"
)

func LoadConfig() (types.Config, error) {
	var config types.Config
	configPath := filepath.Join(os.Getenv("HOME"), ".singulatron", "cliConfig.yaml")

	file, err := os.Open(configPath)
	if err != nil {
		return config, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("failed to decode config file: %v", err)
	}

	return config, nil
}

func SaveConfig(config types.Config) error {
	configPath := filepath.Join(os.Getenv("HOME"), ".singulatron", "cliConfig.yaml")

	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config file for writing: %v", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to encode config file: %v", err)
	}

	return nil
}
