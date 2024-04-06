package handler

import (
	"fmt"
)

func init() {
	prepareRepo()
	// register the handler
	registeredHandlers["version"] = &VersionHandler{}
	// register the aliases
	registeredAliases["ver"] = &VersionHandler{}
}

type VersionHandler struct{}

func (h *VersionHandler) Handle(env *HandlerEnv) error {
	fmt.Printf("KeyKeeper v0.1.0\n")
	return nil
}
