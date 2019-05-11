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

func (u *EnvComponent) StringAndLength() (string, int) {
	var formattedValue string
	rawValue := os.Getenv(u.envVar)

	if u.Format == nil {
		formattedValue = rawValue
	} else {
		formattedValue = u.Format.Colourise(rawValue)
	}

	return formattedValue, len(rawValue)
}

func (u *EnvComponent) String() string {
	s, _ := u.StringAndLength()
	return s
}

func (u *EnvComponent) Length() int {
	_, l := u.StringAndLength()
	return l
}
