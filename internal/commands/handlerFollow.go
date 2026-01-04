package commands

import (
	"fmt"
	"context"
	"time"

	"gator/internal/state"
	"gator/internal/database"

	"github.com/google/uuid"
)

func HandlerFollow(s *state.State, cmd Command, currentUser database.User) error {

	if cmd.Args == nil || len(cmd.Args) < 1 {
		return fmt.Errorf("usage: follow <url>")
	}

	feed, err := s.DB.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:  uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: currentUser.ID,
		FeedID: feed.ID,
	}

	feedFollowRow, err := s.DB.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("\nFollowed Feed:\n")
	fmt.Printf(" - UserID: %s\n", feedFollowRow.UserID)
	fmt.Printf(" - FeedID: %s\n\n", feedFollowRow.FeedID)

	return nil

}