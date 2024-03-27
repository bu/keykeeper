package command

import "log"

// Unknown is a command that is run when the requested command is not recognized
// by the application. Original Pass will search this unknown command in the
// repository to check if it is an item. If so, pass to "list" command to show the
// item. If not, display an error message and leads the user to the help command.
type Unknown struct{}

func (u *Unknown) Run(*Env) error {
	log.Printf("Unknown command\n")
	return nil
}
