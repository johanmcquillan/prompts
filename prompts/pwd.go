package prompts

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	homeSymbol = "~"
	upOne = ".."

	pathSeparator = string(os.PathSeparator)
)

type FunctionalComponent struct {
	Formatter
	function func(...interface{}) string
}

func MakeFunctionalComponent(f func(...interface{}) string) *FunctionalComponent {
	return &FunctionalComponent{
		function: f,
	}
}

func MakeRelativeWDComponent() *FunctionalComponent {
	return &FunctionalComponent{
		function: func(...interface{}) string {
			return insertPathHomeSymbol(os.Getenv(envPWD))
		},
	}
}

func insertPathHomeSymbol(path string) string {
	homePath := os.Getenv(envHome)
	if homePath == "" {
		return ""
	}

	relPath, err := filepath.Rel(homePath, path)
	if err != nil || strings.Contains(relPath, upOne) {
		return ""
	}

	return filepath.Join(homeSymbol, relPath)
}
