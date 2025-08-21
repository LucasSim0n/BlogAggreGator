package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/LucasSim0n/BlogAggreGator/internal/database"
	"github.com/LucasSim0n/BlogAggreGator/internal/rss"
)

const ReqURL = "https://www.wagslane.dev/index.xml"

func AggHandler(s *State, cmd Command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("The agg command requires a 'time between requests' argument:\negg. gator agg <10m>")
	}

	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", duration)

	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *State) error {

	nextFetch, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	feed, err := rss.FetchFeed(context.Background(), nextFetch.Url)
	if err != nil {
		return err
	}

	fetch := database.MarkFeedFetchedParams{
		ID:            nextFetch.ID,
		UpdatedAt:     time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}
	err = s.DB.MarkFeedFetched(context.Background(), fetch)
	if err != nil {
		return err
	}

	fmt.Println(feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
	}
	fmt.Print("\n\n")
	return nil
}
