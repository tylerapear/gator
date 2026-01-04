package commands

import (
	"fmt"
	"context"
	
	"gator/internal/state"
	"gator/internal/database"
)

func HandlerFollowing(s *state.State, cmd Command, currentUser database.User) error {

	feedFollows, err := s.DB.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return err
	}

	fmt.Printf("\nFeeds followed by %s:\n", currentUser.Name)
	for _, feedFollow := range feedFollows {
		fmt.Printf("\n - %s", feedFollow.FeedName)
	}
	fmt.Println("\n")

	return nil
}