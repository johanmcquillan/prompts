package prompts

import "strings"

const separator = ":"

type Prompt struct {
	Components []Component
	showEmptyElements bool
}

func MakePrompt() *Prompt {
	return &Prompt{}
}

func (p *Prompt) WithComponent(c Component) *Prompt {
	p.Components = append(p.Components, c)
	return p
}

func (p *Prompt) WithEnvVar(envName string) *Prompt {
	p.Components = append(p.Components, MakeEnvComponent(envName))
	return p
}

func (p *Prompt) WithUser() *Prompt {
	return p.WithComponent(MakeUser())
}

func (p *Prompt) String() string {
	var subStrings []string
	for _, component := range p.Components {
		if component.Length() > 0 {
			subStrings = append(subStrings, component.String())
		}
	}

	return strings.Join(subStrings, separator)
}
