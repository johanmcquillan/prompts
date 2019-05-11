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

func (c *StaticComponent) GenerateElement() Element {
	if c.Formatter == nil {
		return Element{c.Value, len(c.Value)}
	}
	return Element{c.Colourise(c.Value), len(c.Value)}
}

func (c *StaticComponent) WithFormatter(formatter Formatter) *StaticComponent {
	c.Formatter = formatter
	return c
}
