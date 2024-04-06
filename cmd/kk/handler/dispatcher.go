package handler

import (
	"fmt"
	"keykeeper/cmd/kk/lang"
	"strings"
)

func init() {
	prepareRepo()
	registeredHandlers["dispatch"] = &Dispatcher{}
}

// Dispatcher is a special handler that is used to dispatch other handlers.
type Dispatcher struct {
}

func (h *Dispatcher) Handle(env *HandlerEnv) error {
	// first we extract the command first
	// normally, it was called like kk --arguments list --flags core
	requestCommand := "list"

	if len(env.Args) > 0 {
		requestCommand = env.Args[0]
	}

	requestCommand = strings.ToLower(requestCommand)

	// if requested is dispatch, we return an error
	if requestCommand == "dispatch" {
		return fmt.Errorf("%s", lang.GetString("error.dispatch.dispatching_dispatch"))
	}

	// check if we can first match a reserved command
	handler, ok := registeredHandlers[requestCommand]
	if ok {
		// remove the command from the arguments
		if len(env.Args) > 1 {
			env.Args = env.Args[1:]
		} else {
			env.Args = []string{"."}
		}

		return handler.Handle(env)
	}

	// if not, we check if this was a path of items
	if env.Repo.CheckExists(requestCommand) {
		return registeredHandlers["list"].Handle(env)
	}

	// if not, we check the aliases
	handler, ok = registeredAliases[requestCommand]
	if ok {
		// remove the alias from the arguments
		env.Args = env.Args[1:]
		return handler.Handle(env)
	}

	// finally, we return an error
	return fmt.Errorf("%s", lang.GetString("error.unknown_command", env.Args[0]))
}
