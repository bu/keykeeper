package handler

import (
	"keykeeper/cmd/kk/file"
	"keykeeper/cmd/kk/model"
	"sync"
)

var repoInit sync.Once

var registeredHandlers map[string]HandlerInterface
var registeredAliases map[string]HandlerInterface

func prepareRepo() {
	repoInit.Do(func() {
		registeredHandlers = make(map[string]HandlerInterface)
		registeredAliases = make(map[string]HandlerInterface)
	})
}

type HandlerInterface interface {
	Handle(env *HandlerEnv) error
}

type HandlerEnv struct {
	Args        []string
	Flags       []string
	Repo        *model.Repo
	FileHandler map[string]file.Handler
}
