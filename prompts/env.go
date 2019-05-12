package prompts

import "os"

const (
	envHome = "HOME"
	envPWD  = "PWD"
	envUser = "USER"
)

type envFetcher interface {
	GetEnv(string) string
}

type EnvComponent struct {
	envFetcher
	Formatter
	envVar string
}

type ActualEnv struct {}

func (ActualEnv) GetEnv(envVar string) string {
	return os.Getenv(envVar)
}

func MakeEnvComponent(envVar string) *EnvComponent {
	return &EnvComponent{
		envFetcher: ActualEnv{},
		envVar: envVar,
	}
}

func (c *EnvComponent) WithFormatter(formatter Formatter) *EnvComponent {
	c.Formatter = formatter
	return c
}

func (c *EnvComponent) MakeElement() Element {
	rawValue := c.GetEnv(c.envVar)

	if c.Formatter == nil {
		return Element{rawValue, len(rawValue)}
	}

	return Element{c.Format(rawValue), len(rawValue)}
}
