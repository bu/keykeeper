package command

import "log"

type Move struct {
}

func (c *Move) Run(*Env) error {
	log.Println("Move command")
	return nil
}
