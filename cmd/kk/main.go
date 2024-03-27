package main

import (
	"fmt"
	"keykeeper/cmd/kk/command"
	"keykeeper/cmd/kk/util"
	"os"
)

// main is the entry point of the application. It checks initial states
// and runs the requested command.
func main() {
	var cmd command.Command
	var requestCmd string

	// first we create a repository object to check the state of the repository
	repo := util.NewRepo(os.Getenv("PASSWORD_STORE_DIR"))

	// if the repository is not ready, run "FirstRun" command
	if !repo.IsReady() {
		requestCmd = "init"
	}

	// if there is no command, means repo is ready and user wants to list the passwords
	if requestCmd == "" && len(os.Args) < 2 {
		requestCmd = "list"
	}

	// if still there is no command, get the command from the arguments
	if requestCmd == "" {
		requestCmd = os.Args[1]
	}

	// switch the command and create the command object
	switch requestCmd {
	case "init":
		cmd = &command.FirstRun{}
	case "show":
		cmd = &command.List{}
	case "ls":
		cmd = &command.List{}
	case "list":
		cmd = &command.List{}
	case "help":
		cmd = &command.Help{}
	case "rm":
		cmd = &command.Remove{}
	case "mv":
		cmd = &command.Move{}
	case "edit":
		cmd = &command.Edit{}
	case "search":
		cmd = &command.Search{}
	case "grep":
		cmd = &command.Search{}
	case "generate":
		cmd = &command.Generate{}
	case "git":
		cmd = &command.Git{}
	case "version":
		cmd = &command.Version{}
	default:
		// if the command is not recognized, run the unknown command
		cmd = &command.Unknown{}
	}

	// run the command with the repository object
	if err := cmd.Run(&command.Env{
		Repo: repo,
	}); err != nil {
		fmt.Printf("[keykeeper] error = %v", err)
		os.Exit(1)
	}
}
