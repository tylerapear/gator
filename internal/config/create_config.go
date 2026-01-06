package config

import (
	"fmt"
	"os"
	"encoding/json"
	"path/filepath"
)

func CreateConfig(DBUrl string) error {
	
	config := Config{
		DBUrl: DBUrl,
		CurrentUserName: "",
	}

	jsonContent, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return fmt.Errorf("Error marshaling json: %v\n", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error setting home directory: %v\n", err)
	}

	err = os.WriteFile(filepath.Join(homeDir, configFileName), jsonContent, 0644)
	if err != nil {
		return fmt.Errorf("Error writing config: %v\n", err)
	}

	return nil

}