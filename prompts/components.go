package prompts

import "fmt"

type Component interface {
	fmt.Stringer
	Length() int
}