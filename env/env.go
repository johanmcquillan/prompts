package env

import (
	"os"

	"github.com/johanmcquillan/prompts"
)

const (
	EnvHome  = "HOME"
	EnvPWD   = "PWD"
	EnvUser  = "USER"
	EnvShell = "SHELL"
)

type envFetcher interface {
	GetEnv(string) string
}

type EnvComponent struct {
	envFetcher
	prompts.Formatter
	envVar string
}

type Env struct{}

func (Env) GetEnv(envVar string) string {
	return os.Getenv(envVar)
}

func MakeEnvComponent(envVar string) *EnvComponent {
	return &EnvComponent{
		envFetcher: Env{},
		envVar:     envVar,
	}
}

func MakeUserComponent() *EnvComponent {
	return &EnvComponent{
		envFetcher: Env{},
		envVar:     EnvUser,
	}
}

func (c *EnvComponent) WithFormatter(formatter prompts.Formatter) *EnvComponent {
	c.Formatter = formatter
	return c
}

func (c *EnvComponent) MakeElement() prompts.Element {
	rawValue := c.GetEnv(c.envVar)

	if c.Formatter == nil {
		return prompts.Element{
			Output: rawValue,
			Length: len(rawValue),
		}
	}

	return prompts.Element{
		Output: c.Format(rawValue),
		Length: len(rawValue),
	}
}
