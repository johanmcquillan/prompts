package prompts

type StaticComponent struct {
	Formatter
	Value string
}

func MakeStaticComponent(value string) *StaticComponent {
	return &StaticComponent{
		Value: value,
	}
}

func (c *StaticComponent) WithFormatter(formatter Formatter) *StaticComponent {
	c.Formatter = formatter
	return c
}

func (c *StaticComponent) MakeElement() Element {
	if c.Formatter == nil {
		return Element{
			Output: c.Value,
			Length: len(c.Value),
		}
	}
	return Element{
		Output: c.Format(c.Value),
		Length: len(c.Value),
	}
}
