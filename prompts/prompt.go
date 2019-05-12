package prompts

import (
	"fmt"
	"strings"

	"github.com/jessevdk/go-flags"
)

const (
	defaultEnder     = "$"
	defaultSeparator = ":"
	defaultPanicMsg  = "<Prompt panicked!>"
	// Upon a panic, this prompt will be used instead
	defaultFallBack = defaultPanicMsg + defaultEnder + " "
)

type Prompt struct {
	Opts
	PromptState
	Ender
	Components []Component
	Separator  string
}

type PromptState struct {
	CurrentLength int
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

func (p *Prompt) String(exitCode int) (output string) {
	defer func() {
		// Recover and return a fall back prompt
		if err := recover(); err != nil {
			if p.NoRecover {
				panic(err)
			}
		output = p.getFallBack()
		}
	}()

	var elements, unresolvedElements []*Element
	var dynamicComponents []*DynamicComponent
	for _, component := range p.Components {
		if dComponent, ok := component.(*DynamicComponent); ok {
			e := Element{}
			elements = append(elements, &e)
			unresolvedElements = append(unresolvedElements, &e)
			dynamicComponents = append(dynamicComponents, dComponent)
		} else if e := component.MakeElement(); e.Length > 0 {
			elements = append(elements, &e)
			p.CurrentLength += e.Length
		}
	}

	for i, dComponent := range dynamicComponents {
		e := dComponent.MakeElement()
		*unresolvedElements[i] = e
		p.CurrentLength += e.Length
	}

	var subStrings []string
	for _, e := range elements {
		if e != nil && e.Length > 0 {
			subStrings = append(subStrings, e.Output)
		}
	}

	return fmt.Sprintf("%s%s ", strings.Join(subStrings, p.Separator), p.End(exitCode).Output)
}

func (p *Prompt) WithComponent(c Component) *Prompt {
	p.Components = append(p.Components, c)
	return p
}

func (p *Prompt) WithDynamicComponent(c *DynamicComponent) *Prompt {
	c.PromptState = &p.PromptState
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

func (p *Prompt) WithOpts(opts Opts) *Prompt {
	p.Opts = opts
	return p
}

func (p *Prompt) getFallBack() string {
	if p.FallBack == "" {
		return defaultFallBack
	}
	return p.FallBack
}
