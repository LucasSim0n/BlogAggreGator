package config

import (
	"context"
	"fmt"
	"github.com/LucasSim0n/BlogAggreGator/internal/database"
	"github.com/google/uuid"
	"time"
)

func AddFeedHandler(s *State, cmd Command, user database.User) error {

	if len(cmd.Args) < 2 {
		return fmt.Errorf("addfeed requires 2 arguments:\ngator addfeed <name> <url>")
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
		UserID:    user.ID,
	}

	s.DB.CreateFeed(context.Background(), fd)

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedId,
	}

	s.DB.CreateFeedFollow(context.Background(), feedFollow)

	fmt.Println(fd)
	return nil
}
