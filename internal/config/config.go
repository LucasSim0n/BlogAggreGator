package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const cfgFile = "/.gatorconfig.json"

type State struct {
	Cfg *config
}

type config struct {
	DbURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func ReadConfig() (*config, error) {
	var cfg *config

	cfgPath, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	file, err := os.Open(cfgPath)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	body, err := io.ReadAll(file)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(body, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(home, cfgFile)
	return fullPath, nil
}
