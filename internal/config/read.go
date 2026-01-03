package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

func Read() (*Config, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(filepath.Join(homeDir, configFileName))
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}