package prompts

type FunctionalComponent struct {
	Formatter
	Function func() string
}

func MakeFunctionalComponent(f func() string) *FunctionalComponent {
	return &FunctionalComponent{
		Function: f,
	}
}

func (c *FunctionalComponent) WithFormatter(formatter Formatter) *FunctionalComponent {
	c.Formatter = formatter
	return c
}

func (c *FunctionalComponent) MakeElement() Element {
	rawValue := c.Function()
	if c.Formatter == nil {
		return Element{
			Output: rawValue,
			Length: len(rawValue),
		}
	}
	return Element{
		Output: c.Format(rawValue),
		Length: len(rawValue),
	}
}
