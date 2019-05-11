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

func (c *FunctionalComponent) StringAndLength() (string, int) {
	rawValue := c.function()
	if c.Formatter == nil {
		return rawValue, len(rawValue)
	}
	return c.Colourise(rawValue), len(rawValue)
}

func (c *FunctionalComponent) String() string {
	s, _ := c.StringAndLength()
	return s
}

func (c *FunctionalComponent) Length() int {
	_, l := c.StringAndLength()
	return l
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
