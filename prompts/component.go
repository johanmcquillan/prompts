package prompts

type Element struct {
	Output string
	Length int
}

type Component interface {
	Formatter
	GenerateElement() Element
}

type Formatter interface {
	Colourise(string) string
}