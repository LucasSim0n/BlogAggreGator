package config

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/LucasSim0n/BlogAggreGator/internal/database"
)

func BrowseHandler(s *State, cmd Command, user database.User) error {

	var limit int32 = 2

	if len(cmd.Args) > 0 {
		limit = validLimit(cmd.Args[0])
	}

	req := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), req)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Description)
	}

	return nil
}

func validLimit(str string) int32 {
	var baseLim int32 = 2
	lim, err := strconv.Atoi(str)
	if err != nil {
		return baseLim
	}

	if lim < 1 || lim > math.MaxInt32 {
		return baseLim
	}

	return int32(lim)
}
