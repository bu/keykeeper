package handler

import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"
)

// ReleaseVersion is the version of the application
// TODO: update this version before release
const ReleaseVersion = "0.1.0"

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

	buildTime, err := time.Parse(time.RFC3339, versionTime)
	if err != nil {
		buildTime = time.Time{}
	}

	fmt.Printf("[KeyKeeper] v%s\nBuild #%s, built on %s\n", ReleaseVersion, version[:7], buildTime.Format("January 02, 2006"))
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("Copyright (C) 2024 Buwei Chiu\n")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("KeyKeeper is a password manager that helps you manage your passwords securely.")
	fmt.Println("This software is distributed under the MIT License. See LICENSE for more information.")
	fmt.Println("For more information, please visit https://github.com/bu/keykeeper")

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
