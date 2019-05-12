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
	Ender
	Components        []Component
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

func (p *Prompt) getFallBack() string {
	if p.FallBack == "" {
		return defaultFallBack
	}

	return p.FallBack
}

func (p *Prompt) String(exitCode int) (output string) {
	defer func() {
		if err := recover(); err != nil {
			if p.NoRecover {
				panic(err)
			}
		output = p.getFallBack()
		}
	}()

	var subStrings []string
	for _, component := range p.Components {
		e := component.MakeElement()
		if e.Length > 0 || p.Opts.ShowAll {
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

func (p *Prompt) WithOpts(opts Opts) *Prompt {
	p.Opts = opts
	return p
}
