package commands

import (
	"fmt"
	"context"
	"time"

	"gator/internal/database"
	"gator/internal/state"

	"github.com/google/uuid"
)

func HandlerAddFeed(s *state.State, cmd Command) error {

	if cmd.Args == nil || len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	currentUser, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	paramsCreate := database.CreateFeedParams{
		ID:  uuid.New(),
		Name: feedName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url: feedURL,
		UserID: currentUser.ID,
	}

	feed, err := s.DB.CreateFeed(context.Background(), paramsCreate)
	if err != nil {
		return err
	}

	fmt.Printf("\nFeed successfully created by %s:\n", currentUser.Name)
	fmt.Printf(" - ID: %v\n", feed.ID)
	fmt.Printf(" - Name: %s\n", feed.Name)
	fmt.Printf(" - URL: %s\n", feed.Url)

	paramsFollow := database.CreateFeedFollowParams{
		ID:  uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: currentUser.ID,
		FeedID: feed.ID,
	}

	feedFollowRow, err := s.DB.CreateFeedFollow(context.Background(), paramsFollow)
	if err != nil {
		return err
	}

	fmt.Printf("\nFollowed Feed:\n")
	fmt.Printf(" - UserID: %s\n", feedFollowRow.UserID)
	fmt.Printf(" - FeedID: %s\n\n", feedFollowRow.FeedID)

	return nil
}