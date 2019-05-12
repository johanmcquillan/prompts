package prompts

import (
	"fmt"
	"strings"

	"github.com/jessevdk/go-flags"
)

const (
	defaultEnder = "$"
	defaultSeparator = ":"
)

type Prompt struct {
	Opts
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

func (p *Prompt) Print() string {
	s := p.String(p.ExitCode)
	fmt.Print(s)
	return s
}

func (p *Prompt) PrintWithExitCode(exitCode int) string {
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

func (p *Prompt) ParseArgs() *Prompt {
	p.Opts = Opts{}
	_, err := flags.Parse(&p.Opts)
	if err != nil {
		panic(err)
	}

	return p
}

func (p *Prompt) WithArgs(args []string) *Prompt {
	p.Opts = Opts{}
	_, err := flags.ParseArgs(&p.Opts, args)
	if err != nil {
		panic(err)
	}

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
