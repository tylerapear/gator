package commands

import (
	"context"
	"fmt"
	"strconv"
	//"time"

	"gator/internal/state"
	"gator/internal/database"
)

func HandlerBrowse(s *state.State, cmd Command, currentUser database.User) error {

	var limit int32 = 2
	if cmd.Args == nil || len(cmd.Args) > 0 {
		n64, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("Error parsing int: %v\n", err)
		}
		limit = int32(n64)
	}

	params := database.GetPostsForUserParams{
		UserID: currentUser.ID,
		Limit: limit,
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error getting user's posts: %v\n", err)
	}

	for _, post := range posts {

		feed, err := s.DB.GetFeedByID(context.Background(), post.FeedID)
		if err != nil {
			return fmt.Errorf("Error retrieving feed: %v\n", err)
		}

		fmt.Printf("\nTitle: %s\n", post.Title)
		fmt.Printf("Url: %s\n", post.Url)
		if post.PublishedAt.Valid {
			fmt.Printf("Published At: %s\n", post.PublishedAt.Time.Format("2006-01-02 @ 15:04"))
		} else {
			fmt.Println("Published At: unknown")
		}
		fmt.Printf("Feed: %s\n", feed.Name)
	}
	fmt.Println("")

	return nil

}