package commands

import (
	"fmt"
	"os"
	"context"

	"gator/internal/state"
)

func HandlerReset(s *state.State, cmd Command) error {

	err := s.DB.Reset(context.Background())
	if err != nil {
		fmt.Printf("failed to reset data: %w", err)
		os.Exit(1)
	}

	fmt.Println("All data has been reset.")
	return nil
}