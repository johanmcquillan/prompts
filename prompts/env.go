package prompts

import "os"

const envUser = "USER"

type EnvComponent struct {
	Format *AnsiFormat
	envVar  string
}

func MakeEnvComponent(envVar string) Component {
	return &EnvComponent{
		envVar: envVar,
	}
}

func MakeUserComponent() Component {
	return MakeEnvComponent(envUser)
}

func (c *EnvComponent) StringAndLength() (string, int) {
	var formattedValue string
	rawValue := os.Getenv(c.envVar)

	if c.Format == nil {
		formattedValue = rawValue
	} else {
		formattedValue = c.Format.Colourise(rawValue)
	}

	return formattedValue, len(rawValue)
}

func (c *EnvComponent) String() string {
	s, _ := c.StringAndLength()
	return s
}

func (c *EnvComponent) Length() int {
	_, l := c.StringAndLength()
	return l
}
