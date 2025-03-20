package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	DefaultPath     string `json:"defaultPath"`
	DefaultBookName string `json:"defaultBookName"`
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
		return nil, errors.New("config key 'defaultPath' is missing")
	}
	if config.DefaultBookName == "" {
		return nil, errors.New("config key 'defaultBookName' is missing")
	}

	return &config, nil
}
