package main

import (
	"fmt"

	"github.com/LucasSim0n/BlogAggreGator/internal/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}

	cfg.SetUser()

	cfg, err = config.ReadConfig()
	fmt.Println(cfg.CurrentUser)
	fmt.Println(cfg.DbURL)

}
