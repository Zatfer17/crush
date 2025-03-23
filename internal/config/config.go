package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DefaultPath      string `json:"defaultPath"`
	DefaultWorkspace string `json:"defaultWorkspace"`
}

func InitConfig() (*Config, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(homePath, ".config/zurg/config.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	if config.DefaultPath == "" {
		return nil, err
	}

	if config.DefaultWorkspace == "" {
		return nil, err
	}

	return &config, nil
}
