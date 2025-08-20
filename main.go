package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/LucasSim0n/BlogAggreGator/internal/config"
	"github.com/LucasSim0n/BlogAggreGator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s := &config.State{
		Cfg: cfg,
		DB:  dbQueries,
	}

	var cmds config.Commands
	cmds.Register("login", config.LoginHandler)
	cmds.Register("register", config.RegisterHandler)
	cmds.Register("reset", config.ResetHandler)
	cmds.Register("users", config.UsersHandler)
	cmds.Register("agg", config.AggHandler)
	cmds.Register("addfeed", config.AddFeedHandler)

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
