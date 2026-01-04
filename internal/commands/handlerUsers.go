package commands

import (
	"fmt"
	"context"

	"gator/internal/state"
)

func HandlerUsers(s *state.State, cmd Command) error {
	
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to retrieve users: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("No users found.")
		return nil
	}

	fmt.Println("Registered Users:")
	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Printf(" * %s (current)\n", user.Name)
			continue
		}
		fmt.Printf(" * %s\n", user.Name)
	}

	return nil

}