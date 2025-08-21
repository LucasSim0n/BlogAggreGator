package config

import (
	"context"
	"fmt"
	"github.com/LucasSim0n/BlogAggreGator/internal/database"
	"github.com/google/uuid"
	"time"
)

func AddFeedHandler(s *State, cmd Command) error {

	if len(cmd.Args) < 2 {
		return fmt.Errorf("addfeed requires 2 arguments:\ngator addfeed <name> <url>")
	}

	userData, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUser)
	if err != nil {
		return err
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]
	feedId := uuid.New()

	fd := database.CreateFeedParams{
		ID:        feedId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    userData.ID,
	}

	s.DB.CreateFeed(context.Background(), fd)

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userData.ID,
		FeedID:    feedId,
	}

	s.DB.CreateFeedFollow(context.Background(), feedFollow)

	fmt.Println(fd)
	return nil
}
