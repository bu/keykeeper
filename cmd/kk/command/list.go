package command

import "log"

type List struct{}

func (l *List) Run(env *Env) error {
	log.Printf("current repo = %s\n", env.Repo.GetPath())
	log.Printf("List command executed\n")
	return nil
}
