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
	function func() string
}

func MakeFunctionalComponent(f func() string) *FunctionalComponent {
	return &FunctionalComponent{
		function: f,
	}
}

func MakeRelativeWDComponent() *FunctionalComponent {
	return &FunctionalComponent{
		function: func() string {
			return insertPathHomeSymbol(os.Getenv(envPWD))
		},
	}
}

func (c *FunctionalComponent) GenerateElement() Element {
	rawValue := c.function()
	if c.Formatter == nil {
		return Element{rawValue, len(rawValue)}
	}
	return Element{c.Colourise(rawValue), len(rawValue)}
}

func (c *FunctionalComponent) WithFormatter(formatter Formatter) *FunctionalComponent {
	c.Formatter = formatter
	return c
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
