/*

This command is handled differently from the rest. 
It can be called without a state, since it is meant to configure a database connection.

*/

package commands

import (
	"fmt"

	"gator/internal/config"
)

func HandlerInit(cmd Command) error {

	if cmd.Args == nil || len(cmd.Args) < 1 {
		return fmt.Errorf("usage: init <DBUrl> (eg. postgres://username:password@localhost:5432/gator)")
	}

	DBUrl := cmd.Args[0]

	err := config.CreateConfig(DBUrl)
	if err != nil {
		return fmt.Errorf("Error initializing: %v\n", err)
	}

	fmt.Println("\nGator successfully initialized.\n")

	return nil

}