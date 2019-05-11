package prompts

type StaticComponent struct {
	*AnsiFormat
	Value string
}

func MakeStaticComponent(value string) Component {
	return &StaticComponent{
		Value: value,
	}
}

func (c *StaticComponent) StringAndLength() (string, int) {
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
