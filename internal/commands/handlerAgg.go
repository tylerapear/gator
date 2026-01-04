package commands

import (
	"context"
	"fmt"
	"gator/internal/state"
	"gator/internal/rssfeed"
)

func HandlerAgg(s *state.State, cmd Command) error {

	feed, err := rssfeed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}
	fmt.Printf("Fetched feed: %v\n", *feed)

	return nil
}