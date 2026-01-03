package commands

import (
	"fmt"
	"time"
	"os"
	"context"

	"gator/internal/state"
	"gator/internal/database"

	"github.com/google/uuid"
)

func HandlerRegister(s *state.State, cmd Command) error {

	if cmd.Args == nil || len(cmd.Args) < 1 {
		return fmt.Errorf("name argument is required")
	}

	fmt.Println(cmd.Args)

	params := database.CreateUserParams{
		ID:  uuid.New(),
		Name: cmd.Args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.DB.CreateUser(context.Background(), params)
	if err != nil {
		fmt.Println("failed to create user:", err)
		os.Exit(1)
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User %s registered and set as current user\n", user.Name)

	s.DB.GetUser(context.Background(), user.Name)
	fmt.Printf("User details:\n ID=%s,\n Name=%s,\n CreatedAt=%s,\n UpdatedAt=%s\n", user.ID, user.Name, user.CreatedAt, user.UpdatedAt)

	return nil

}