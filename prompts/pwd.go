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
			return substitutePathPrefix(os.Getenv(envHome), os.Getenv(envPWD), homeSymbol)
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

func substitutePathPrefix(prefixPath, fullPath, substitution string) string {
	relativePath, err := filepath.Rel(prefixPath, fullPath)
	if err != nil || strings.Contains(relativePath, upOne) {
		return fullPath
	}

	return filepath.Join(substitution, relativePath)
}
