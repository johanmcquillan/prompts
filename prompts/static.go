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

func (c *StaticComponent) StringAndLength() (string, int) {
	if c.Formatter == nil {
		return c.Value, len(c.Value)
	}
	return c.Colourise(c.Value), len(c.Value)
}

func (c *StaticComponent) String() string {
	s, _ := c.StringAndLength()
	return s
}

func (c *StaticComponent) Length() int {
	_, l := c.StringAndLength()
	return l
}
