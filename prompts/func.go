package prompts

type FunctionalComponent struct {
	Formatter
	function func() string
}

func MakeFunctionalComponent(f func() string) *FunctionalComponent {
	return &FunctionalComponent{
		function: f,
	}
}

func (c *FunctionalComponent) WithFormatter(formatter Formatter) *FunctionalComponent {
	c.Formatter = formatter
	return c
}

func (c *FunctionalComponent) MakeElement() Element {
	rawValue := c.function()
	if c.Formatter == nil {
		return Element{rawValue, len(rawValue)}
	}
	return Element{c.Format(rawValue), len(rawValue)}
}
