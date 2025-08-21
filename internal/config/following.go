package config

import (
	"context"
	"fmt"
)

func FollowingHandler(s *State, cmd Command) error {
	user, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUser)
	if err != nil {
		return err
	}

	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}
