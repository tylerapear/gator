package main

import (
	"fmt"
	"os"

	"gator/internal/config"
	"gator/internal/commands"
	"gator/internal/state"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	currentState := state.State{
		Config: cfg,
	}
	

	commandsMap := commands.Commands{
		CommandsMap: make(map[string]func(*state.State, commands.Command) error),
	}

	commandsMap.Register("login", commands.HandlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	cmd := commands.Command{
		Name: args[1],
		Args: args[2:],
	}

	err = commandsMap.Run(&currentState, cmd)
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}

	fmt.Println("Final State Config User:", currentState.Config)

}