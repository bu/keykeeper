package command

import "keykeeper/cmd/kk/util"

type Command interface {
	Run(env *Env) error
}

type Env struct {
	Repo *util.Repo
}
