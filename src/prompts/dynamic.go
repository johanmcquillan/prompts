package prompts

type DynamicComponent struct {
	Formatter
	*PromptState
	Function func(p *PromptState) string
}

func MakeDynamicComponent(f func(p *PromptState) string) *DynamicComponent {
	return &DynamicComponent{
		Function: f,
	}
}

func (c *DynamicComponent) ForPrompt(p *PromptState) *DynamicComponent {
	c.PromptState = p
	return c
}

func (c *DynamicComponent) WithFormatter(formatter Formatter) *DynamicComponent {
	c.Formatter = formatter
	return c
}

func (c *DynamicComponent) MakeElement() Element {
	rawValue := c.Function(c.PromptState)
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
