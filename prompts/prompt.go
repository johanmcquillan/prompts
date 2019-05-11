package prompts

import (
	"fmt"
	"strings"
)

const (
	defaultEnder     = "$"
	defaultSeparator = ":"
)

type Prompt struct {
	Components []Component
	ShowEmptyElements bool
	Ender, Separator string
}

func MakePrompt() *Prompt {
	return &Prompt{
		Ender:     defaultEnder,
		Separator: defaultSeparator,
	}
}

func (p *Prompt) String() string {
	var subStrings []string
	for _, component := range p.Components {
		e := component.GenerateElement()
		if e.Length > 0 || p.ShowEmptyElements {
			subStrings = append(subStrings, e.Output)
		}
	}

	return fmt.Sprintf("%s%s ", strings.Join(subStrings, p.Separator), p.Ender)
}

func (p *Prompt) WithComponent(c Component) *Prompt {
	p.Components = append(p.Components, c)
	return p
}

func (p *Prompt) WithEnvVar(envName string) *Prompt {
	p.Components = append(p.Components, MakeEnvComponent(envName))
	return p
}

func (p *Prompt) WithFormattedEnvVar(envName string, f Formatter) *Prompt {
	p.Components = append(p.Components, MakeEnvComponent(envName).WithFormatter(f))
	return p
}

func (p *Prompt) WithUser() *Prompt {
	return p.WithComponent(MakeUserComponent())
}

func (p *Prompt) WithFormattedUser(f Formatter) *Prompt {
	return p.WithComponent(MakeUserComponent().WithFormatter(f))
}

func (p *Prompt) WithFullWorkingDir() *Prompt {
	return p.WithComponent(MakeFullWDComponent())
}

func (p *Prompt) WithFormattedFullWorkingDir(f Formatter) *Prompt {
	return p.WithComponent(MakeFullWDComponent().WithFormatter(f))
}

func (p *Prompt) WithRelativeWorkingDir() *Prompt {
	return p.WithComponent(MakeRelativeWDComponent())
}

func (p *Prompt) WithFormattedRelativeWorkingDir(f Formatter) *Prompt {
	return p.WithComponent(MakeRelativeWDComponent().WithFormatter(f))
}

func (p *Prompt) WithGitBranch() *Prompt {
	return p.WithComponent(MakeGitBranchComponent())
}

func (p *Prompt) WithFormattedGitBranch(f Formatter) *Prompt {
	return p.WithComponent(MakeGitBranchComponent().WithFormatter(f))
}

func (p *Prompt) IncludeEmptyElements() *Prompt {
	p.ShowEmptyElements = true
	return p
}

func (p *Prompt) ExcludeEmptyElements() *Prompt {
	p.ShowEmptyElements = false
	return p
}
