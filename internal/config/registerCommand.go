package config

import (
	"context"
	"fmt"
	"time"

	"github.com/LucasSim0n/BlogAggreGator/internal/database"
	"github.com/google/uuid"
)

func RegisterHandler(s *State, cmd Command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("The register command expects the username as argument: Gator register <username>")
	}
	user := cmd.Args[0]
	cup := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      user,
	}

	_, err := s.DB.CreateUser(context.Background(), cup)
	if err != nil {
		return err
	}

	cmd.Name = "login"
	fmt.Printf("The user %s has been succesfuly created.\n", user)
	err = LoginHandler(s, cmd)
	if err != nil {
		return err
	}

	return nil
}
