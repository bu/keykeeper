package model

import "strings"

// ListFlag is a struct that holds the flags for the list command.
// It is used to pass the flags to the list handler.
type ListFlag struct {
	NoColor     bool
	NoRecursive bool
	NoHeader    bool

	List         bool
	ListFileOnly bool

	TabIndent    bool
	Space2Indent bool
	Space4Indent bool
}

// ParseListFlag is a function that parses the flags for the list command.
// It returns a ListFlag struct.
func ParseListFlag(flags []string) *ListFlag {
	listFlag := ListFlag{}

	for _, flag := range flags {
		switch flag {
		case "-nc", "--no-color":
			listFlag.NoColor = true
		case "-nr", "--no-recursive":
			listFlag.NoRecursive = true
		case "-nh", "--no-header":
			listFlag.NoHeader = true
		case "-l", "--list":
			listFlag.List = true
		case "-f", "--list-file-only":
			listFlag.ListFileOnly = true
		case "-ti", "--tab-indent":
			listFlag.TabIndent = true
		case "-2si", "--space-2-indent":
			listFlag.Space2Indent = true
		case "-4si", "--space-4-indent":
			listFlag.Space4Indent = true
		}
	}

	return &listFlag
}

// GetIndentCharacter is a function that returns the indent character based on the flags.
func (f *ListFlag) GetIndentCharacter() string {
	if f.TabIndent {
		return "\t"
	}

	if f.Space2Indent {
		return strings.Repeat("\u00a0", 2)
	}

	return strings.Repeat("\u00a0", 4)
}
