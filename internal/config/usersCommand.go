package config

import (
	"context"
	"fmt"
)

func UsersHandler(s *State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting users from db")
	}

	for _, user := range users {
		if user == s.Cfg.CurrentUser {
			fmt.Printf("%s (current)\n", user)
		} else {
			fmt.Println(user)
		}
	}

	return nil
}
