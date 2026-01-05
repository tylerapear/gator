package commands

import (
	"context"
	"fmt"
	"time"
	"database/sql"

	"gator/internal/state"
	"gator/internal/rssfeed"
	"gator/internal/database"

	"github.com/google/uuid"
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
			fmt.Printf("Error scraping feed: %v\n", err)
		}
	}

	return nil
}

func scrapeFeeds(s *state.State) error {

	feed, err := s.DB.GetNextFeed(context.Background())
	if err != nil {
		return err
	}

	feedParams := database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: time.Now(),
	}

	err = s.DB.MarkFeedFetched(context.Background(), feedParams)
	if err != nil {
		return err
	}

	feed, err = s.DB.GetNextFeed(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Feed: %s\n", feed.Name)
	fmt.Printf("LastFetchedAt: %v", feed.LastFetchedAt)

	fmt.Println(feed.Url)
	fetchedFeed, err := rssfeed.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %v", err)
	}
	//fmt.Println(fetchedFeed)
	fmt.Println("before scraping")
	for _, item := range fetchedFeed.Channel.Items {
		fmt.Println("scraping item")

		pubTime, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
		if err != nil {
			return err
		}

		postParams := database.CreatePostParams{
			ID:  uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: pubTime, Valid: true},
			FeedID: feed.ID,
		}

		_, err = s.DB.CreatePost(context.Background(), postParams)
		if err != nil {
			return err
		}
		fmt.Printf("\nFeed Item Title: %s", item.Title)

	}
	fmt.Println("\n")

	return nil

}

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt sql.NullTime
	FeedID      uuid.UUID
}