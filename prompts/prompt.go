package prompts

import "strings"

const separator = ":"

type Prompt struct {
	Contexts []Component
}

func MakePrompt() *Prompt {
	return &Prompt{}
}

func (p *Prompt) WithComponent(c Component) *Prompt {
	p.Contexts = append(p.Contexts, c)
	return p
}

func (p *Prompt) WithUser() *Prompt {
	return p.WithComponent(MakeUser())
}

func (p *Prompt) String() string {
	subStrings := make([]string, len(p.Contexts))
	for i, c := range p.Contexts {
		subStrings[i] = c.String()
	}

	return strings.Join(subStrings, separator)
}
