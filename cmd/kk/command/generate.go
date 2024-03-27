package command

import "log"

type Generate struct{}

func (g *Generate) Run(*Env) error {
	log.Printf("generate command")
	return nil
}
