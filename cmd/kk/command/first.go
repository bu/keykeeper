package command

import "log"

type FirstRun struct {
}

func (c *FirstRun) Run(*Env) error {
	log.Printf("Running the first run command\n")
	return nil
}
