package env

import (
	"path/filepath"

	"github.com/johanmcquillan/prompts/src/prompts"
)

type ShellComponent struct {
	envFetcher
	prompts.Formatter
	envVar string
}

func (c *ShellComponent) GetShell() string {
	shell := c.GetEnv(EnvShell)
	return filepath.Base(shell)
}

func MakeShellComponent() *ShellComponent {
	return &ShellComponent{
		envFetcher: Env{},
	}
}

func (c *ShellComponent) WithFormatter(formatter prompts.Formatter) *ShellComponent {
	c.Formatter = formatter
	return c
}

func (c *ShellComponent) MakeElement() prompts.Element {
	rawValue := c.GetShell()

	if c.Formatter == nil {
		return prompts.Element{rawValue, len(rawValue)}
	}

	return prompts.Element{c.Format(rawValue), len(rawValue)}
}
