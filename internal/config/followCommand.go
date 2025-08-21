package config

import (
	"context"
	"fmt"
	"time"

	"github.com/LucasSim0n/BlogAggreGator/internal/database"
	"github.com/google/uuid"
)

func FollowHandler(s *State, cmd Command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("The follow command takes a URL as argument:\ngator follow <URL>")
	}
	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	user, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUser)
	if err != nil {
		return err
	}

	follow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), follow)
	if err != nil {
		return err
	}

	fmt.Printf("%s successfully followed by %s", feed.Name, user.Name)
	return nil
}
