package main

import (
	"fmt"
	"os"

	"github.com/LucasSim0n/BlogAggreGator/internal/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}

	s := &config.State{
		Cfg: cfg,
	}

	var cmds config.Commands
	cmds.Register("login", config.LoginHandler)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: please insert ani argument: Gator <command>")
		fmt.Println("Try 'Gator help' for more information")
		os.Exit(1)
	}
	com := config.Command{
		Name: args[1],
		Args: args[2:],
	}

	err = cmds.Run(s, com)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
