package prompts

import (
	"fmt"
	"strings"
)

const (
	defaultEnder = "$"
	defaultSeparator = ":"
)

type Prompt struct {
	Ender
	Components        []Component
	ShowEmptyElements bool
	Separator         string
}

func MakePrompt() *Prompt {
	return &Prompt{
		Ender:     MakeStaticEnder(defaultEnder),
		Separator: defaultSeparator,
	}
}

func (p *Prompt) Print(exitCode int) string {
	s := p.String(exitCode)
	fmt.Print(s)
	return s
}

func (p *Prompt) String(exitCode int) string {
	var subStrings []string
	for _, component := range p.Components {
		e := component.MakeElement()
		if e.Length > 0 || p.ShowEmptyElements {
			subStrings = append(subStrings, e.Output)
		}
	}

	return fmt.Sprintf("%s%s ", strings.Join(subStrings, p.Separator), p.End(exitCode).Output)
}

func (p *Prompt) WithComponent(c Component) *Prompt {
	p.Components = append(p.Components, c)
	return p
}

func (p *Prompt) WithEnder(e Ender) *Prompt {
	p.Ender = e
	return p
}

func (p *Prompt) IncludeEmptyElements() *Prompt {
	p.ShowEmptyElements = true
	return p
}

func (p *Prompt) ExcludeEmptyElements() *Prompt {
	p.ShowEmptyElements = false
	return p
}
