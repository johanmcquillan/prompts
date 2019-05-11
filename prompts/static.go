package prompts

import "os"

const envUser = "USER"

type User struct {
	Colour int
	name string
}

func MakeUser() Component {
	return &User{
		name: os.Getenv(envUser),
	}
}

func (u *User) String() string {
	return u.name
}

func (u *User) Length() int {
	return len(u.name)
}
