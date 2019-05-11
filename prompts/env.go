package prompts

import "os"

const (
	envUser = "USER"
	envPWD = "PWD"
	envHome = "HOME"
)

type EnvComponent struct {
	Formatter
	envVar  string
}

func MakeEnvComponent(envVar string) *EnvComponent {
	return &EnvComponent{
		envVar: envVar,
	}
}

func MakeUserComponent() *EnvComponent {
	return MakeEnvComponent(envUser)
}

func MakeFullWDComponent() *EnvComponent {
	return MakeEnvComponent(envPWD)
}

func (c *EnvComponent) StringAndLength() (string, int) {
	rawValue := os.Getenv(c.envVar)

	if c.Formatter == nil {
		return rawValue, len(rawValue)
	}
	return c.Colourise(rawValue), len(rawValue)
}

func (c *EnvComponent) String() string {
	s, _ := c.StringAndLength()
	return s
}

func (c *EnvComponent) Length() int {
	_, l := c.StringAndLength()
	return l
}
