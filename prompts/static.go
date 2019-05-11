package prompts

import "os"

const envUser = "USER"

type EnvComponent struct {
	Colour int
	value  string
}

func MakeEnvComponent(envVar string) Component {
	return &EnvComponent{
		value: os.Getenv(envVar),
	}
}

func MakeUser() Component {
	return MakeEnvComponent(envUser)
}

func (u *EnvComponent) String() string {
	return u.value
}

func (u *EnvComponent) Length() int {
	return len(u.value)
}
