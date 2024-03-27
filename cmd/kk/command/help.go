package command

import "log"

type Help struct{}

func (h *Help) Run(*Env) error {
	log.Printf("Help command executed\n")
	return nil
}
