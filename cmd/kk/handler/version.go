package handler

import (
	"fmt"
	"runtime/debug"
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
	version, versionTime := h.extractVcsInfo()

	fmt.Printf("KeyKeeper %s (built at %s)\n", version, versionTime)
	return nil
}

func (h *VersionHandler) extractVcsInfo() (string, string) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown", ""
	}

	ver := ""
	verTime := ""

	for _, kv := range info.Settings {
		if kv.Key == "vcs.revision" {
			ver = kv.Value
		}

		if kv.Key == "vcs.time" {
			verTime = kv.Value
		}
	}

	return ver, verTime
}
