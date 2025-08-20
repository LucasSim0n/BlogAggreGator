package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

func LoginHandler(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("The login command expects the username as argument: Gator login <username>")
	}

	user := cmd.Args[0]
	_, err := s.DB.GetUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("User %s doesn't exist", user)
	}

	err = writeUser(*s.Cfg, user)
	if err != nil {
		return err
	}

	fmt.Printf("The user has been set to '%s'\n", user)

	return nil
}

func writeUser(cfg config, user string) error {

	cfgPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

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
