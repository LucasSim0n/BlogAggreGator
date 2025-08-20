package config

import (
	"context"
	"fmt"

	"github.com/LucasSim0n/BlogAggreGator/internal/rss"
)

const ReqURL = "https://www.wagslane.dev/index.xml"

func AggHandler(s *State, cmd Command) error {

	feed, err := rss.FetchFeed(context.Background(), ReqURL)
	if err != nil {
		return err
	}

	fmt.Println(*feed)

	return nil
}
