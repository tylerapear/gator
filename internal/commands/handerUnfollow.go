package commands

import (
	"fmt"
	"context"

	"gator/internal/state"
	"gator/internal/database"
)

func HandlerUnfollow(s *state.State, cmd Command, currentUser database.User) error {

	if cmd.Args == nil || len(cmd.Args) < 1 {
		return fmt.Errorf("usage: unfollow <url>")
	}

	feed, err := s.DB.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: currentUser.ID,
	}

	err = s.DB.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("\nDeleted Feed Follow: %s\n", feed.Url)

	return nil
}