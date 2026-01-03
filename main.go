package main

import (
	"fmt"

	"gator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	err = cfg.SetUser("Tyler")
	if err != nil {
		fmt.Println("Error setting user:", err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	fmt.Println(cfg)
	


}