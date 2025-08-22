package config

import "fmt"

func HelpHandler(s *State, cmd Command) error {

	fmt.Println(help)

	return nil
}

const help = `
	BlogAggreGator is a lightweight command-line application that fetches and aggregates blog or RSS feed posts.

	Usage: gator <command> [options]

	Avaliable commands:

	- help: Displays this help message
	- register <username> : Adds a user to the database
	- login <username> : Logs in with an existing user
	- users: Displays all the registered users and the currently logged one
	- addfeed <name> <url> : Adds new a feed to the database
	- feeds: Displays all the registered feeds by the given name, the asociated url and the owner
	- follow <url> : Subscribes the logged user to an already registered feed
	- following: Displays de feeds witch the logged user currently follows
	- unfollow <url> : Unfollows the provided feed
	- browse [limit] : Displays the [limit] most recent posts from the feeds followed by the user. If not given or invalid, limit defaults to 2.
	- agg <timeUnit> : Checks for fetches in the registered feeds one by one from last updated to most recently in cycles of <timeUnit> time. egg. gator agg 10m
	- reset: Resets the database to 0. Seriously... be careful :D
	`
