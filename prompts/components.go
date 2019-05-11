package prompts

import "fmt"

type Component interface {
	fmt.Stringer
	Formatter
	Length() int
	StringAndLength() (string, int)
}

type Formatter interface {
	Colourise(string) string
}
