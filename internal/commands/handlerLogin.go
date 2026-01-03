package commands

import (
	"fmt"
	"gator/internal/state"
)

func HandlerLogin(s *state.State, cmd Command) error {

	if cmd.Args == nil || len(cmd.Args) < 1 {
		return fmt.Errorf("username argument is required")
	}

	username := cmd.Args[0]

	err := s.Config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User set to %s\n", username)
	return nil
}