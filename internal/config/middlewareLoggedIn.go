package config

import (
	"context"

	"github.com/LucasSim0n/BlogAggreGator/internal/database"
)

func LoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, c Command) error {
		u, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUser)
		if err != nil {
			return err
		}
		err = handler(s, c, u)
		if err != nil {
			return err
		}
		return nil
	}
}
