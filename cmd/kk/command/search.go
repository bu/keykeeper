package command

import "log"

type Search struct {
}

func (s *Search) Run(*Env) error {
	log.Printf("Search command")
	return nil
}
