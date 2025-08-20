package config

import (
	"context"
	"fmt"
)

func ResetHandler(s *State, cmd Command) error {
	err := s.DB.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error dropping users table: %v", err)
	}

	return nil
}
