package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const cfgFile = "/.gatorconfig.json"

type Config struct {
	DbURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func ReadConfig() (Config, error) {
	var cfg Config
	home, err := os.UserHomeDir()
	if err != nil {
		return cfg, err
	}

	cfgPath := home + cfgFile
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

func (cfg *Config) SetUser() error {
	fmt.Print("Set user ~> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	user := scanner.Text()

	err := write(cfg, user)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg *Config, user string) error {

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	cfgPath := home + cfgFile

	cfg.CurrentUser = user

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(cfgPath, data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
