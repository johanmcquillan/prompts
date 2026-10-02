package colors

import (
	"path/filepath"
	"strings"

	"github.com/johanmcquillan/prompts/env"
)

type ShellType string

const (
	UnknownShell ShellType = ""
	BASH                   = "bash"
	ZSH                    = "zsh"
	// RawANSI emits plain ANSI escape codes without any shell prompt escaping,
	// for output that is not a shell prompt (e.g. a status line). It is never
	// detected from $SHELL, so must be set with SetShellType.
	RawANSI = "raw-ansi"
)

func toShellType(s string) ShellType {
	shellType := ShellType(strings.ToLower(filepath.Base(s)))
	switch shellType {
	case BASH, ZSH:
		return shellType
	default:
		return UnknownShell
	}
}

func getShellType() ShellType {
	return toShellType(env.GetShell())
}
