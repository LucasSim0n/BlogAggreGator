package config

import (
	"context"
	"fmt"

	"github.com/LucasSim0n/BlogAggreGator/internal/database"
)

func UnfollowHandler(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("The unfollow command requires a URL as an argument:\ngator unfollow <URL>")
	}

	feed, err := s.DB.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	del := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.DB.DeleteFeedFollow(context.Background(), del)

	return err
}
