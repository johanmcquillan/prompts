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
	envVar  string
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

func MakeUserComponent() *EnvComponent {
	return MakeEnvComponent(envUser)
}

func MakeFullWDComponent() *EnvComponent {
	return MakeEnvComponent(envPWD)
}

func (c *EnvComponent) StringAndLength() (string, int) {
	rawValue := c.GetEnv(c.envVar)

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
