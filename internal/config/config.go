package config

import (
	"encoding/json"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(homeDir, configFileName), nil
}

func Read() (Config, error) {
	fullUrl, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(fullUrl)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil

}

func write(cfg Config) error {
	fullUrl, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	if err := os.WriteFile(fullUrl, data, 0644); err != nil {
		return err
	}

	return nil
}

func SetUser(user_name string) error {
	// Read current and assign in a variable
	currentConfig, err := Read()
	if err != nil {
		return err
	}

	// Update current variable
	currentConfig.CurrentUserName = user_name

	// Save back to the json file
	if err := write(currentConfig); err != nil {
		return err
	}

	return nil

}
