package main

import _ "github.com/lib/pq"

import (
	"fmt"
	"os"
	"database/sql"

	"gator/internal/config"
	"gator/internal/commands"
	"gator/internal/state"
	"gator/internal/database"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	cmd := commands.Command{
		Name: args[1],
		Args: args[2:],
	}

	if cmd.Name == "init" {
		err := commands.HandlerInit(cmd)
		if err != nil {
			fmt.Printf("Error initializing: %v\n", err)
		}
		return
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		fmt.Println("Make sure gator is initialized: 'gator init <database url>'")
		return
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	dbQueries := database.New(db)

	currentState := state.State{
		Config: cfg,
		DB:     dbQueries,
	}

	commandsMap := commands.Commands{
		CommandsMap: make(map[string]func(*state.State, commands.Command) error),
	}

	commandsMap.Register("login", commands.HandlerLogin)
	commandsMap.Register("register", commands.HandlerRegister)
	commandsMap.Register("reset", commands.HandlerReset)
	commandsMap.Register("users", commands.HandlerUsers)
	commandsMap.Register("agg", commands.HandlerAgg)
	commandsMap.Register("addfeed", middlewareLoggedIn(commands.HandlerAddFeed))
	commandsMap.Register("feeds", commands.HandlerFeeds)
	commandsMap.Register("follow", middlewareLoggedIn(commands.HandlerFollow))
	commandsMap.Register("following", middlewareLoggedIn(commands.HandlerFollowing))
	commandsMap.Register("unfollow", middlewareLoggedIn(commands.HandlerUnfollow))
	commandsMap.Register("browse", middlewareLoggedIn(commands.HandlerBrowse))

	err = commandsMap.Run(&currentState, cmd)
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}

}