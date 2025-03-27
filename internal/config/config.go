package config

import (
    "encoding/json"
    "os"
    "path/filepath"
)

type Config struct {
    DefaultPath string `json:"defaultPath"`
}

func InitConfig() (*Config, error) {
    homePath, err := os.UserHomeDir()
    if err != nil {
        return nil, err
    }

    configDir := filepath.Join(homePath, ".config/crush")
    configPath := filepath.Join(configDir, "config.json")
    defaultCrushPath := filepath.Join(homePath, "Documents/Crush")

    if err := os.MkdirAll(configDir, 0755); err != nil {
        return nil, err
    }

    if err := os.MkdirAll(defaultCrushPath, 0755); err != nil {
        return nil, err
    }

    data, err := os.ReadFile(configPath)
    if err != nil {
        if os.IsNotExist(err) {

			config := Config{
                DefaultPath: defaultCrushPath,
            }
            
            data, err = json.MarshalIndent(config, "", "  ")
            if err != nil {
                return nil, err
            }

            if err := os.WriteFile(configPath, data, 0644); err != nil {
                return nil, err
            }

            return &config, nil
        }
        return nil, err
    }

    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, err
    }

    if config.DefaultPath == "" {
        config.DefaultPath = defaultCrushPath
        data, err = json.MarshalIndent(config, "", "  ")
        if err != nil {
            return nil, err
        }

        if err := os.WriteFile(configPath, data, 0644); err != nil {
            return nil, err
        }
    }

    return &config, nil
}