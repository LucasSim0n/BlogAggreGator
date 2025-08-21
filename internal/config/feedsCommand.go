package config

import (
	"context"
	"fmt"
)

func FeedsHandler(s *State, cmd Command) error {
	feeds, err := s.DB.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.DB.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Feed name: %s\n", feed.Name)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		fmt.Printf("Feed owner: %s\n", user)
	}

	return nil
}
