package commands

import (
	"fmt"
	"context"

	"gator/internal/state"
)

func HandlerFeeds(s *state.State, cmd Command) error {

	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("\nRegistered RSS Feeds:\n")
	
	for _, feed := range feeds {

		feedUser, err := s.DB.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf(" - Name: %s", feed.Name)
		fmt.Printf("\n - URL: %s", feed.Url)
		fmt.Printf("\n - Creator: %s\n\n", feedUser.Name)
	}

	return nil
}