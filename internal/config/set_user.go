package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

func (c Config) SetUser(name string) error {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	c.CurrentUserName = name

	updatedContent, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(homeDir, configFileName), updatedContent, 0644)
	if err != nil {
		return err
	}

	return nil
}