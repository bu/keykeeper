package command

import "log"

type Edit struct {
}

func (e *Edit) Run(*Env) error {
	log.Printf("Edit command")
	return nil
}
