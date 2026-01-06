package main

import (
	"context"
	"fmt"

	"gator/internal/state"
	"gator/internal/database"
	"gator/internal/commands"
)

func middlewareLoggedIn(handler func(s *state.State, cmd commands.Command, user database.User) error) func(*state.State, commands.Command) error {
	return func(s *state.State, cmd commands.Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Error getting current user: %v\nTry registering a user: gator register <username>", err)
		}
	
		return handler(s, cmd, user)
	} 
}