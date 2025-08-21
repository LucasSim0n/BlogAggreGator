package config

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/LucasSim0n/BlogAggreGator/internal/database"
	"github.com/LucasSim0n/BlogAggreGator/internal/rss"
	"github.com/google/uuid"
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
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
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

	for _, item := range feed.Channel.Item {

		//NOTE: Not greate, but works by now
		if strings.Contains(item.Description, "<a href") {
			continue
		}

		var date sql.NullTime
		parsed, err := parsePubDate(item.PubDate)

		if err != nil {
			date = sql.NullTime{Time: time.Time{}, Valid: false}
		} else {
			date = sql.NullTime{Time: parsed, Valid: true}
		}

		post := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: date,
			FeedID:      nextFetch.ID,
		}

		err = s.DB.CreatePost(context.Background(), post)
		if err != nil && !strings.Contains(err.Error(), "«posts_url_key»") {
			return err
		}
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

	return nil
}

func parsePubDate(pubDate string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC3339,
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, pubDate)
		if err == nil {
			return t, nil
		}
	}
	return t, err
}
