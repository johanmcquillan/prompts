package colors

import (
	"os"
	"path/filepath"
	"strings"
)

type ShellType string

const (
	UnknownShell ShellType = ""
	BASH                   = "bash"
	ZSH                    = "zsh"
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
	return toShellType(os.Getenv("SHELL"))
}
