package prompts

type Element struct {
	Output string
	Length int
}

type Component interface {
	Formatter
	MakeElement() Element
}

type Ender interface {
	End(int) Element
}

type Formatter interface {
	Format(string) string
}
