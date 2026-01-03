package commands

import (
	"fmt"
	"gator/internal/state"
)

type Commands struct {
	CommandsMap map[string]func(*state.State, Command) error
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	handler, exists := c.CommandsMap[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.CommandsMap[name] = f
	return
}