package command

import "log"

type Remove struct {
}

func (c *Remove) Run(*Env) error {
	log.Println("Remove command")
	return nil
}
