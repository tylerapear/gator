package commands

import (
	"context"
	"fmt"
	"time"
	"database/sql"

	"gator/internal/state"
	"gator/internal/rssfeed"
	"gator/internal/database"
)

func HandlerAgg(s *state.State, cmd Command) error {

	if cmd.Args == nil || len(cmd.Args) < 1 {
		return fmt.Errorf("usage: agg <time_between_reqs> (eg. 1s, 1m, 1h, etc.)")
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("\nCollecting feeds every %v\n\n", time_between_reqs)

	i := 0
	ticker := time.NewTicker(time_between_reqs)
	defer ticker.Stop()
	for ; ; <-ticker.C{
		i += 1
		fmt.Printf("request #%d\n", i)
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Printf("Error scraping feed: %v", err)
		}
	}

	return nil
}

func scrapeFeeds(s *state.State) error {

	feed, err := s.DB.GetNextFeed(context.Background())
	if err != nil {
		return err
	}

	params := database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now()},
		UpdatedAt: time.Now(),
	}

	err = s.DB.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return err
	}

	fetchedFeed, err := rssfeed.FetchFeed(context.Background(), feed.Url)

	for _, item := range fetchedFeed.Channel.Items {
		fmt.Printf("\nFeed Item Title: %s", item.Title)
	}
	fmt.Println("\n")

	return nil

}