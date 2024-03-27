package command

import "log"

type Git struct {
}

func (c *Git) Run(*Env) error {
	log.Printf("Git command")
	return nil
}
