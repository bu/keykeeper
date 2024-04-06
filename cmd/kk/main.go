package main

import (
	"fmt"
	"keykeeper/cmd/kk/file"
	"keykeeper/cmd/kk/handler"
	"keykeeper/cmd/kk/lang"
	"keykeeper/cmd/kk/model"
	"os"
)

// main function is the entry point of the program
func main() {
	// configure language
	err := lang.ConfigureLanguage(
		os.Getenv("LC_ALL"),
		os.Getenv("LANG"),
		"en_us",
	)

	// if there is an error, print it
	if err != nil {
		fmt.Printf("[keykeeper] error during loading locale: %s\n", err)
		return
	}

	// then setup up runner
	runner := handler.NewRunner()

	// get the repo and register it
	repo, err := model.NewRepo(
		os.Getenv("PASSWORD_STORE_DIR"),
	)

	if err != nil {
		fmt.Printf("[keykeeper] %s\n", lang.GetString("error.repo_cannot_be_loaded"))
		return
	}

	runner.RegisterRepo(repo)

	// prepare commands to run
	commands := os.Args[1:]

	// check if repo is ready for use
	if !repo.IsReady() {
		commands = []string{"init"}
	}

	// register the file decrypter
	runner.RegisterFileHandler(file.NewPgpHandler(repo))
	runner.RegisterFileHandler(file.NewNaClHandler(repo))

	// run the runner
	if err := runner.Run(commands); err != nil {
		fmt.Print("\033[H\033[2J\033[3J")
		fmt.Printf("[keykeeper] %s", err)
		fmt.Print("\033[E")
	}
}
