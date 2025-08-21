package config

import (
	"context"
	"fmt"

	"github.com/LucasSim0n/BlogAggreGator/internal/database"
)

func FollowingHandler(s *State, cmd Command, user database.User) error {

	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}
